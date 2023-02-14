package v3

import (
	"fmt"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/vaults"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/backups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/checkpoint"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateCheckpoint(t *testing.T, client *golangsdk.ServiceClient, createOpts checkpoint.CreateOpts) *checkpoint.Checkpoint {
	backup, err := checkpoint.Create(client, createOpts)
	th.AssertNoErr(t, err)

	err = golangsdk.WaitFor(600, func() (bool, error) {
		checkp, err := checkpoint.Get(client, backup.ID)
		if err != nil {
			return false, err
		}
		if checkp.Status == "available" {
			return true, nil
		}
		if checkp.Status == "error" {
			return false, fmt.Errorf("error creating a checkpoint")
		}
		return false, nil
	})
	th.AssertNoErr(t, err)

	return backup
}

func RestoreBackup(t *testing.T, client *golangsdk.ServiceClient, id string, opts backups.RestoreBackupOpts) error {
	errRest := backups.RestoreBackup(client, id, opts)
	th.AssertNoErr(t, errRest)

	err := golangsdk.WaitFor(600, func() (bool, error) {
		back, err := backups.Get(client, id)
		if err != nil {
			return false, err
		}
		if back.Status == "available" {
			return true, nil
		}
		if back.Status == "error" {
			return false, fmt.Errorf("error restoring a backup")
		}
		return false, nil
	})
	th.AssertNoErr(t, err)

	return nil
}

func waitForBackupDelete(client *golangsdk.ServiceClient, secs int, id string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := backups.Get(client, id)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}

func CreateCBR(t *testing.T, client *golangsdk.ServiceClient) (*vaults.Vault, vaults.AssociateResourcesOpts, checkpoint.CreateOpts, *checkpoint.Checkpoint) {
	// Create Vault for further backup
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
	vault, err := vaults.Create(client, opts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, vaults.Delete(client, vault.ID))
	})

	// Create Volume
	volume := openstack.CreateVolume(t)
	t.Cleanup(func() {
		openstack.DeleteVolume(t, volume.ID)
	})

	// Associate server to the vault
	aOpts := vaults.AssociateResourcesOpts{
		Resources: []vaults.ResourceCreate{
			{
				ID:   volume.ID,
				Type: "OS::Cinder::Volume",
			},
		},
	}
	associated, err := vaults.AssociateResources(client, vault.ID, aOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(associated))

	// Create vault checkpoint
	optsVault := checkpoint.CreateOpts{
		VaultID: vault.ID,
		Parameters: checkpoint.CheckpointParam{
			Description: "go created backup",
			Incremental: true,
			Name:        tools.RandomString("go-checkpoint", 5),
		},
	}
	checkp := CreateCheckpoint(t, client, optsVault)
	return vault, aOpts, optsVault, checkp
}
