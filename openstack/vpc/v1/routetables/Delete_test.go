package routetables_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/routetables"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/project-id/routetables/route-table-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})
	th.AssertNoErr(t, routetables.Delete(serviceClient(), "route-table-id"))
}

func TestDeleteInvalidID(t *testing.T) {
	if err := routetables.Delete(serviceClient(), "../vpcs/vpc-id"); err == nil {
		t.Fatal("expected invalid ID error")
	}
}
