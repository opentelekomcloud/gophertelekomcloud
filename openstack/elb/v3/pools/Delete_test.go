package pools_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/pools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/pools/pool-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})
	th.AssertNoErr(t, pools.Delete(serviceClient(), "pool-id"))
}

func TestDeleteInvalidID(t *testing.T) {
	if err := pools.Delete(serviceClient(), "../listeners/listener-id"); err == nil {
		t.Fatal("expected invalid ID error")
	}
}
