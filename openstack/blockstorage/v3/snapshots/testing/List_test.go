package testing

import (
	"net/http"
	"strings"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v3/snapshots"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	actual, err := snapshots.List(client.ServiceClient(), snapshots.ListOpts{
		Marker:       "289da7f8-6440-407c-9fb4-7db01ec49164",
		Offset:       1,
		Limit:        2,
		Name:         "snapshot-001",
		SortDir:      "asc",
		Status:       "available",
		VolumeID:     "521752a6-acf6-4b2d-bc7a-119f9148cd8c",
		NameLike:     "snapshot",
		StatusLike:   "avail",
		VolumeIDLike: "521752",
		SortKey:      "name",
		WithCount:    true,
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
		Count: 3,
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestListZeroOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		if r.URL.RawQuery != "" {
			t.Fatalf("expected no query string, got %q", r.URL.RawQuery)
		}
		w.Header().Add("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"snapshots":[],"snapshots_links":null,"count":0}`))
	})

	actual, err := snapshots.List(client.ServiceClient(), snapshots.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &snapshots.ListResponse{
		Snapshots: []snapshots.Snapshot{},
	}, actual)
}

func TestListEscapesQueryValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("name"); got != "snapshot&status=error" {
			t.Fatalf("unexpected name query value %q", got)
		}
		if got := r.URL.Query()["status"]; len(got) != 1 || got[0] != "available" {
			t.Fatalf("query injection changed status values: %#v", got)
		}
		if strings.Contains(r.URL.RawQuery, "name=snapshot&status=error") {
			t.Fatalf("name query value was not escaped: %q", r.URL.RawQuery)
		}
		_, _ = w.Write([]byte(`{"snapshots":[]}`))
	})

	_, err := snapshots.List(client.ServiceClient(), snapshots.ListOpts{
		Name:   "snapshot&status=error",
		Status: "available",
	})
	th.AssertNoErr(t, err)
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

func TestListInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"snapshots":`))
	})

	actual, err := snapshots.List(client.ServiceClient(), snapshots.ListOpts{})
	if err == nil {
		t.Fatal("expected a response extraction error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}

func TestIDFromNameUsesAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Query().Get("marker") {
		case "":
			th.TestFormValues(t, r, map[string]string{
				"limit": "1000",
				"name":  "snapshot-001",
			})
			_, _ = w.Write([]byte(`{
				"snapshots": [{"id":"other","name":"other"}],
				"snapshots_links": [{"href":"` + th.Server.URL + `/snapshots?limit=1000&marker=first&name=snapshot-001","rel":"next"}]
			}`))
		case "first":
			_, _ = w.Write([]byte(`{
				"snapshots": [{"id":"target","name":"snapshot-001"}],
				"snapshots_links": null
			}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	actual, err := snapshots.IDFromName(client.ServiceClient(), "snapshot-001")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, actual, "target")
}

func TestIDFromNameDuplicate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Query().Get("marker") {
		case "":
			_, _ = w.Write([]byte(`{
				"snapshots": [{"id":"first","name":"duplicate"}],
				"snapshots_links": [{"href":"` + th.Server.URL + `/snapshots?limit=1000&marker=first&name=duplicate","rel":"next"}]
			}`))
		case "first":
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

func TestIDFromNameNotFound(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"snapshots":[]}`))
	})

	_, err := snapshots.IDFromName(client.ServiceClient(), "missing")
	if _, ok := err.(golangsdk.ErrResourceNotFound); !ok {
		t.Fatalf("expected ErrResourceNotFound, got %T: %v", err, err)
	}
}
