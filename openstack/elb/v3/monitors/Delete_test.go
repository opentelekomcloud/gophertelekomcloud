package monitors_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/monitors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/healthmonitors/monitor-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})
	th.AssertNoErr(t, monitors.Delete(serviceClient(), "monitor-id"))
}

func TestDeleteErrors(t *testing.T) {
	for _, status := range []int{http.StatusConflict, http.StatusOK, http.StatusAccepted} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/healthmonitors/monitor-id", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
			})
			if err := monitors.Delete(serviceClient(), "monitor-id"); err == nil {
				t.Fatal("expected delete error")
			}
		})
	}
}

func TestDeleteInvalidID(t *testing.T) {
	if err := monitors.Delete(serviceClient(), "../pools/pool-id"); err == nil {
		t.Fatal("expected invalid ID error")
	}
}
