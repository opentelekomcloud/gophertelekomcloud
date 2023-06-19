package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/snapshots"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/volumes"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVolumes(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	volume1, err := CreateVolume(t, client)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { DeleteVolume(t, client, volume1) })

	volume2, err := CreateVolume(t, client)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { DeleteVolume(t, client, volume2) })

	// Update volume
	updatedVolumeName := ""
	updatedVolumeDescription := ""
	updatedVolume, err := volumes.Update(client, volumes.UpdateOpts{
		VolumeId:    volume1.ID,
		Name:        updatedVolumeName,
		Description: updatedVolumeDescription,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, updatedVolume)
	th.AssertEquals(t, updatedVolume.Name, updatedVolumeName)
	th.AssertEquals(t, updatedVolume.Description, updatedVolumeDescription)

	listOpts := volumes.ListOpts{
		Limit: 1,
	}

	err = volumes.List(client, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		actual, err := volumes.ExtractVolumes(page)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, 1, len(actual))

		var found bool
		for _, v := range actual {
			if v.ID == volume1.ID || v.ID == volume2.ID {
				found = true
			}
		}

		th.AssertEquals(t, found, true)

		return true, nil
	})

	th.AssertNoErr(t, err)
}

func TestVolumesMultiAttach(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	volumeName := tools.RandomString("ACPTTEST", 16)

	volOpts := volumes.CreateOpts{
		Size:        1,
		Name:        volumeName,
		Description: "Testing creation of multiattach enabled volume",
		Multiattach: true,
	}

	vol, err := volumes.Create(client, volOpts)
	th.AssertNoErr(t, err)

	err = volumes.WaitForStatus(client, vol.ID, "available", 60)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, vol.Multiattach, true)

	err = volumes.Delete(client, volumes.DeleteOpts{VolumeId: vol.ID})
	th.AssertNoErr(t, err)
}

func TestVolumesCascadeDelete(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	vol, err := CreateVolume(t, client)
	th.AssertNoErr(t, err)

	err = volumes.WaitForStatus(client, vol.ID, "available", 60)
	th.AssertNoErr(t, err)

	snapshot1, err := CreateSnapshot(t, client, vol)
	th.AssertNoErr(t, err)

	snapshot2, err := CreateSnapshot(t, client, vol)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to delete volume: %s", vol.ID)

	err = volumes.Delete(client, volumes.DeleteOpts{
		VolumeId: vol.ID,
		Cascade:  true,
	})
	if err != nil {
		t.Fatalf("Unable to delete volume %s: %v", vol.ID, err)
	}

	for _, sid := range []string{snapshot1.ID, snapshot2.ID} {
		err := tools.WaitFor(func() (bool, error) {
			_, err := snapshots.Get(client, sid)
			if err != nil {
				return true, nil
			}
			return false, nil
		})
		th.AssertNoErr(t, err)
		t.Logf("Successfully deleted snapshot: %s", sid)
	}

	err = tools.WaitFor(func() (bool, error) {
		_, err := volumes.Get(client, vol.ID)
		if err != nil {
			return true, nil
		}
		return false, nil
	})
	th.AssertNoErr(t, err)

	t.Logf("Successfully deleted volume: %s", vol.ID)
}
