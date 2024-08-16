package v3

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/backups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDdsBackupLifeCycle(t *testing.T) {
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
		delJob, err := backups.Delete(client, backup.BackupId)
		err = waitForJobCompleted(client, 600, delJob.JobId)
		th.AssertNoErr(t, err)
		t.Logf("Deleted DDSv3 backup: %s", backup.BackupId)
	})

	t.Logf("Attempting to list DDSv3 backups for instance: %s", instanceId)
	backupList, err := backups.List(client, backups.ListBackupsOpts{
		InstanceId: instanceId,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, backupList.TotalCount, 2)

	t.Logf("Attempting to get download links for DDSv3 backup: %s", backup.BackupId)
	links, err := backups.ListBackupDownloadLinks(client, backups.BackupLinkOpts{
		InstanceId: instanceId,
		BackupId:   backup.BackupId,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(links.Files), 1)

	t.Logf("Attempting to set backup policy for DDSv3 instance: %s", instanceId)
	err = backups.SetBackupPolicy(client, backups.ModifyBackupPolicyOpts{
		InstanceId: instanceId,
		BackupPolicy: &instances.BackupStrategy{
			StartTime: "01:00-02:00",
			KeepDays:  pointerto.Int(365),
			Period:    "1,2,3,5,6,7",
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to get backup policy for DDSv3 instance: %s", instanceId)
	getPolicy, err := backups.GetBackupPolicy(client, instanceId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "01:00-02:00", getPolicy.StartTime)
	th.AssertEquals(t, 365, *getPolicy.KeepDays)
	th.AssertEquals(t, "1,2,3,5,6,7", getPolicy.Period)

	t.Logf("Attempting to modify backup policy for DDSv3 instance: %s", instanceId)
	err = backups.SetBackupPolicy(client, backups.ModifyBackupPolicyOpts{
		InstanceId: instanceId,
		BackupPolicy: &instances.BackupStrategy{
			StartTime: "03:00-04:00",
			KeepDays:  pointerto.Int(7),
			Period:    "1,2,3",
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to get again backup policy for DDSv3 instance: %s", instanceId)
	getPolicyModified, err := backups.GetBackupPolicy(client, instanceId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "03:00-04:00", getPolicyModified.StartTime)
	th.AssertEquals(t, 7, *getPolicyModified.KeepDays)
	th.AssertEquals(t, "1,2,3", getPolicyModified.Period)
}

func waitForBackupAvailable(client *golangsdk.ServiceClient, secs int, backupId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		listOpts := backups.ListBackupsOpts{
			BackupId: backupId,
		}
		backupList, err := backups.List(client, listOpts)
		if err != nil {
			return false, err
		}
		if backupList.TotalCount == 1 {
			b := backupList.Backups
			if len(b) == 1 && b[0].Status == "COMPLETED" {
				return true, nil
			}
			return false, nil
		}
		return false, nil
	})
}
