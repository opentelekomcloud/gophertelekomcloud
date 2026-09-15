package routetables_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/routetables"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/routetables", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"routetable":{"name":"route-table-test","routes":[{"type":"ecs","destination":"10.10.10.0/24","nexthop":"server-id","description":""}],"vpc_id":"vpc-id","description":"route table"}}`)
		_, _ = w.Write([]byte(`{"routetable":` + routeTableJSON + `}`))
	})
	actual, err := routetables.Create(serviceClient(), routetables.CreateOpts{
		Name: "route-table-test",
		Routes: []routetables.RouteOpts{{
			Type: "ecs", Destination: "10.10.10.0/24", NextHop: "server-id", Description: pointerto.String(""),
		}},
		VpcID: "vpc-id", Description: "route table",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "route-table-id", actual.ID)
	th.AssertEquals(t, "2026-09-15T08:00:00", actual.CreatedAt)
}

func TestCreateMissingVpcID(t *testing.T) {
	actual, err := routetables.Create(serviceClient(), routetables.CreateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected validation error, got %#v, %v", actual, err)
	}
}
