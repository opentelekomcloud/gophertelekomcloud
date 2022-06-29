package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/backups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/checkpoint"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/vaults"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBackupLifecycle(t *testing.T) {
	clientCbr, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

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
	vault, err := vaults.Create(clientCbr, opts).Extract()
	th.AssertNoErr(t, err)
	defer func() {
		th.AssertNoErr(t, vaults.Delete(clientCbr, vault.ID).ExtractErr())
	}()

	// Create Volume
	volume := openstack.CreateVolume(t)
	defer openstack.DeleteVolume(t, volume.ID)

	// Associate server to the vault
	resourceType := "OS::Cinder::Volume"
	aOpts := vaults.AssociateResourcesOpts{
		Resources: []vaults.ResourceCreate{
			{
				ID:   volume.ID,
				Type: resourceType,
			},
		},
	}
	associated, err := vaults.AssociateResources(clientCbr, vault.ID, aOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(associated))

	// Create vault checkpoint
	description := "go created backup"
	checkName := tools.RandomString("go-checkpoint", 5)
	optsVault := checkpoint.CreateOpts{
		VaultId: vault.ID,
		Parameters: checkpoint.CheckpointParam{
			Description: description,
			Incremental: true,
			Name:        checkName,
		},
	}
	checkp := CreateChekpoint(t, clientCbr, optsVault)
	th.AssertEquals(t, vault.ID, checkp.Vault.Id)
	th.AssertEquals(t, description, checkp.ExtraInfo.Description)
	th.AssertEquals(t, checkName, checkp.ExtraInfo.Name)

	checkpointGet, err := checkpoint.Get(clientCbr, checkp.Id).Extract()
	th.AssertNoErr(t, err)
	// Checks are disabled due to STO-10008 bug
	// th.AssertEquals(t, description, checkpointGet.ExtraInfo.Description)
	// th.AssertEquals(t, checkName, checkpointGet.ExtraInfo.Name)
	th.AssertEquals(t, "available", checkpointGet.Status)
	th.AssertEquals(t, vault.ID, checkpointGet.Vault.Id)

	allPages, err := backups.List(clientCbr, backups.ListOpts{CheckpointId: checkp.Id}).AllPages()
	th.AssertNoErr(t, err)
	allBackups, err := backups.ExtractBackups(allPages)
	bOpts := backups.RestoreBackupOpts{
		Restore: backups.RestoreBackupStruct{
			VolumeId: allBackups[0].ResourceID,
		},
	}
	backupErr := RestoreBackup(t, clientCbr, allBackups[0].ID, bOpts)
	th.AssertNoErr(t, backupErr)
	deletePage := backups.Delete(clientCbr, allBackups[0].ID).ExtractErr()
	th.AssertNoErr(t, deletePage)
	waitForBackupDelete(clientCbr, 600, allBackups[0].ID)
}
