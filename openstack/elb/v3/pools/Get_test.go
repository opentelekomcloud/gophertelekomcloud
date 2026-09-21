package pools_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/pools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/pools/pool-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"pool":` + poolJSON + `}`))
	})
	actual, err := pools.Get(serviceClient(), "pool-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "pool-id", actual.ID)
	th.AssertEquals(t, "PROTECTED", actual.ProtectionStatus)
	th.AssertEquals(t, 50, actual.AZAffinity.AZMinimumHealthyMemberPercentage)
}

func TestGetInvalidID(t *testing.T) {
	actual, err := pools.Get(serviceClient(), "../listeners/listener-id")
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
