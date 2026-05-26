package testing

import (
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v2/snapshots"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateResponse(t)

	options := snapshots.CreateOpts{VolumeID: "1234", Name: "snapshot-001"}
	n, err := snapshots.Create(client.ServiceClient(), options)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, n.VolumeID, "1234")
	th.AssertEquals(t, n.Name, "snapshot-001")
	th.AssertEquals(t, n.ID, "d32019d3-bc6e-4319-9c1d-6722fc136a22")
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	v, err := snapshots.Get(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v.Name, "snapshot-001")
	th.AssertEquals(t, v.ID, "d32019d3-bc6e-4319-9c1d-6722fc136a22")
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateResponse(t)

	options := snapshots.UpdateOpts{
		Name:        "snapshot-001-update",
		Description: "Weekly backup",
	}

	v, err := snapshots.Update(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22", options)

	th.AssertNoErr(t, err)
	th.AssertEquals(t, v.Name, "snapshot-001-update")
	th.AssertEquals(t, v.Description, "Weekly backup")
	th.AssertEquals(t, v.UpdatedAt, time.Date(2020, 3, 27, 15, 55, 3, 0, time.UTC))
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	actual, err := snapshots.List(client.ServiceClient(), snapshots.ListOpts{})
	th.AssertNoErr(t, err)

	expected := []snapshots.Snapshot{
		{
			ID:          "289da7f8-6440-407c-9fb4-7db01ec49164",
			Name:        "snapshot-001",
			VolumeID:    "521752a6-acf6-4b2d-bc7a-119f9148cd8c",
			Status:      "available",
			Size:        30,
			CreatedAt:   time.Date(2020, 3, 27, 15, 35, 3, 0, time.UTC),
			Description: "Daily Backup",
		},
		{
			ID:          "96c3bda7-c82a-4f50-be73-ca7621794835",
			Name:        "snapshot-002",
			VolumeID:    "76b8950a-8594-4e5b-8dce-0dfa9c696358",
			Status:      "available",
			Size:        25,
			CreatedAt:   time.Date(2020, 3, 27, 15, 35, 3, 0, time.UTC),
			Description: "Weekly Backup",
		},
	}

	th.CheckDeepEquals(t, expected, actual.Snapshots)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteResponse(t)

	err := snapshots.Delete(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertNoErr(t, err)
}

func TestRollback(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockRollbackResponse(t)

	options := snapshots.RollbackOpts{
		VolumeID: "5aa119a8-d25b-45a7-8d1b-88e127885635",
		Name:     "volume-001",
	}
	actual, err := snapshots.Rollback(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22", options)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, actual.VolumeID, "5aa119a8-d25b-45a7-8d1b-88e127885635")
}
