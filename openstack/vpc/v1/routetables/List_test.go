package routetables_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/routetables"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListAllOptionsAndPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	requests := 0
	th.Mux.HandleFunc("/project-id/routetables", func(w http.ResponseWriter, r *http.Request) {
		requests++
		th.AssertEquals(t, "2", r.URL.Query().Get("limit"))
		th.AssertEquals(t, "table-filter", r.URL.Query().Get("id"))
		th.AssertEquals(t, "vpc-id", r.URL.Query().Get("vpc_id"))
		th.AssertEquals(t, "network-id", r.URL.Query().Get("subnet_id"))
		switch r.URL.Query().Get("marker") {
		case "start":
			_, _ = fmt.Fprintf(w, `{"routetables":[%s,%s]}`, strings.ReplaceAll(routeTableJSON, "route-table-id", "table-1"), strings.ReplaceAll(routeTableJSON, "route-table-id", "table-2"))
		case "table-2":
			_, _ = fmt.Fprintf(w, `{"routetables":[%s]}`, strings.ReplaceAll(routeTableJSON, "route-table-id", "table-3"))
		case "table-3":
			_, _ = w.Write([]byte(`{"routetables":[]}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})
	limit := 2
	actual, err := routetables.List(serviceClient(), routetables.ListOpts{
		Limit: &limit, Marker: "start", ID: "table-filter", VpcID: "vpc-id", SubnetID: "network-id",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, len(actual))
	th.AssertEquals(t, 3, requests)
}

func TestListInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/routetables", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"routetables":`))
	})
	actual, err := routetables.List(serviceClient(), routetables.ListOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected extraction error, got %#v, %v", actual, err)
	}
}
