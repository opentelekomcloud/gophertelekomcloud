package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v3/snapshots"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vbs/v2/backups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVBSV2BackupLifecycle(t *testing.T) {
	client, err := clients.NewVbsV2Client()
	th.AssertNoErr(t, err)

	volume := openstack.CreateVolume(t)
	defer func() {
		deleteVolumeSnapshots(t, volume.ID)
		openstack.DeleteVolume(t, volume.ID)
	}()

	opts := backups.CreateOpts{
		VolumeId:    volume.ID,
		Name:        tools.RandomString("vbs-instance-", 6),
		Description: tools.RandomString("Here: ", 20),
	}
	job, err := backups.Create(client, opts).ExtractJobResponse()
	th.AssertNoErr(t, err)

	if err := backups.WaitForJobSuccess(client, 600, job.JobID); err != nil {
		t.Fatalf("error waiting for backup to be created: %s", err)
	}
	t.Log("backup successfully created")

	v, err := backups.GetJobEntity(client, job.JobID, "backup_id")
	th.AssertNoErr(t, err)
	backupID := v.(string)

	defer func() {
		err = backups.Delete(client, backupID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Log("backup successfully deleted")
	}()

	backupDetails, err := backups.Get(client, backupID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, opts.Description, backupDetails.Description)
	th.AssertEquals(t, opts.Name, backupDetails.Name)
}

func deleteSnapshot(client *golangsdk.ServiceClient, id string) error {
	err := snapshots.WaitForStatus(client, id, "available", 600)
	if err != nil {
		return err
	}

	err = snapshots.Delete(client, id).ExtractErr()
	if err != nil {
		return err
	}
	return nil
}

func deleteVolumeSnapshots(t *testing.T, volumeID string) {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	allPages, err := snapshots.List(client, snapshots.ListOpts{
		VolumeID: volumeID,
	}).AllPages()
	th.AssertNoErr(t, err)

	snapshotSlice, err := snapshots.ExtractSnapshots(allPages)
	th.AssertNoErr(t, err)

	for _, s := range snapshotSlice {
		th.AssertNoErr(t, deleteSnapshot(client, s.ID))
	}
}
