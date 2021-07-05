package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation/protocols"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestProtocolCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(protocolURL, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, createRequestBody)
		w.WriteHeader(201)
		_, _ = fmt.Fprint(w, getResponseBody)
	})

	protocol, err := protocols.Create(fake.ServiceClient(), providerID, protocolID, protocols.CreateOpts{
		MappingID: mappingID,
	}).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, mappingID, protocol.MappingID)
}

func TestProtocolList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(listURL, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, listResponseBody)
	})

	pages, err := protocols.List(fake.ServiceClient(), providerID).AllPages()
	th.AssertNoErr(t, err)
	prots, err := protocols.ExtractProtocols(pages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(prots))
}

func TestProtocolGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(protocolURL, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, getResponseBody)
	})

	protocol, err := protocols.Get(fake.ServiceClient(), providerID, protocolID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, mappingID, protocol.MappingID)
}

func TestProtocolUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(protocolURL, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, getResponseBody)
	})

	opts := protocols.UpdateOpts{
		MappingID: "SAML2",
	}
	protocol, err := protocols.Update(fake.ServiceClient(), providerID, protocolID, opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, mappingID, protocol.MappingID)
}

func TestProtocolDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(protocolURL, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		w.WriteHeader(204)
	})
}
