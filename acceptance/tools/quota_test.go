package tools_test

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
)

func TestAcquireQuotaSerializesBeyondCapacity(t *testing.T) {
	resource := tools.RandomString("quota-test-", 6)
	tools.RegisterQuotaCapacity(resource, 2)

	var inFlight int32
	var maxInFlight int32
	var mu sync.Mutex

	worker := func() {
		release := tools.AcquireQuota(t, resource, 1)
		defer release()

		cur := atomic.AddInt32(&inFlight, 1)
		mu.Lock()
		if cur > maxInFlight {
			maxInFlight = cur
		}
		mu.Unlock()

		time.Sleep(20 * time.Millisecond)
		atomic.AddInt32(&inFlight, -1)
	}

	var wg sync.WaitGroup
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker()
		}()
	}
	wg.Wait()

	if maxInFlight > 2 {
		t.Fatalf("expected at most 2 concurrent quota holders, got %d", maxInFlight)
	}
}

func TestAcquireQuotaReleaseIsIdempotent(t *testing.T) {
	resource := tools.RandomString("quota-test-", 6)
	tools.RegisterQuotaCapacity(resource, 1)

	release := tools.AcquireQuota(t, resource, 1)
	release()
	release()

	// A second acquire must succeed immediately since the first was released.
	done := make(chan struct{})
	go func() {
		second := tools.AcquireQuota(t, resource, 1)
		second()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("expected second AcquireQuota to proceed after release")
	}
}

// TestAcquireQuotaIsCrossProcess proves the quota is enforced across
// separate OS processes, not just goroutines within one binary. This is the
// scenario `go test ./acceptance/...` actually creates: one test binary per
// package, executed concurrently by the go tool. It spawns several worker
// subprocesses (this same test binary, re-invoked with
// TEST_QUOTA_WORKER=1) that each acquire the quota, record their
// acquisition into a shared counter file while holding it, and fail if more
// than `capacity` workers are ever recorded as concurrently holding it.
func TestAcquireQuotaIsCrossProcess(t *testing.T) {
	const capacity = 2
	const workers = 6

	if os.Getenv("TEST_QUOTA_WORKER") == "1" {
		runQuotaWorker(t)
		return
	}

	quotaDir := t.TempDir()
	counterPath := quotaDir + "/counter"
	resultsPath := quotaDir + "/violation"

	resource := tools.RandomString("quota-proc-test-", 6)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cmd := exec.Command(os.Args[0], "-test.run=^TestAcquireQuotaIsCrossProcess$", "-test.v")
			cmd.Env = append(os.Environ(),
				"TEST_QUOTA_WORKER=1",
				"QUOTA_WORKER_RESOURCE="+resource,
				"QUOTA_WORKER_CAPACITY="+fmt.Sprint(capacity),
				"QUOTA_WORKER_COUNTER_PATH="+counterPath,
				"QUOTA_WORKER_VIOLATION_PATH="+resultsPath,
			)
			if out, err := cmd.CombinedOutput(); err != nil {
				t.Errorf("worker subprocess failed: %v\n%s", err, out)
			}
		}()
	}
	wg.Wait()

	if _, err := os.Stat(resultsPath); err == nil {
		data, _ := os.ReadFile(resultsPath)
		t.Fatalf("quota was violated across processes: %s", data)
	}
}

// runQuotaWorker is the subprocess entry point invoked by
// TestAcquireQuotaIsCrossProcess. It acquires the quota, atomically bumps a
// shared counter file (via the same flock mechanism AcquireQuota itself
// uses, applied to a separate lock file) to track how many processes
// currently hold the quota, and records a violation if that count ever
// exceeds the configured capacity.
func runQuotaWorker(t *testing.T) {
	resource := os.Getenv("QUOTA_WORKER_RESOURCE")
	capacity := os.Getenv("QUOTA_WORKER_CAPACITY")
	counterPath := os.Getenv("QUOTA_WORKER_COUNTER_PATH")
	violationPath := os.Getenv("QUOTA_WORKER_VIOLATION_PATH")

	var cap int
	_, _ = fmt.Sscanf(capacity, "%d", &cap)
	tools.RegisterQuotaCapacity(resource, cap)

	release := tools.AcquireQuota(t, resource, 1)
	defer release()

	cur := bumpCounter(t, counterPath, 1)
	if cur > cap {
		_ = os.WriteFile(violationPath, []byte(fmt.Sprintf("observed %d concurrent holders, capacity is %d", cur, cap)), 0o644)
	}

	time.Sleep(300 * time.Millisecond)

	bumpCounter(t, counterPath, -1)
}

// bumpCounter atomically adds delta to the integer stored in path (creating
// it with value 0 if absent), guarded by an flock on the same file, and
// returns the value immediately after the update.
func bumpCounter(t *testing.T, path string, delta int) int {
	t.Helper()

	f, err := os.OpenFile(path+".lock", os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		t.Fatalf("failed to open counter lock: %v", err)
	}
	defer f.Close()

	if err := flockExclusive(f); err != nil {
		t.Fatalf("failed to lock counter: %v", err)
	}
	defer flockUnlock(f)

	cur := 0
	if data, err := os.ReadFile(path); err == nil && len(data) > 0 {
		_, _ = fmt.Sscanf(string(data), "%d", &cur)
	}
	cur += delta
	if err := os.WriteFile(path, []byte(fmt.Sprint(cur)), 0o644); err != nil {
		t.Fatalf("failed to write counter: %v", err)
	}
	return cur
}

func flockExclusive(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
}

func flockUnlock(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
}
