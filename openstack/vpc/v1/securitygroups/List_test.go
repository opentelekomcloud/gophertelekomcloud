package securitygroups_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/securitygroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListAllOptionsAndPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	requests := 0
	th.Mux.HandleFunc("/project-id/security-groups", func(w http.ResponseWriter, r *http.Request) {
		requests++
		th.TestMethod(t, r, http.MethodGet)
		th.AssertEquals(t, "2", r.URL.Query().Get("limit"))
		th.AssertEquals(t, "vpc-id", r.URL.Query().Get("vpc_id"))
		th.AssertEquals(t, "0", r.URL.Query().Get("enterprise_project_id"))
		switch r.URL.Query().Get("marker") {
		case "start":
			_, _ = fmt.Fprintf(w, `{"security_groups":[%s,%s]}`,
				strings.ReplaceAll(groupJSON, `"group-id"`, `"group-1"`),
				strings.ReplaceAll(groupJSON, `"group-id"`, `"group-2"`))
		case "group-2":
			_, _ = fmt.Fprintf(w, `{"security_groups":[%s]}`,
				strings.ReplaceAll(groupJSON, `"group-id"`, `"group-3"`))
		case "group-3":
			_, _ = w.Write([]byte(`{"security_groups":[]}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})
	limit := 2
	actual, err := securitygroups.List(serviceClient(), securitygroups.ListOpts{
		Marker: "start", Limit: &limit, VpcID: "vpc-id", EnterpriseProjectID: "0",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, len(actual))
	th.AssertEquals(t, "group-3", actual[2].ID)
	th.AssertEquals(t, 3, requests)
}

func TestListNoOptionsAndInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/security-groups", func(w http.ResponseWriter, r *http.Request) {
		th.AssertEquals(t, "", r.URL.RawQuery)
		_, _ = w.Write([]byte(`{"security_groups":`))
	})
	actual, err := securitygroups.List(serviceClient(), securitygroups.ListOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and extraction error, got %#v, %v", actual, err)
	}
}
