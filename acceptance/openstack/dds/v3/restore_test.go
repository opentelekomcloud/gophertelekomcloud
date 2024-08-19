package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/backups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDdsRestoreLifeCycle(t *testing.T) {
	instanceId := os.Getenv("DDS_INSTANCE_ID")

	if instanceId == "" {
		t.Skip("`DDS_INSTANCE_ID` need to be defined")
	}

	client, err := clients.NewDdsV3Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create DDSv3 backup for instance: %s", instanceId)
	backupOpts := backups.CreateOpts{
		Backup: &backups.Backup{
			InstanceId:  instanceId,
			Name:        tools.RandomString("dmd-dds-backup-", 5),
			Description: "this is backup for dds",
		},
	}
	backup, err := backups.Create(client, backupOpts)
	th.AssertNoErr(t, err)
	err = waitForBackupAvailable(client, 600, backup.BackupId)
	th.AssertNoErr(t, err)
	t.Logf("DDSv3 backup successfully created")
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete DDSv3 backup: %s", backup.BackupId)
		delJob, errDel := backups.Delete(client, backup.BackupId)
		err = waitForJobCompleted(client, 600, delJob.JobId)
		th.AssertNoErr(t, errDel)
		t.Logf("Deleted DDSv3 backup: %s", backup.BackupId)
	})

	t.Logf("Attempting to restore DDSv3 instance to Original: %s", instanceId)
	restoreJob, err := instances.RestoreToOriginal(client, instances.RestoreToOriginalOpts{
		Source: instances.Source{
			InstanceId: instanceId,
			Type:       "backup",
			BackupId:   backup.BackupId,
		},
		Target: instances.Target{InstanceId: instanceId},
	})
	th.AssertNoErr(t, err)
	err = waitForJobCompleted(client, 600, *restoreJob)
	th.AssertNoErr(t, err)
}
