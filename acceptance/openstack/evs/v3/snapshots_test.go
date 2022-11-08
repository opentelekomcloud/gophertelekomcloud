package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/snapshots"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSnapshots(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	volume1, err := CreateVolume(t, client)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { DeleteVolume(t, client, volume1) })

	snapshot1, err := CreateSnapshot(t, client, volume1)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { DeleteSnapshot(t, client, snapshot1) })

	// Update snapshot
	updatedSnapshotName := tools.RandomString("ACPTTEST", 16)
	updatedSnapshotDescription := tools.RandomString("ACPTTEST", 16)
	updateOpts := snapshots.UpdateOpts{
		Name:        updatedSnapshotName,
		Description: updatedSnapshotDescription,
	}
	t.Logf("Attempting to update snapshot: %s", updatedSnapshotName)
	updatedSnapshot, err := snapshots.Update(client, snapshot1.ID, updateOpts)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, updatedSnapshot)
	th.AssertEquals(t, updatedSnapshot.Name, updatedSnapshotName)
	th.AssertEquals(t, updatedSnapshot.Description, updatedSnapshotDescription)

	volume2, err := CreateVolume(t, client)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { DeleteVolume(t, client, volume2) })

	snapshot2, err := CreateSnapshot(t, client, volume2)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { DeleteSnapshot(t, client, snapshot2) })

	listOpts := snapshots.ListOpts{
		Limit: 1,
	}

	err = snapshots.List(client, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		actual, err := snapshots.ExtractSnapshots(page)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, 1, len(actual))

		var found bool
		for _, v := range actual {
			if v.ID == snapshot1.ID || v.ID == snapshot2.ID {
				found = true
			}
		}

		th.AssertEquals(t, found, true)

		return true, nil
	})

	th.AssertNoErr(t, err)
}
