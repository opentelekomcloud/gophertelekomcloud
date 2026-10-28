package monitors_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/monitors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/healthmonitors/monitor-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{"healthmonitor":{"delay":10,"timeout":5,"max_retries":3,"max_retries_down":2,"url_path":"/health","domain_name":"example.com","http_method":"GET","expected_codes":"200-299","name":"health-check","admin_state_up":true,"monitor_port":8080,"type":"HTTP","http_version":"HTTP/1.1","content_check_enable":false,"send_content":"","receive_content":"pong","content_format":"TEXT"}}`)
		_, _ = w.Write([]byte(`{"healthmonitor":` + monitorJSON + `}`))
	})
	empty, receive, version := "", "pong", "HTTP/1.1"
	actual, err := monitors.Update(serviceClient(), "monitor-id", monitors.UpdateOpts{
		Delay: 10, Timeout: 5, MaxRetries: 3, MaxRetriesDown: 2, URLPath: "/health",
		DomainName: "example.com", HTTPMethod: "GET", ExpectedCodes: "200-299", Name: "health-check",
		AdminStateUp: pointerto.Bool(true), MonitorPort: 8080, Type: "HTTP", HTTPVersion: &version,
		ContentCheckEnable: pointerto.Bool(false), SendContent: &empty, ReceiveContent: &receive, ContentFormat: "TEXT",
	})
	th.AssertNoErr(t, err)
	assertMonitor(t, actual)
}

func TestUpdateOmittedFields(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/healthmonitors/monitor-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{"healthmonitor":{}}`)
		_, _ = w.Write([]byte(`{"healthmonitor":{"id":"monitor-id"}}`))
	})
	actual, err := monitors.Update(serviceClient(), "monitor-id", monitors.UpdateOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "monitor-id", actual.ID)
}

func TestUpdateErrors(t *testing.T) {
	for _, status := range []int{http.StatusConflict, http.StatusOK, http.StatusAccepted} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/healthmonitors/monitor-id", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				_, _ = w.Write([]byte(`{"healthmonitor":`))
			})
			actual, err := monitors.Update(serviceClient(), "monitor-id", monitors.UpdateOpts{Name: "updated"})
			if err == nil || actual != nil {
				t.Fatalf("expected error and nil monitor, got %#v, %v", actual, err)
			}
		})
	}
}

func TestUpdateInvalidID(t *testing.T) {
	actual, err := monitors.Update(serviceClient(), "../pools/pool-id", monitors.UpdateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
