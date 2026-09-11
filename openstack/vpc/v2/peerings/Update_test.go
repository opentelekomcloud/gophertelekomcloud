package peerings_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/peerings"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/vpc/peerings/peering-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{"peering":{"name":"updated-name","description":""}}`)
		_, _ = w.Write([]byte(`{"peering":` + peeringJSON + `}`))
	})
	actual, err := peerings.Update(serviceClient(), "peering-id", peerings.UpdateOpts{
		Name:        pointerto.String("updated-name"),
		Description: pointerto.String(""),
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "peering-id", actual.ID)
}

func TestUpdateMissingAttributes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	actual, err := peerings.Update(serviceClient(), "peering-id", peerings.UpdateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and validation error, got %#v, %v", actual, err)
	}
}

func TestUpdateInvalidID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	actual, err := peerings.Update(serviceClient(), "../peerings/other-id", peerings.UpdateOpts{
		Name: pointerto.String("updated-name"),
	})
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and invalid ID error, got %#v, %v", actual, err)
	}
}
