package peerings_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/peerings"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/vpc/peerings/peering-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"peering":` + peeringJSON + `}`))
	})
	actual, err := peerings.Get(serviceClient(), "peering-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "peering-id", actual.ID)
	th.AssertEquals(t, "accept-vpc-id", actual.AcceptVpcInfo.VpcID)
	th.AssertEquals(t, "2026-09-11T10:01:00", actual.UpdatedAt)
}

func TestGetInvalidID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	actual, err := peerings.Get(serviceClient(), "../peerings/other-id")
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and invalid ID error, got %#v, %v", actual, err)
	}
}
