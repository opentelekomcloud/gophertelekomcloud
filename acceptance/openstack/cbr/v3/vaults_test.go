package v3

import (
	"fmt"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	volumesV3 "github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v3/volumes"
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

func createVolume(t *testing.T) *volumesV3.Volume {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)
	vol, err := volumesV3.Create(client, volumesV3.CreateOpts{
		Size:       10,
		VolumeType: "SSD",
	}).Extract()
	th.AssertNoErr(t, err)

	err = golangsdk.WaitFor(300, func() (bool, error) {
		volume, err := volumesV3.Get(client, vol.ID).Extract()
		if err != nil {
			return false, err
		}
		if volume.Status == "available" {
			return true, nil
		}
		if volume.Status == "error" {
			return false, fmt.Errorf("error creating a volume")
		}
		return false, nil
	})
	th.AssertNoErr(t, err)

	return vol
}

func removeVolume(t *testing.T, id string) {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, volumesV3.Delete(client, id).ExtractErr())
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
	volume := createVolume(t)
	defer removeVolume(t, volume.ID)

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
