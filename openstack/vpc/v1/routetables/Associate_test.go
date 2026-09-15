package routetables_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/routetables"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAssociate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/routetables/route-table-id/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"routetable":{"subnets":{"associate":["network-id"]}}}`)
		_, _ = w.Write([]byte(`{"routetable":` + routeTableJSON + `}`))
	})
	actual, err := routetables.Associate(serviceClient(), "route-table-id", routetables.AssociateOpts{Subnets: []string{"network-id"}})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "network-id", actual.Subnets[0].ID)
}

func TestAssociateMissingSubnets(t *testing.T) {
	actual, err := routetables.Associate(serviceClient(), "route-table-id", routetables.AssociateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected validation error, got %#v, %v", actual, err)
	}
}
