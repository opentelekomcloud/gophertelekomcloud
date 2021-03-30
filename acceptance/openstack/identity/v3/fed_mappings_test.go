package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation/mappings"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFederatedProviderLifecycle(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow iam:identityProviders:createMapping to be performed.")
	}

	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	createOpts := mappings.CreateOpts{
		Rules: []mappings.RuleOpts{
			{
				Local: []mappings.LocalRuleOpts{
					{
						User: mappings.UserOpts{
							Name: "samltestid",
						},
						Group: mappings.GroupOpts{
							Name: "power_user",
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

	mapping, err := mappings.Create(client, "ACME", createOpts).Extract()
	th.AssertNoErr(t, err)

	defer func() {
		err = mappings.Delete(client, mapping.ID).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	th.AssertEquals(t, createOpts.Rules[0].Local[0].User.Name, mapping.Rules[0].Local[0].User.Name)

	got, err := mappings.Get(client, mapping.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, mapping, got)

	pages, err := mappings.List(client).AllPages()
	th.AssertNoErr(t, err)

	mappingList, err := mappings.ExtractMappings(pages)
	th.AssertNoErr(t, err)
	found := false
	for _, m := range mappingList {
		if m.ID == mapping.ID {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("created mapping not found in the list")
	}

	updateOpts := mappings.UpdateOpts{
		Rules: []mappings.RuleOpts{
			{
				Local: []mappings.LocalRuleOpts{
					{
						User: mappings.UserOpts{
							Name: "samltestid-{0}",
						},
						Group: mappings.GroupOpts{
							Name: "power_user",
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
	updated, err := mappings.Update(client, mapping.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	// th.AssertEquals(t, false, updated.Enabled)

	got2, err := mappings.Get(client, mapping.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, updated, got2)
}
