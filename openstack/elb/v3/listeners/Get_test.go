package listeners_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/listeners/listener-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"listener":` + listenerJSON + `}`))
	})
	actual, err := listeners.Get(serviceClient(), "listener-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "listener-id", actual.ID)
}

func TestGetInvalidID(t *testing.T) {
	actual, err := listeners.Get(serviceClient(), "../pools/pool-id")
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
