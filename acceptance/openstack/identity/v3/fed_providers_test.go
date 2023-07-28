package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation/mappings"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation/protocols"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation/providers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFederatedProviderLifecycle(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Requires iam:identityProviders:createIdentityProvider permission")
	}

	client, err := clients.NewIdentityV3AdminClient()
	th.AssertNoErr(t, err)

	cOpts := providers.CreateOpts{
		ID:          tools.RandomString("test-", 5),
		Description: tools.RandomString("This is ", 30),
		Enabled:     true,
	}

	provider, err := providers.Create(client, cOpts).Extract()
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = providers.Delete(client, provider.ID).Err
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

	mappingCreateOpts := mappings.CreateOpts{
		Rules: []mappings.RuleOpts{
			{
				Local: []mappings.LocalRuleOpts{
					{
						User: &mappings.UserOpts{
							Name: "{0}",
						},
					},
					{
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

	mappingName := tools.RandomString("muh", 3)
	mapping, err := mappings.Create(client, mappingName, mappingCreateOpts).Extract()
	th.AssertNoErr(t, err)

	protocolsCreateOpts := protocols.CreateOpts{
		MappingID: mapping.ID,
	}

	_, err = protocols.Create(client, provider.ID, "oidc", protocolsCreateOpts).Extract()
	th.AssertNoErr(t, err)

	nClient, err := clients.NewIdentityV30AdminClient()
	th.AssertNoErr(t, err)

	signingKey := "{\"keys\":[{\"kty\":\"RSA\",\"e\":\"AQAB\",\"use\":\"sig\",\"n\":\"example\",\"kid\":\"kid_example\",\"alg\":\"RS256\"}]}"

	oidc, err := providers.CreateOIDC(nClient, providers.CreateOIDCOpts{
		IdpIp:      provider.ID,
		AccessMode: "program",
		IdpUrl:     "https://accounts.example.com",
		ClientId:   "client_id_example",
		SigningKey: signingKey,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, oidc.AccessMode, "program")

	updatedOidc, err := providers.UpdateOIDC(nClient, providers.UpdateOIDCOpts{
		IdpIp:    provider.ID,
		ClientId: "new_client_id",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updatedOidc.ClientId, "new_client_id")

	getOIDC, err := providers.GetOIDC(nClient, provider.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, getOIDC)
}
