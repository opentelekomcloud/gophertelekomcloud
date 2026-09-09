package securitygrouprules_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/securitygrouprules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/security-group-rules/rule-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})
	th.AssertNoErr(t, securitygrouprules.Delete(serviceClient(), "rule-id"))
}

func TestDeleteInvalidID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	if err := securitygrouprules.Delete(serviceClient(), "../security-groups/group-id"); err == nil {
		t.Fatal("expected invalid ID error")
	}
}
