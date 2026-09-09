package securitygroups_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/securitygroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/security-groups/group-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})
	th.AssertNoErr(t, securitygroups.Delete(serviceClient(), "group-id"))
}

func TestDeleteInvalidID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	if err := securitygroups.Delete(serviceClient(), "../vpcs/vpc-id"); err == nil {
		t.Fatal("expected invalid ID error")
	}
}
