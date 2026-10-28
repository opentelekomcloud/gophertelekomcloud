package monitors_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/monitors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/healthmonitors/monitor-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"healthmonitor":` + monitorJSON + `}`))
	})
	actual, err := monitors.Get(serviceClient(), "monitor-id")
	th.AssertNoErr(t, err)
	assertMonitor(t, actual)
}

func TestGetErrors(t *testing.T) {
	for _, body := range []string{`{"error":"not found"}`, `{"healthmonitor":`} {
		t.Run(body, func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/healthmonitors/monitor-id", func(w http.ResponseWriter, r *http.Request) {
				if body == `{"error":"not found"}` {
					w.WriteHeader(http.StatusNotFound)
				}
				_, _ = w.Write([]byte(body))
			})
			actual, err := monitors.Get(serviceClient(), "monitor-id")
			if err == nil || actual != nil {
				t.Fatalf("expected error and nil monitor, got %#v, %v", actual, err)
			}
		})
	}
}

func TestGetInvalidID(t *testing.T) {
	actual, err := monitors.Get(serviceClient(), "../pools/pool-id")
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
