package peerings_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/peerings"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/vpc/peerings", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"peering":{"name":"peering-test","description":"peering description","request_vpc_info":{"vpc_id":"request-vpc-id"},"accept_vpc_info":{"vpc_id":"accept-vpc-id","tenant_id":"accept-tenant-id"}}}`)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"peering":` + peeringJSON + `}`))
	})
	actual, err := peerings.Create(serviceClient(), peerings.CreateOpts{
		Name:        "peering-test",
		Description: "peering description",
		RequestVpcInfo: peerings.VpcInfo{
			VpcID: "request-vpc-id",
		},
		AcceptVpcInfo: peerings.VpcInfo{
			VpcID:    "accept-vpc-id",
			TenantID: "accept-tenant-id",
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "peering-id", actual.ID)
	th.AssertEquals(t, "request-tenant-id", actual.RequestVpcInfo.TenantID)
	th.AssertEquals(t, "2026-09-11T10:00:00", actual.CreatedAt)
}

func TestCreateMissingRequiredOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	actual, err := peerings.Create(serviceClient(), peerings.CreateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and validation error, got %#v, %v", actual, err)
	}
}

func TestCreateError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/vpc/peerings", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	actual, err := peerings.Create(serviceClient(), peerings.CreateOpts{
		Name:           "peering-test",
		RequestVpcInfo: peerings.VpcInfo{VpcID: "request-vpc-id"},
		AcceptVpcInfo:  peerings.VpcInfo{VpcID: "accept-vpc-id"},
	})
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and request error, got %#v, %v", actual, err)
	}
}
