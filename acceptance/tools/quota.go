package tools

import (
	"sync"
	"testing"
)

// defaultQuotas holds the built-in capacity for well-known limited
// resources. Add an entry here (or call RegisterQuotaCapacity before first
// use) to support quota tracking for an additional service/resource.
var defaultQuotas = map[string]int{
	"eip": 3,
}

var (
	quotaRegistryMu sync.Mutex
	quotaRegistry   = map[string]chan struct{}{}
)

// quotaChannel returns the shared semaphore channel for the given resource,
// creating it on first use. The channel capacity is taken from
// defaultQuotas (or a previous RegisterQuotaCapacity call) when known,
// otherwise it defaults to a single slot (serialized access), so an
// AcquireQuota call for an unregistered resource never silently grants more
// concurrency than is safe.
func quotaChannel(resource string) chan struct{} {
	quotaRegistryMu.Lock()
	defer quotaRegistryMu.Unlock()

	if ch, ok := quotaRegistry[resource]; ok {
		return ch
	}

	capacity := 1
	if def, ok := defaultQuotas[resource]; ok {
		capacity = def
	}

	ch := make(chan struct{}, capacity)
	quotaRegistry[resource] = ch
	return ch
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
	quotaRegistry[resource] = make(chan struct{}, capacity)
}

// AcquireQuota reserves `units` slots of the named resource quota (e.g.
// "eip"), blocking until enough slots are free. Every test that consumes one
// or more units of a limited resource should call AcquireQuota before
// creating that resource; the reserved units are released automatically
// when the test (and its subtests) finish via t.Cleanup, allowing other
// tests waiting on the same quota to proceed.
//
// The returned release function may be called earlier than test cleanup if
// a test wants to free its quota units as soon as it no longer needs them;
// it is safe to call multiple times (only the first call has an effect).
func AcquireQuota(t *testing.T, resource string, units int) func() {
	t.Helper()
	if units <= 0 {
		units = 1
	}

	ch := quotaChannel(resource)
	if units > cap(ch) {
		t.Fatalf("requested %d quota units for %q, but its total capacity is only %d", units, resource, cap(ch))
	}

	for i := 0; i < units; i++ {
		ch <- struct{}{}
	}

	var release func()
	var once sync.Once
	release = func() {
		once.Do(func() {
			for i := 0; i < units; i++ {
				<-ch
			}
		})
	}

	t.Cleanup(release)
	return release
}
