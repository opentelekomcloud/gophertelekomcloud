package listeners_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestForceDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/listeners/listener-id/force", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})
	th.AssertNoErr(t, listeners.ForceDelete(serviceClient(), "listener-id"))
}

func TestForceDeleteInvalidID(t *testing.T) {
	if err := listeners.ForceDelete(serviceClient(), "../pools/pool-id"); err == nil {
		t.Fatal("expected invalid ID error")
	}
}
