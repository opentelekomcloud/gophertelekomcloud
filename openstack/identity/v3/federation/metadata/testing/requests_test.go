package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation/metadata"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestImportMetadataRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(metadataURL, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, importMetadataBody)

		w.WriteHeader(201)
		_, _ = fmt.Fprint(w, `{ "message": "Import metadata successful"}`)
	})

	opts := metadata.ImportOpts{
		DomainID: domainID,
		Metadata: data,
	}

	err := metadata.Import(fake.ServiceClient(), "ACME", "saml", opts).Err
	th.AssertNoErr(t, err)
}

func TestGetMetadataRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(metadataURL, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		w.WriteHeader(200)
		_, _ = fmt.Fprint(w, getMetadataResponseBody)
	})

	meta, err := metadata.Get(fake.ServiceClient(), "ACME", "saml").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, domainID, meta.DomainID)
	th.AssertEquals(t, data, meta.Data)
}
