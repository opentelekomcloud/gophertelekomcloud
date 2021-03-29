package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation/providers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestProviderCreateResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(createURI, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		w.WriteHeader(201)
		_, _ = fmt.Fprintf(w, jsonResponse)
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
