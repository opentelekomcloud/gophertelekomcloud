package testing

import (
	"net/http"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v2/snapshots"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	actual, err := snapshots.Get(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertNoErr(t, err)

	expected := &snapshots.Snapshot{
		ID:          "d32019d3-bc6e-4319-9c1d-6722fc136a22",
		CreatedAt:   time.Date(2017, 5, 30, 3, 35, 3, 0, time.UTC),
		UpdatedAt:   time.Time{},
		Name:        "snapshot-001",
		Description: "Daily backup",
		VolumeID:    "521752a6-acf6-4b2d-bc7a-119f9148cd8c",
		Status:      "available",
		Size:        0,
		Metadata:    map[string]string{},
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestGetErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "get failed", http.StatusInternalServerError)
	})

	actual, err := snapshots.Get(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	if err == nil {
		t.Fatal("expected an error response")
	}
	if actual != nil {
		t.Fatalf("expected nil snapshot, got %#v", actual)
	}
}
