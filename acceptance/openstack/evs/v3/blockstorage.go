// Package v3 contains common functions for creating block storage based
// resources for use in acceptance tests. See the `*_test.go` files for
// example usages.
package v3

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/qos"
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
		Size:        1,
		Name:        volumeName,
		Description: volumeDescription,
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
	// so block until the snapshoth as been deleted.
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

	err := volumes.Delete(client, volume.ID, volumes.DeleteOpts{})
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

// DeleteVolumeType will delete a volume type. A fatal error will occur if the
// volume type failed to be deleted. This works best when used as a deferred
// function.
func DeleteVolumeType(t *testing.T, client *golangsdk.ServiceClient, vt *volumetypes.VolumeType) {
	t.Logf("Attempting to delete volume type: %s", vt.ID)

	err := volumetypes.Delete(client, vt.ID)
	if err != nil {
		t.Fatalf("Unable to delete volume %s: %v", vt.ID, err)
	}

	t.Logf("Successfully deleted volume type: %s", vt.ID)
}

// CreateQoS will create a QoS with one spec and a random name. An
// error will be returned if the volume was unable to be created.
func CreateQoS(t *testing.T, client *golangsdk.ServiceClient) (*qos.QoS, error) {
	name := tools.RandomString("ACPTTEST", 16)
	t.Logf("Attempting to create QoS: %s", name)

	createOpts := qos.CreateOpts{
		Name:     name,
		Consumer: qos.ConsumerFront,
		Specs: map[string]string{
			"read_iops_sec": "20000",
		},
	}

	qs, err := qos.Create(client, createOpts)
	if err != nil {
		return nil, err
	}

	tools.PrintResource(t, qs)
	th.AssertEquals(t, qs.Consumer, "front-end")
	th.AssertEquals(t, qs.Name, name)
	th.AssertDeepEquals(t, qs.Specs, createOpts.Specs)

	t.Logf("Successfully created QoS: %s", qs.ID)

	return qs, nil
}

// DeleteQoS will delete a QoS. A fatal error will occur if the QoS
// failed to be deleted. This works best when used as a deferred function.
func DeleteQoS(t *testing.T, client *golangsdk.ServiceClient, qs *qos.QoS) {
	t.Logf("Attempting to delete QoS: %s", qs.ID)

	deleteOpts := qos.DeleteOpts{
		Force: true,
	}

	err := qos.Delete(client, qs.ID, deleteOpts)
	if err != nil {
		t.Fatalf("Unable to delete QoS %s: %v", qs.ID, err)
	}

	t.Logf("Successfully deleted QoS: %s", qs.ID)
}
