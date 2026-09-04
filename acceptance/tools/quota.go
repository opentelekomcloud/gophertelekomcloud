package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

// defaultQuotas holds the built-in capacity for well-known limited
// resources. Add an entry here (or call RegisterQuotaCapacity before first
// use) to support quota tracking for an additional service/resource.
var defaultQuotas = map[string]int{
	"eip": 3,
}

// Quota slots are enforced with real, cross-process file locks (rather than
// an in-memory channel/semaphore) because `go test ./acceptance/...` runs a
// separate test binary per package, and those binaries execute concurrently.
// An in-process-only semaphore would give every package its own independent
// copy of the quota, allowing far more concurrent requests against a
// resource (e.g. EIP creation) than the account actually supports, which can
// manifest as unrelated failures such as "TLS handshake timeout" under CI
// load. Lock files are placed in a directory shared by all test binaries of
// a run.
var (
	quotaRegistryMu sync.Mutex
	quotaRegistry   = map[string]int{}
)

// quotaDir returns the shared directory used to store per-resource quota
// slot lock files, creating it if necessary. It can be overridden with the
// OTC_ACC_QUOTA_DIR environment variable (e.g. to isolate concurrent CI runs
// on a shared host); otherwise a fixed path under os.TempDir() is used so
// all test binaries in the same run agree on the location.
func quotaDir() string {
	dir := os.Getenv("OTC_ACC_QUOTA_DIR")
	if dir == "" {
		dir = filepath.Join(os.TempDir(), "otc-acceptance-quota")
	}
	_ = os.MkdirAll(dir, 0o755)
	return dir
}

// quotaCapacity returns the configured capacity for a resource, registering
// it with the default (or 1, if unknown) on first use.
func quotaCapacity(resource string) int {
	quotaRegistryMu.Lock()
	defer quotaRegistryMu.Unlock()

	if capacity, ok := quotaRegistry[resource]; ok {
		return capacity
	}

	capacity := 1
	if def, ok := defaultQuotas[resource]; ok {
		capacity = def
	}
	quotaRegistry[resource] = capacity
	return capacity
}

// RegisterQuotaCapacity sets the capacity for a resource quota before its
// first use. This lets tests/suites configure a limit for a resource that
// isn't one of the built-in defaults (e.g. a new service with its own
// availability limit). It has no effect if the quota for the resource has
// already been created by a prior AcquireQuota/RegisterQuotaCapacity call.
func RegisterQuotaCapacity(resource string, capacity int) {
	quotaRegistryMu.Lock()
	defer quotaRegistryMu.Unlock()

	if _, ok := quotaRegistry[resource]; ok {
		return
	}
	if capacity <= 0 {
		capacity = 1
	}
	quotaRegistry[resource] = capacity
}

// slotPath returns the lock-file path for a given resource/slot pair.
func slotPath(resource string, slot int) string {
	return filepath.Join(quotaDir(), fmt.Sprintf("%s.slot%d.lock", resource, slot))
}

// tryLockSlot attempts a non-blocking exclusive flock on the given slot's
// lock file. It returns the open file handle (to be closed/unlocked on
// release) and true on success, or false if the slot is already held by
// another process/goroutine.
func tryLockSlot(resource string, slot int) (*os.File, bool) {
	f, err := os.OpenFile(slotPath(resource, slot), os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return nil, false
	}
	if err := tryLockFile(f); err != nil {
		_ = f.Close()
		return nil, false
	}
	return f, true
}

// AcquireQuota reserves `units` slots of the named resource quota (e.g.
// "eip"), blocking until enough slots are free across every acceptance test
// process running concurrently. Every test that consumes one or more units
// of a limited resource should call AcquireQuota before creating that
// resource; the reserved units are released automatically when the test
// (and its subtests) finish via t.Cleanup, allowing other tests waiting on
// the same quota to proceed.
//
// The returned release function may be called earlier than test cleanup if
// a test wants to free its quota units as soon as it no longer needs them;
// it is safe to call multiple times (only the first call has an effect).
func AcquireQuota(t *testing.T, resource string, units int) func() {
	t.Helper()
	if units <= 0 {
		units = 1
	}

	capacity := quotaCapacity(resource)
	if units > capacity {
		t.Fatalf("requested %d quota units for %q, but its total capacity is only %d", units, resource, capacity)
	}

	held := make([]*os.File, 0, units)
	for len(held) < units {
		acquiredThisPass := false
		for slot := 0; slot < capacity && len(held) < units; slot++ {
			if f, ok := tryLockSlot(resource, slot); ok {
				held = append(held, f)
				acquiredThisPass = true
			}
		}
		if len(held) < units {
			if !acquiredThisPass {
				time.Sleep(500 * time.Millisecond)
			}
		}
	}

	var release func()
	var once sync.Once
	release = func() {
		once.Do(func() {
			for _, f := range held {
				_ = unlockFile(f)
				_ = f.Close()
			}
		})
	}

	t.Cleanup(release)
	return release
}
