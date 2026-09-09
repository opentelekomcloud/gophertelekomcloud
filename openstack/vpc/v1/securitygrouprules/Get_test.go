package securitygrouprules_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/securitygrouprules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/security-group-rules/rule-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"security_group_rule":` + ruleJSON + `}`))
	})
	actual, err := securitygrouprules.Get(serviceClient(), "rule-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "rule-id", actual.ID)
	th.AssertEquals(t, 80, *actual.PortRangeMax)
}

func TestGetInvalidID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	actual, err := securitygrouprules.Get(serviceClient(), "../security-groups/group-id")
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and invalid ID error, got %#v, %v", actual, err)
	}
}
