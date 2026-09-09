package securitygroups_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/securitygroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/security-groups/group-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"security_group":` + groupJSON + `}`))
	})
	actual, err := securitygroups.Get(serviceClient(), "group-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "group-id", actual.ID)
	th.AssertEquals(t, "vpc-id", actual.VpcID)
	th.AssertEquals(t, (*int)(nil), actual.SecurityGroupRules[0].PortRangeMin)
}

func TestGetInvalidID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	actual, err := securitygroups.Get(serviceClient(), "../vpcs/vpc-id")
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and invalid ID error, got %#v, %v", actual, err)
	}
}

func TestGetInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/security-groups/group-id", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"security_group":`))
	})
	actual, err := securitygroups.Get(serviceClient(), "group-id")
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and extraction error, got %#v, %v", actual, err)
	}
}
