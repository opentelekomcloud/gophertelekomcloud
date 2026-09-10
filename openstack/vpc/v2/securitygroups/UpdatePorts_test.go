package securitygroups_test

import (
	"net/http"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/securitygroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	c := client.ServiceClient()
	c.ProjectID = "project-id"
	return c
}

func TestUpdatePorts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/security-groups/group-id/instance/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"ports":[{"id":"port-1"},{"id":"port-2"}],"action":"add"}`)
		_, _ = w.Write([]byte(`{"fail":[]}`))
	})
	actual, err := securitygroups.UpdatePorts(serviceClient(), "group-id", securitygroups.UpdatePortsOpts{
		Ports: []securitygroups.Port{{ID: "port-1"}, {ID: "port-2"}}, Action: "add",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(actual))
}

func TestUpdatePortsFailureResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/security-groups/group-id/instance/action", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"fail":[{"id":"port-1","error_code":"VPC.0608","error_msg":"not found"}]}`))
	})
	actual, err := securitygroups.UpdatePorts(serviceClient(), "group-id", securitygroups.UpdatePortsOpts{
		Ports: []securitygroups.Port{{ID: "port-1"}}, Action: "remove",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "port-1", actual[0].ID)
	th.AssertEquals(t, "VPC.0608", actual[0].ErrorCode)
}

func TestUpdatePortsInvalidID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	actual, err := securitygroups.UpdatePorts(serviceClient(), "../vpcs/vpc-id", securitygroups.UpdatePortsOpts{
		Ports: []securitygroups.Port{{ID: "port-1"}}, Action: "add",
	})
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and invalid ID error, got %#v, %v", actual, err)
	}
}
