package securitygrouprules_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/securitygrouprules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListAllOptionsAndPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	requests := 0
	th.Mux.HandleFunc("/project-id/security-group-rules", func(w http.ResponseWriter, r *http.Request) {
		requests++
		th.AssertEquals(t, "2", r.URL.Query().Get("limit"))
		th.AssertEquals(t, "group-id", r.URL.Query().Get("security_group_id"))
		th.AssertEquals(t, "10.0.0.0/24", r.URL.Query().Get("remote_ip_prefix"))
		switch r.URL.Query().Get("marker") {
		case "start":
			_, _ = fmt.Fprintf(w, `{"security_group_rules":[%s,%s]}`,
				strings.ReplaceAll(ruleJSON, `"rule-id"`, `"rule-1"`),
				strings.ReplaceAll(ruleJSON, `"rule-id"`, `"rule-2"`))
		case "rule-2":
			_, _ = fmt.Fprintf(w, `{"security_group_rules":[%s]}`,
				strings.ReplaceAll(ruleJSON, `"rule-id"`, `"rule-3"`))
		case "rule-3":
			_, _ = w.Write([]byte(`{"security_group_rules":[]}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})
	limit := 2
	actual, err := securitygrouprules.List(serviceClient(), securitygrouprules.ListOpts{
		Marker: "start", Limit: &limit, SecurityGroupID: "group-id", RemoteIPPrefix: "10.0.0.0/24",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, len(actual))
	th.AssertEquals(t, "rule-3", actual[2].ID)
	th.AssertEquals(t, 3, requests)
}

func TestListInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/security-group-rules", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"security_group_rules":`))
	})
	actual, err := securitygrouprules.List(serviceClient(), securitygrouprules.ListOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and extraction error, got %#v, %v", actual, err)
	}
}
