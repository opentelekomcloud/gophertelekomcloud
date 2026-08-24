package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v2/volumes"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v2/cloudvolumes"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v2/snapshots"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEVSv2List(t *testing.T) {
	client, err := clients.NewBlockStorageV2Client()
	th.AssertNoErr(t, err)

	createOpts := volumes.CreateOpts{
		Size:       40,
		Name:       tools.RandomString("tf-evs-disk-", 4),
		VolumeType: "SSD",
	}

	resp, err := volumes.Create(client, createOpts)
	th.AssertNoErr(t, err)

	err = waitForEvsAvailable(client, 100, resp.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = volumes.Delete(client, resp.ID, volumes.DeleteOpts{})
		th.AssertNoErr(t, err)
	})

	list, err := cloudvolumes.List(client, cloudvolumes.ListOpts{
		ID: resp.ID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, list)
}

func TestEVSv2SnapshotWorkflow(t *testing.T) {
	client, err := clients.NewBlockStorageV2Client()
	th.AssertNoErr(t, err)

	volumeName := tools.RandomString("gopher-evs-disk-", 4)
	createVolumeOpts := volumes.CreateOpts{
		Size:       10,
		Name:       volumeName,
		VolumeType: "SSD",
	}

	volume, err := volumes.Create(client, createVolumeOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = volumes.Delete(client, volume.ID, volumes.DeleteOpts{})
		th.AssertNoErr(t, err)
	})

	err = waitForEvsAvailable(client, 100, volume.ID)
	th.AssertNoErr(t, err)

	snapshotName := tools.RandomString("tf-evs-snapshot-", 4)
	snapshot, err := snapshots.Create(client, snapshots.CreateOpts{
		VolumeID:    volume.ID,
		Name:        snapshotName,
		Description: "EVS v2 acceptance test snapshot",
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = snapshots.Delete(client, snapshot.ID)
		th.AssertNoErr(t, err)
		err = waitForSnapshotDeleted(client, 300, snapshot.ID)
		th.AssertNoErr(t, err)
	})

	err = snapshots.WaitForStatus(client, snapshot.ID, "available", 300)
	th.AssertNoErr(t, err)

	updatedSnapshotName := tools.RandomString("tf-evs-snapshot-updated-", 4)
	updatedSnapshot, err := snapshots.Update(client, snapshot.ID, snapshots.UpdateOpts{
		Name:        updatedSnapshotName,
		Description: "EVS v2 acceptance test snapshot updated",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updatedSnapshot.Name, updatedSnapshotName)
	th.AssertEquals(t, updatedSnapshot.Description, "EVS v2 acceptance test snapshot updated")

	gotSnapshot, err := snapshots.Get(client, snapshot.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, gotSnapshot.ID, snapshot.ID)
	th.AssertEquals(t, gotSnapshot.VolumeID, volume.ID)

	list, err := snapshots.List(client, snapshots.ListOpts{
		ID:       snapshot.ID,
		VolumeID: volume.ID,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(list.Snapshots), 1)
	th.AssertEquals(t, list.Snapshots[0].ID, snapshot.ID)

	rollback, err := snapshots.Rollback(client, snapshot.ID, snapshots.RollbackOpts{
		VolumeID: volume.ID,
		Name:     volumeName,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, rollback.VolumeID, volume.ID)

	err = waitForEvsAvailable(client, 300, volume.ID)
	th.AssertNoErr(t, err)
}

func waitForEvsAvailable(client *golangsdk.ServiceClient, secs int, volId string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		vol, err := volumes.Get(client, volId)
		if err != nil {
			return false, err
		}

		if vol.Status == "available" {
			return true, nil
		}
		return false, nil
	})
}

func waitForSnapshotDeleted(client *golangsdk.ServiceClient, secs int, snapshotID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := snapshots.Get(client, snapshotID)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}
		return false, nil
	})
}
