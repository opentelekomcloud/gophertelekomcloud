package listeners_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/listeners/listener-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})
	th.AssertNoErr(t, listeners.Delete(serviceClient(), "listener-id"))
}

func TestDeleteInvalidID(t *testing.T) {
	if err := listeners.Delete(serviceClient(), "../pools/pool-id"); err == nil {
		t.Fatal("expected invalid ID error")
	}
}
