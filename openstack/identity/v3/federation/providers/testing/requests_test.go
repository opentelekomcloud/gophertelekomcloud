package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation/providers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestProviderCreateRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(providerURI, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		w.WriteHeader(201)
		_, _ = fmt.Fprint(w, providerResponse)
	})

	opts := providers.CreateOpts{
		ID:          providerID,
		Description: providerDescription,
		Enabled:     true,
	}
	p, err := providers.Create(fake.ServiceClient(), opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, opts.ID, p.ID)
	th.AssertEquals(t, opts.Enabled, p.Enabled)
	th.AssertEquals(t, 0, len(p.RemoteIDs))
}

func TestProviderGetRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(providerURI, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		_, _ = fmt.Fprint(w, providerResponse)
	})

	p, err := providers.Get(fake.ServiceClient(), providerID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, p.Enabled)
	th.AssertEquals(t, providerDescription, p.Description)
}

func TestProviderListRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(listURI, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, providerListResponse)
	})

	pages, err := providers.List(fake.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)

	providerList, err := providers.ExtractProviders(pages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(providerList))
}

func TestProviderUpdateRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(providerURI, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, updatedProviderResposnse)
	})

	iFalse := false
	opts := providers.UpdateOpts{
		Enabled: &iFalse,
	}
	p, err := providers.Update(fake.ServiceClient(), providerID, opts).Extract()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, iFalse, p.Enabled)
}

func TestProviderDeleteRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(providerURI, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(204)
	})

	err := providers.Delete(fake.ServiceClient(), providerID).Err
	th.AssertNoErr(t, err)
}
