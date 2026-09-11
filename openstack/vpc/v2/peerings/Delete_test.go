package peerings_test

import (
	"net/http"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/peerings"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	return client.ServiceClient()
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/vpc/peerings/peering-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})
	err := peerings.Delete(serviceClient(), "peering-id")
	th.AssertNoErr(t, err)
}

func TestDeleteInvalidID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	if err := peerings.Delete(serviceClient(), "../peerings/other-id"); err == nil {
		t.Fatal("expected invalid peering ID error")
	}
}
