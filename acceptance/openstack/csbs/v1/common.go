package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cbr/v3"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/checkpoint"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/vaults"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/csbs/v1/backup"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateCSBS(t *testing.T, client *golangsdk.ServiceClient, ecsId string) *backup.Checkpoint {
	backupName := tools.RandomString("backup-", 3)

	t.Logf("Attempting to create CSBS backup")
	checkpoint, err := backup.Create(client, ecsId, backup.CreateOpts{
		BackupName:   backupName,
		Description:  "bla-bla",
		ResourceType: "OS::Nova::Server",
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete CSBS backup: %s", checkpoint.Id)
		err = backup.Delete(client, checkpoint.Id)
		th.AssertNoErr(t, err)

		err = waitForBackupDeleted(client, 600, checkpoint.Id)
		th.AssertNoErr(t, err)
		t.Logf("Deleted CSBS backup: %s", checkpoint.Id)
	})

	err = waitForBackupCreated(client, 600, checkpoint.Id)
	th.AssertNoErr(t, err)

	return checkpoint
}

func waitForBackupCreated(client *golangsdk.ServiceClient, secs int, backupID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		csbsBackup, err := backup.Get(client, backupID)
		if err != nil {
			return false, err
		}

		if csbsBackup.Id == "error" {
			return false, err
		}

		if csbsBackup.Status == "available" {
			return true, nil
		}

		return false, nil
	})
}

func waitForBackupDeleted(client *golangsdk.ServiceClient, secs int, backupID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := backup.Get(client, backupID)
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
	checkp := v3.CreateCheckpoint(t, client, optsVault)
	return vault, aOpts, optsVault, checkp
}
