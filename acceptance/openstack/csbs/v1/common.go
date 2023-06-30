package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
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
