package monitors_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/monitors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/healthmonitors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"healthmonitor":{"pool_id":"pool-id","type":"HTTP","delay":10,"timeout":5,"max_retries":3,"max_retries_down":2,"url_path":"/health","domain_name":"example.com","http_method":"GET","expected_codes":"200-299","project_id":"project-id","name":"health-check","admin_state_up":true,"monitor_port":8080,"http_version":"HTTP/1.1","content_check_enable":false,"send_content":"ping","receive_content":"pong","content_format":"TEXT"}}`)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"healthmonitor":` + monitorJSON + `}`))
	})
	actual, err := monitors.Create(serviceClient(), monitors.CreateOpts{
		PoolID: "pool-id", Type: monitors.TypeHTTP, Delay: 10, Timeout: 5, MaxRetries: 3,
		MaxRetriesDown: 2, URLPath: "/health", DomainName: "example.com", HTTPMethod: "GET",
		ExpectedCodes: "200-299", ProjectID: "project-id", Name: "health-check",
		AdminStateUp: pointerto.Bool(true), MonitorPort: 8080, HTTPVersion: "HTTP/1.1",
		ContentCheckEnable: pointerto.Bool(false), SendContent: "ping", ReceiveContent: "pong", ContentFormat: "TEXT",
	})
	th.AssertNoErr(t, err)
	assertMonitor(t, actual)
}

func TestCreateRequiredFields(t *testing.T) {
	valid := monitors.CreateOpts{PoolID: "pool-id", Type: monitors.TypeUDPConnect, Delay: 10, Timeout: 5, MaxRetries: 3}
	for _, field := range []string{"pool", "type", "delay", "timeout", "retries"} {
		t.Run(field, func(t *testing.T) {
			opts := valid
			switch field {
			case "pool":
				opts.PoolID = ""
			case "type":
				opts.Type = ""
			case "delay":
				opts.Delay = 0
			case "timeout":
				opts.Timeout = 0
			case "retries":
				opts.MaxRetries = 0
			}
			actual, err := monitors.Create(serviceClient(), opts)
			if err == nil || actual != nil {
				t.Fatalf("expected required field error, got %#v, %v", actual, err)
			}
		})
	}
}

func TestCreateErrors(t *testing.T) {
	for _, status := range []int{http.StatusBadRequest, http.StatusCreated} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/healthmonitors", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				_, _ = w.Write([]byte(`{"healthmonitor":`))
			})
			actual, err := monitors.Create(serviceClient(), monitors.CreateOpts{
				PoolID: "pool-id", Type: monitors.TypeTCP, Delay: 10, Timeout: 5, MaxRetries: 3,
			})
			if err == nil || actual != nil {
				t.Fatalf("expected error and nil monitor, got %#v, %v", actual, err)
			}
		})
	}
}
