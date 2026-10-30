package certificates_test

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/certificates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	return client.ServiceClient()
}

const certificateJSON = `{
	"id":"certificate-id","project_id":"project-id","name":"test-certificate",
	"description":"test description","type":"server","domain":"example.com,*.example.org",
	"private_key":"test-key","certificate":"test-certificate-content","admin_state_up":true,
	"created_at":"2026-10-30T10:00:00Z","updated_at":"2026-10-30T10:01:00Z",
	"expire_time":"2027-10-30T10:00:00Z"
}`

func assertCertificate(t *testing.T, actual *certificates.Certificate) {
	t.Helper()
	th.AssertDeepEquals(t, certificates.Certificate{
		ID: "certificate-id", ProjectID: "project-id", Name: "test-certificate",
		Description: "test description", Type: "server", Domain: "example.com,*.example.org",
		PrivateKey: "test-key", Certificate: "test-certificate-content", AdminStateUp: true,
		CreatedAt: "2026-10-30T10:00:00Z", UpdatedAt: "2026-10-30T10:01:00Z",
		ExpireTime: "2027-10-30T10:00:00Z",
	}, *actual)
}
