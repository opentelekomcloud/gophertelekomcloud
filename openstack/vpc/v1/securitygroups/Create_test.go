package securitygroups_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/securitygroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/security-groups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"security_group":{"name":"sg-test","vpc_id":"vpc-id","enterprise_project_id":"0"}}`)
		_, _ = w.Write([]byte(`{"security_group":` + groupJSON + `}`))
	})
	actual, err := securitygroups.Create(serviceClient(), securitygroups.CreateOpts{
		Name: "sg-test", VpcID: "vpc-id", EnterpriseProjectID: "0",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "group-id", actual.ID)
	th.AssertEquals(t, "default", actual.Description)
	th.AssertEquals(t, 1, len(actual.SecurityGroupRules))
}

func TestCreateMissingRequiredOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	actual, err := securitygroups.Create(serviceClient(), securitygroups.CreateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and validation error, got %#v, %v", actual, err)
	}
}

func TestCreateError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/security-groups", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	actual, err := securitygroups.Create(serviceClient(), securitygroups.CreateOpts{Name: "sg-test"})
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and request error, got %#v, %v", actual, err)
	}
}
