package testing

import (
	"net/http"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v2/snapshots"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateResponse(t)

	actual, err := snapshots.Create(client.ServiceClient(), snapshots.CreateOpts{
		VolumeID:    "1234",
		Force:       true,
		Name:        "snapshot-001",
		Description: "Daily backup",
		Metadata: map[string]string{
			"environment": "test",
		},
	})
	th.AssertNoErr(t, err)

	expected := &snapshots.Snapshot{
		ID:          "d32019d3-bc6e-4319-9c1d-6722fc136a22",
		CreatedAt:   time.Date(2017, 5, 30, 3, 35, 3, 0, time.UTC),
		UpdatedAt:   time.Time{},
		Name:        "snapshot-001",
		Description: "Daily backup",
		VolumeID:    "1234",
		Status:      "creating",
		Size:        0,
		Metadata: map[string]string{
			"environment": "test",
		},
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestCreateMissingVolumeID(t *testing.T) {
	_, err := snapshots.Create(client.ServiceClient(), snapshots.CreateOpts{})
	if err == nil {
		t.Fatal("expected an error for a missing volume_id")
	}
}

func TestCreateErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "create failed", http.StatusInternalServerError)
	})

	_, err := snapshots.Create(client.ServiceClient(), snapshots.CreateOpts{VolumeID: "1234"})
	if err == nil {
		t.Fatal("expected an error response")
	}
}
