package tools_test

import (
	"sync"
	"sync/atomic"
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
