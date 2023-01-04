package v3

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/snapshots"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/volumes"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/volumetypes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// CreateSnapshot will create a snapshot of the specified volume.
// Snapshot will be assigned a random name and description.
func CreateSnapshot(t *testing.T, client *golangsdk.ServiceClient, volume *volumes.Volume) (*snapshots.Snapshot, error) {
	snapshotName := tools.RandomString("ACPTTEST", 16)
	snapshotDescription := tools.RandomString("ACPTTEST", 16)
	t.Logf("Attempting to create snapshot: %s", snapshotName)

	createOpts := snapshots.CreateOpts{
		VolumeID:    volume.ID,
		Name:        snapshotName,
		Description: snapshotDescription,
	}

	snapshot, err := snapshots.Create(client, createOpts)
	if err != nil {
		return snapshot, err
	}

	err = snapshots.WaitForStatus(client, snapshot.ID, "available", 60)
	if err != nil {
		return snapshot, err
	}

	tools.PrintResource(t, snapshot)
	th.AssertEquals(t, snapshot.Name, snapshotName)
	th.AssertEquals(t, snapshot.VolumeID, volume.ID)

	t.Logf("Successfully created snapshot: %s", snapshot.ID)

	return snapshot, nil
}

// CreateVolume will create a volume with a random name and size of 1GB. An
// error will be returned if the volume was unable to be created.
func CreateVolume(t *testing.T, client *golangsdk.ServiceClient) (*volumes.Volume, error) {
	volumeName := tools.RandomString("ACPTTEST", 16)
	volumeDescription := tools.RandomString("ACPTTEST-DESC", 16)
	t.Logf("Attempting to create volume: %s", volumeName)

	createOpts := volumes.CreateOpts{
		Size:             1,
		Name:             volumeName,
		Description:      volumeDescription,
		AvailabilityZone: clients.EnvOS.GetEnv("AVAILABILITY_ZONE"),
	}

	volume, err := volumes.Create(client, createOpts)
	if err != nil {
		return volume, err
	}

	err = volumes.WaitForStatus(client, volume.ID, "available", 60)
	if err != nil {
		return volume, err
	}

	tools.PrintResource(t, volume)
	th.AssertEquals(t, volume.Name, volumeName)
	th.AssertEquals(t, volume.Description, volumeDescription)
	th.AssertEquals(t, volume.Size, 1)

	t.Logf("Successfully created volume: %s", volume.ID)

	return volume, nil
}

// CreateVolumeWithType will create a volume of the given volume type
// with a random name and size of 1GB. An error will be returned if
// the volume was unable to be created.
func CreateVolumeWithType(t *testing.T, client *golangsdk.ServiceClient, vt *volumetypes.VolumeType) (*volumes.Volume, error) {
	volumeName := tools.RandomString("ACPTTEST", 16)
	volumeDescription := tools.RandomString("ACPTTEST-DESC", 16)
	t.Logf("Attempting to create volume: %s", volumeName)

	createOpts := volumes.CreateOpts{
		Size:        1,
		Name:        volumeName,
		Description: volumeDescription,
		VolumeType:  vt.Name,
	}

	volume, err := volumes.Create(client, createOpts)
	if err != nil {
		return volume, err
	}

	err = volumes.WaitForStatus(client, volume.ID, "available", 60)
	if err != nil {
		return volume, err
	}

	tools.PrintResource(t, volume)
	th.AssertEquals(t, volume.Name, volumeName)
	th.AssertEquals(t, volume.Description, volumeDescription)
	th.AssertEquals(t, volume.Size, 1)
	th.AssertEquals(t, volume.VolumeType, vt.Name)

	t.Logf("Successfully created volume: %s", volume.ID)

	return volume, nil
}

// DeleteSnapshot will delete a snapshot. A fatal error will occur if the
// snapshot failed to be deleted.
func DeleteSnapshot(t *testing.T, client *golangsdk.ServiceClient, snapshot *snapshots.Snapshot) {
	err := snapshots.Delete(client, snapshot.ID)
	if err != nil {
		t.Fatalf("Unable to delete snapshot %s: %+v", snapshot.ID, err)
	}

	// Volumes can't be deleted until their snapshots have been,
	// so block until the snapshot as been deleted.
	err = tools.WaitFor(func() (bool, error) {
		_, err := snapshots.Get(client, snapshot.ID)
		if err != nil {
			return true, nil
		}

		return false, nil
	})
	if err != nil {
		t.Fatalf("Error waiting for snapshot to delete: %v", err)
	}

	t.Logf("Deleted snapshot: %s", snapshot.ID)
}

// DeleteVolume will delete a volume. A fatal error will occur if the volume
// failed to be deleted. This works best when used as a deferred function.
func DeleteVolume(t *testing.T, client *golangsdk.ServiceClient, volume *volumes.Volume) {
	t.Logf("Attempting to delete volume: %s", volume.ID)

	err := volumes.Delete(client, volumes.DeleteOpts{
		VolumeId: volume.ID,
	})
	if err != nil {
		t.Fatalf("Unable to delete volume %s: %v", volume.ID, err)
	}

	// VolumeTypes can't be deleted until their volumes have been,
	// so block until the volume is deleted.
	err = tools.WaitFor(func() (bool, error) {
		_, err := volumes.Get(client, volume.ID)
		if err != nil {
			return true, nil
		}

		return false, nil
	})
	if err != nil {
		t.Fatalf("Error waiting for volume to delete: %v", err)
	}

	t.Logf("Successfully deleted volume: %s", volume.ID)
}
