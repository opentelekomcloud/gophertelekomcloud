package monitors_test

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/monitors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListFilters(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/healthmonitors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.AssertDeepEquals(t, url.Values{
			"id": {"legacy-id", "id-1", "id-2"}, "name": {"legacy-name", "name-1", "name-2"},
			"type": {"HTTP", "TCP", "UDP_CONNECT"}, "delay": {"10", "15", "20"},
			"timeout": {"5"}, "max_retries": {"3", "4", "5"}, "http_method": {"GET", "HEAD", "OPTIONS"},
			"url_path": {"/health", "/ready", "/live"}, "expected_codes": {"200", "201", "202"},
			"admin_state_up": {"false"}, "enterprise_project_id": {"enterprise-1", "enterprise-2"},
			"limit": {"1"}, "marker": {"start"}, "page_reverse": {"true"}, "monitor_port": {"80", "8080"},
			"domain_name": {"example.com", "test.example"}, "max_retries_down": {"2", "3"},
			"http_version": {"HTTP/1.0", "HTTP/1.1"}, "content_check_enable": {"false"},
			"content_format": {"TEXT", "HEX"}, "tenant_id": {"tenant-id"}, "project_id": {"project-id"},
			"pool_id": {"pool-id"}, "status": {"ACTIVE"}, "sort_key": {"name"}, "sort_dir": {"desc"},
		}, r.URL.Query())
		_, _ = w.Write([]byte(`{"healthmonitors":[` + monitorJSON + `],"page_info":{}}`))
	})
	actual, err := monitors.List(serviceClient(), monitors.ListOpts{
		ID: "legacy-id", Name: "legacy-name", Type: "HTTP", Delay: 10, Timeout: 5, MaxRetries: 3,
		HTTPMethod: "GET", URLPath: "/health", ExpectedCodes: "200", AdminStateUp: pointerto.Bool(false),
		EnterpriseProjectID: []string{"enterprise-1", "enterprise-2"}, Limit: 1, Marker: "start", PageReverse: true,
		MonitorPort: []int{80, 8080}, DomainName: []string{"example.com", "test.example"}, MaxRetriesDown: []int{2, 3},
		IDs: []string{"id-1", "id-2"}, Names: []string{"name-1", "name-2"}, Types: []string{"TCP", "UDP_CONNECT"},
		Delays: []int{15, 20}, MaxRetriesValues: []int{4, 5}, HTTPMethods: []string{"HEAD", "OPTIONS"},
		URLPaths: []string{"/ready", "/live"}, ExpectedCodesValues: []string{"201", "202"},
		HTTPVersion: []string{"HTTP/1.0", "HTTP/1.1"}, ContentCheckEnable: pointerto.Bool(false),
		ContentFormat: []string{"TEXT", "HEX"}, TenantID: "tenant-id", ProjectID: "project-id",
		PoolID: "pool-id", Status: "ACTIVE", SortKey: "name", SortDir: "desc",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(actual))
	assertMonitor(t, &actual[0])
}

func TestListPagination(t *testing.T) {
	for _, reverse := range []bool{false, true} {
		t.Run(fmt.Sprint(reverse), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			requests := 0
			th.Mux.HandleFunc("/healthmonitors", func(w http.ResponseWriter, r *http.Request) {
				requests++
				th.AssertEquals(t, "1", r.URL.Query().Get("limit"))
				th.AssertEquals(t, "health-check", r.URL.Query().Get("name"))
				marker := "next"
				if reverse {
					marker = "previous"
				}
				switch r.URL.Query().Get("marker") {
				case "start":
					_, _ = fmt.Fprintf(w, `{"healthmonitors":[%s],"page_info":{"next_marker":"next","previous_marker":"previous","current_count":1}}`, monitorJSON)
				case marker:
					_, _ = fmt.Fprintf(w, `{"healthmonitors":[%s],"page_info":{}}`,
						strings.ReplaceAll(monitorJSON, "monitor-id", "second-id"))
				default:
					t.Fatal("unexpected extra page")
				}
			})
			actual, err := monitors.List(serviceClient(), monitors.ListOpts{
				Limit: 1, Marker: "start", PageReverse: reverse, Name: "health-check",
			})
			th.AssertNoErr(t, err)
			th.AssertEquals(t, 3, requests)
			th.AssertEquals(t, 2, len(actual))
			th.AssertEquals(t, "second-id", actual[1].ID)
		})
	}
}

func TestListEmpty(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/healthmonitors", func(w http.ResponseWriter, r *http.Request) {
		th.AssertEquals(t, "", r.URL.RawQuery)
		_, _ = w.Write([]byte(`{"healthmonitors":[],"page_info":{}}`))
	})
	actual, err := monitors.List(serviceClient(), monitors.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(actual))
}

func TestListErrors(t *testing.T) {
	for _, status := range []int{http.StatusForbidden, http.StatusOK} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/healthmonitors", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				_, _ = w.Write([]byte(`{"healthmonitors":`))
			})
			actual, err := monitors.List(serviceClient(), monitors.ListOpts{})
			if err == nil || actual != nil {
				t.Fatalf("expected error and nil monitors, got %#v, %v", actual, err)
			}
		})
	}
}
