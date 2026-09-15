package routetables_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/routetables"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/routetables/route-table-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"routetable":` + routeTableJSON + `}`))
	})
	actual, err := routetables.Get(serviceClient(), "route-table-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "vpc-id", actual.VpcID)
	th.AssertEquals(t, "10.10.10.0/24", actual.Routes[0].DestinationCIDR)
}

func TestGetInvalidID(t *testing.T) {
	actual, err := routetables.Get(serviceClient(), "../vpcs/vpc-id")
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
