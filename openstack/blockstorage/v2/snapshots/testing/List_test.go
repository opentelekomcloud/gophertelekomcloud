package testing

import (
	"net/http"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v2/snapshots"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	actual, err := snapshots.List(client.ServiceClient(), snapshots.ListOpts{
		AllTenants: true,
		Name:       "snapshot-001",
		Status:     "available",
		TenantID:   "0c2eba2c5af04d3f9e9d0d410b371fde",
		VolumeID:   "521752a6-acf6-4b2d-bc7a-119f9148cd8c",
		Marker:     "289da7f8-6440-407c-9fb4-7db01ec49164",
		Offset:     1,
		Limit:      2,
	})
	th.AssertNoErr(t, err)

	expected := &snapshots.ListResponse{
		Snapshots: []snapshots.Snapshot{
			{
				ID:          "289da7f8-6440-407c-9fb4-7db01ec49164",
				CreatedAt:   time.Date(2017, 5, 30, 3, 35, 3, 0, time.UTC),
				UpdatedAt:   time.Time{},
				Name:        "snapshot-001",
				Description: "Daily Backup",
				VolumeID:    "521752a6-acf6-4b2d-bc7a-119f9148cd8c",
				Status:      "available",
				Size:        0,
				Metadata:    map[string]string{},
			},
			{
				ID:          "96c3bda7-c82a-4f50-be73-ca7621794835",
				CreatedAt:   time.Date(2017, 5, 30, 3, 35, 3, 0, time.UTC),
				UpdatedAt:   time.Date(2017, 5, 31, 3, 35, 3, 0, time.UTC),
				Name:        "snapshot-002",
				Description: "Weekly Backup",
				VolumeID:    "76b8950a-8594-4e5b-8dce-0dfa9c696358",
				Status:      "available",
				Size:        25,
				Metadata: map[string]string{
					"environment": "test",
				},
			},
			{
				ID:          "d32019d3-bc6e-4319-9c1d-6722fc136a22",
				CreatedAt:   time.Date(2017, 6, 1, 3, 35, 3, 0, time.UTC),
				UpdatedAt:   time.Time{},
				Name:        "snapshot-003",
				Description: "Monthly Backup",
				VolumeID:    "521752a6-acf6-4b2d-bc7a-119f9148cd8c",
				Status:      "available",
				Size:        50,
				Metadata:    map[string]string{},
			},
		},
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestListZeroOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestFormValues(t, r, map[string]string{})
		w.Header().Add("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"snapshots":[],"snapshots_links":null}`))
	})

	actual, err := snapshots.List(client.ServiceClient(), snapshots.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &snapshots.ListResponse{
		Snapshots: []snapshots.Snapshot{},
	}, actual)
}

func TestListErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "list failed", http.StatusInternalServerError)
	})

	actual, err := snapshots.List(client.ServiceClient(), snapshots.ListOpts{})
	if err == nil {
		t.Fatal("expected an error response")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}

func TestListExtractionError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"snapshots":`))
	})

	actual, err := snapshots.List(client.ServiceClient(), snapshots.ListOpts{})
	if err == nil {
		t.Fatal("expected an extraction error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}

func TestIDFromNamePagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")

		switch r.URL.Query().Get("marker") {
		case "":
			th.TestFormValues(t, r, map[string]string{
				"limit": "1000",
				"name":  "duplicate",
			})
			_, _ = w.Write([]byte(`{
				"snapshots": [{"id":"first","name":"duplicate"}],
				"snapshots_links": [{"href":"` + th.Server.URL + `/snapshots?limit=1000&marker=first&name=duplicate","rel":"next"}]
			}`))
		case "first":
			th.TestFormValues(t, r, map[string]string{
				"limit":  "1000",
				"marker": "first",
				"name":   "duplicate",
			})
			_, _ = w.Write([]byte(`{
				"snapshots": [{"id":"second","name":"duplicate"}],
				"snapshots_links": null
			}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	_, err := snapshots.IDFromName(client.ServiceClient(), "duplicate")
	actual, ok := err.(golangsdk.ErrMultipleResourcesFound)
	if !ok {
		t.Fatalf("expected ErrMultipleResourcesFound, got %T: %v", err, err)
	}
	th.AssertEquals(t, actual.Count, 2)
}
