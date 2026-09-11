package peerings_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/peerings"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/vpc/peerings", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		query := r.URL.Query()
		th.AssertEquals(t, "peering-id", query.Get("id"))
		th.AssertEquals(t, "peering-test", query.Get("name"))
		th.AssertEquals(t, "ACTIVE", query.Get("status"))
		th.AssertEquals(t, "tenant-id", query.Get("tenant_id"))
		th.AssertEquals(t, "vpc-id", query.Get("vpc_id"))
		th.AssertEquals(t, "marker-id", query.Get("marker"))
		th.AssertEquals(t, "10", query.Get("limit"))
		_, _ = w.Write([]byte(`{"peerings":[` + peeringJSON + `]}`))
	})
	limit := 10
	actual, err := peerings.List(serviceClient(), peerings.ListOpts{
		ID:       "peering-id",
		Name:     "peering-test",
		Status:   "ACTIVE",
		TenantID: "tenant-id",
		VpcID:    "vpc-id",
		Marker:   "marker-id",
		Limit:    &limit,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(actual))
	th.AssertEquals(t, "peering-id", actual[0].ID)
}

func TestListPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	requests := 0
	th.Mux.HandleFunc("/vpc/peerings", func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.URL.Query().Get("marker") == "" {
			_, _ = w.Write([]byte(`{
				"peerings":[` + peeringJSON + `],
				"peerings_links":[{"href":"` + th.Server.URL + `/vpc/peerings?limit=1&marker=peering-id","rel":"next"}]
			}`))
			return
		}
		th.AssertEquals(t, "peering-id", r.URL.Query().Get("marker"))
		_, _ = w.Write([]byte(`{"peerings":[` + peeringJSON + `],"peerings_links":[]}`))
	})
	limit := 1
	actual, err := peerings.List(serviceClient(), peerings.ListOpts{Limit: &limit})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(actual))
	th.AssertEquals(t, 3, requests)
}

func TestExtractPeeringsInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/vpc/peerings", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"peerings":`))
	})
	actual, err := peerings.List(serviceClient(), peerings.ListOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and extraction error, got %#v, %v", actual, err)
	}
}
