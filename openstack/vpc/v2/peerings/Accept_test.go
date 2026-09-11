package peerings_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/peerings"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAccept(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/vpc/peerings/peering-id/accept", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		_, _ = w.Write([]byte(peeringJSON))
	})
	actual, err := peerings.Accept(serviceClient(), "peering-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "peering-id", actual.ID)
	th.AssertEquals(t, "ACTIVE", actual.Status)
}

func TestAcceptInvalidID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	actual, err := peerings.Accept(serviceClient(), "../peerings/other-id")
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and invalid ID error, got %#v, %v", actual, err)
	}
}
