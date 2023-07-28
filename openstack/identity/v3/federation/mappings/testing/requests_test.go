package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation/mappings"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestMappingCreateRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(mappingURI, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		w.WriteHeader(201)
		_, _ = fmt.Fprint(w, mappingResponse)
	})

	opts := mappings.CreateOpts{
		Rules: []mappings.RuleOpts{
			{
				Local: []mappings.LocalRuleOpts{
					{
						User: &mappings.UserOpts{
							Name: "{0}",
						},
						Groups: "[\"admin\",\"manager\"]",
					},
				},
				Remote: []mappings.RemoteRuleOpts{
					{
						Type: "uid",
					},
				},
			},
		},
	}
	p, err := mappings.Create(fake.ServiceClient(), mappingID, opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(opts.Rules), len(p.Rules))
}

func TestMappingGetRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(mappingURI, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		_, _ = fmt.Fprint(w, mappingResponse)
	})

	p, err := mappings.Get(fake.ServiceClient(), mappingID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, mappingID, p.ID)
}

func TestMappingListRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(listURI, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, mappingListResponse)
	})

	pages, err := mappings.List(fake.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)

	mappingList, err := mappings.ExtractMappings(pages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(mappingList))
}

func TestMappingUpdateRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(mappingURI, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, updatedMappingResponse)
	})

	opts := mappings.UpdateOpts{
		Rules: []mappings.RuleOpts{
			{
				Local: []mappings.LocalRuleOpts{
					{
						User: &mappings.UserOpts{
							Name: "samltestid-{0}",
						},
					},
				},
				Remote: []mappings.RemoteRuleOpts{
					{
						Type: "uid",
					},
				},
			},
		},
	}
	p, err := mappings.Update(fake.ServiceClient(), mappingID, opts).Extract()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(opts.Rules), len(p.Rules))
}

func TestMappingDeleteRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(mappingURI, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(204)
	})

	err := mappings.Delete(fake.ServiceClient(), mappingID).Err
	th.AssertNoErr(t, err)
}
