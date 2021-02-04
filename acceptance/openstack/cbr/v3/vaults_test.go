package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/vaults"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVaultLifecycle(t *testing.T) {
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

	opts := vaults.CreateOpts{
		Billing: &vaults.BillingCreate{
			ConsistentLevel: "crash_consistent",
			ObjectType:      "server",
			ProtectType:     "backup",
			Size:            100,
		},
		Description: "gophertelemocloud testing vault",
		Name:        tools.RandomString("cbr-test-", 5),
		Resources:   []vaults.ResourceCreate{},
	}
	vault, err := vaults.Create(client, opts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, opts.Billing.ConsistentLevel, vault.Billing.ConsistentLevel)
	th.AssertEquals(t, opts.Billing.ObjectType, vault.Billing.ObjectType)
	th.AssertEquals(t, opts.Billing.ProtectType, vault.Billing.ProtectType)
	th.AssertEquals(t, opts.Billing.Size, vault.Billing.Size)
	th.AssertEquals(t, opts.Name, vault.Name)
	th.AssertEquals(t, opts.Description, vault.Description)

	defer func() {
		th.AssertNoErr(t, vaults.Delete(client, vault.ID).ExtractErr())
	}()

	getVault, err := vaults.Get(client, vault.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, vault, getVault)

	updateOpts := vaults.UpdateOpts{
		Name: tools.RandomString("cbr-test-2-", 5),
	}
	updated, err := vaults.Update(client, vault.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vault.ID, updated.ID)
	th.AssertEquals(t, updateOpts.Name, updated.Name)

	getUpdated, err := vaults.Get(client, vault.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, updated, getUpdated)
}
