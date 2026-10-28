package monitors_test

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/monitors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	return client.ServiceClient()
}

const monitorJSON = `{
	"id":"monitor-id","name":"health-check","project_id":"project-id","type":"HTTP",
	"delay":10,"timeout":5,"max_retries":3,"max_retries_down":2,"http_method":"GET",
	"url_path":"/health","domain_name":"example.com","expected_codes":"200-299",
	"admin_state_up":true,"monitor_port":8080,"pools":[{"id":"pool-id"}],
	"created_at":"2026-10-26T10:00:00Z","updated_at":"2026-10-26T10:01:00Z",
	"http_version":"HTTP/1.1","content_check_enable":false,"send_content":"ping",
	"receive_content":"pong","content_format":"TEXT"
}`

func assertMonitor(t *testing.T, actual *monitors.Monitor) {
	t.Helper()
	th.AssertDeepEquals(t, monitors.Monitor{
		ID: "monitor-id", Name: "health-check", ProjectID: "project-id", Type: monitors.TypeHTTP,
		Delay: 10, Timeout: 5, MaxRetries: 3, MaxRetriesDown: 2, HTTPMethod: "GET",
		URLPath: "/health", DomainName: "example.com", ExpectedCodes: "200-299",
		AdminStateUp: true, MonitorPort: 8080, Pools: []structs.ResourceRef{{ID: "pool-id"}},
		CreatedAt: "2026-10-26T10:00:00Z", UpdatedAt: "2026-10-26T10:01:00Z",
		HTTPVersion: "HTTP/1.1", ContentCheckEnable: false, SendContent: "ping",
		ReceiveContent: "pong", ContentFormat: "TEXT",
	}, *actual)
}
