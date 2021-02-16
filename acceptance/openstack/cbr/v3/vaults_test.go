package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/policies"
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

func TestVaultResources(t *testing.T) {
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

	opts := vaults.CreateOpts{
		Billing: &vaults.BillingCreate{
			ConsistentLevel: "crash_consistent",
			ObjectType:      "disk",
			ProtectType:     "backup",
			Size:            100,
		},
		Description: "gophertelemocloud testing vault",
		Name:        tools.RandomString("cbr-test-", 5),
		Resources:   []vaults.ResourceCreate{},
	}
	vault, err := vaults.Create(client, opts).Extract()
	th.AssertNoErr(t, err)

	defer func() {
		th.AssertNoErr(t, vaults.Delete(client, vault.ID).ExtractErr())
	}()

	resourceType := "OS::Cinder::Volume"
	volume := openstack.CreateVolume(t)
	defer openstack.DeleteVolume(t, volume.ID)

	aOpts := vaults.AssociateResourcesOpts{
		Resources: []vaults.ResourceCreate{
			{
				ID:   volume.ID,
				Type: resourceType,
				Name: "cbr-vault-test-volume",
			},
		},
	}
	associated, err := vaults.AssociateResources(client, vault.ID, aOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(associated))
	th.AssertEquals(t, volume.ID, associated[0])

	dOpts := vaults.DissociateResourcesOpts{ResourceIDs: associated}
	dissociated, err := vaults.DissociateResources(client, vault.ID, dOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, associated, dissociated)
}

func TestVaultPolicy(t *testing.T) {
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

	opts := vaults.CreateOpts{
		Billing: &vaults.BillingCreate{
			ConsistentLevel: "crash_consistent",
			ObjectType:      "disk",
			ProtectType:     "backup",
			Size:            100,
		},
		Description: "gophertelemocloud testing vault",
		Name:        tools.RandomString("cbr-test-", 5),
		Resources:   []vaults.ResourceCreate{},
	}
	vault, err := vaults.Create(client, opts).Extract()
	th.AssertNoErr(t, err)

	defer func() {
		th.AssertNoErr(t, vaults.Delete(client, vault.ID).ExtractErr())
	}()

	iTrue := true
	policy, err := policies.Create(client, policies.CreateOpts{
		Name: "test-vault-policy",
		OperationDefinition: &policies.PolicyODCreate{
			DailyBackups: 1,
			WeekBackups:  2,
			YearBackups:  3,
			MonthBackups: 4,
			MaxBackups:   10,
			Timezone:     "UTC+03:00",
		},
		Trigger: &policies.Trigger{
			Properties: policies.TriggerProperties{
				Pattern: []string{"FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00"},
			},
		},
		Enabled:       &iTrue,
		OperationType: "backup",
	}).Extract()
	th.AssertNoErr(t, err)

	defer func() {
		th.AssertNoErr(t, policies.Delete(client, policy.ID).ExtractErr())
	}()

	bind, err := vaults.BindPolicy(client, vault.ID, vaults.BindPolicyOpts{PolicyID: policy.ID}).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vault.ID, bind.VaultID)
	th.AssertEquals(t, policy.ID, bind.PolicyID)

	unbind, err := vaults.UnbindPolicy(client, vault.ID, vaults.BindPolicyOpts{PolicyID: policy.ID}).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vault.ID, unbind.VaultID)
	th.AssertEquals(t, policy.ID, unbind.PolicyID)
}
