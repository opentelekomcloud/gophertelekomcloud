package routetables_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/routetables"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/routetables/route-table-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{"routetable":{"name":"updated","description":"","routes":{"add":[{"type":"ecs","destination":"10.0.0.0/24","nexthop":"server-id"}],"mod":[{"type":"ecs","destination":"10.1.0.0/24","nexthop":"server-id"}],"del":[{"destination":"10.2.0.0/24"}]}}}`)
		_, _ = w.Write([]byte(`{"routetable":` + routeTableJSON + `}`))
	})
	actual, err := routetables.Update(serviceClient(), "route-table-id", routetables.UpdateOpts{
		Name: "updated", Description: pointerto.String(""),
		Routes: &routetables.RouteAction{
			Add: []routetables.RouteOpts{{Type: "ecs", Destination: "10.0.0.0/24", NextHop: "server-id"}},
			Mod: []routetables.RouteOpts{{Type: "ecs", Destination: "10.1.0.0/24", NextHop: "server-id"}},
			Del: []routetables.DeleteRouteOpts{{Destination: "10.2.0.0/24"}},
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "route-table-id", actual.ID)
}

func TestUpdateInvalidID(t *testing.T) {
	actual, err := routetables.Update(serviceClient(), "../vpcs/vpc-id", routetables.UpdateOpts{Name: "updated"})
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
