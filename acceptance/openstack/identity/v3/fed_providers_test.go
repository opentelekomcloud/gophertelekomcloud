package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation/providers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFederatedProviderLifecycle(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Requires iam:identityProviders:createIdentityProvider permission")
	}

	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	cOpts := providers.CreateOpts{
		ID:          tools.RandomString("test-", 5),
		Description: tools.RandomString("This is ", 30),
		Enabled:     true,
	}

	provider, err := providers.Create(client, cOpts).Extract()
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = providers.Delete(client, provider.ID).ExtractErr()
		th.AssertNoErr(t, err)
	})

	th.AssertEquals(t, cOpts.Enabled, provider.Enabled)

	got, err := providers.Get(client, provider.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, provider, got)

	pages, err := providers.List(client).AllPages()
	th.AssertNoErr(t, err)

	providerList, err := providers.ExtractProviders(pages)
	th.AssertNoErr(t, err)
	found := false
	for _, p := range providerList {
		if p.ID == provider.ID {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("created provider not found in the list")
	}

	iFalse := false
	uOpts := providers.UpdateOpts{
		Enabled: &iFalse,
	}
	updated, err := providers.Update(client, provider.ID, uOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, false, updated.Enabled)

	got2, err := providers.Get(client, provider.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, updated, got2)
}
