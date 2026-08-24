package testing

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v3/volumes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateResponse(t)

	actual, err := volumes.Create(client.ServiceClient(), volumes.CreateOpts{
		Size:               40,
		AvailabilityZone:   "az-dc-1",
		ConsistencyGroupID: "consistency-group-id",
		Description:        "create for api test",
		Metadata: map[string]string{
			"volume_owner": "openapi",
		},
		Name:        "openapi_vol01",
		ImageID:     "027cf713-45a6-45f0-ac1b-0ccc57ac12e2",
		VolumeType:  "SSD",
		Multiattach: true,
	})
	th.AssertNoErr(t, err)

	expected := &volumes.Volume{
		ID:                 "8dd7c486-8e9f-49fe-bceb-26aa7e312b66",
		Status:             "creating",
		Size:               40,
		AvailabilityZone:   "az-dc-1",
		CreatedAt:          time.Date(2016, 5, 25, 2, 38, 40, 392463000, time.UTC),
		UpdatedAt:          time.Time{},
		Attachments:        []volumes.Attachment{},
		Name:               "openapi_vol01",
		Description:        "create for api test",
		VolumeType:         "SSD",
		SnapshotID:         "",
		SourceVolID:        "",
		Metadata:           map[string]string{"volume_owner": "openapi"},
		UserID:             "39f6696ae23740708d0f358a253c2637",
		Bootable:           "false",
		Encrypted:          false,
		ReplicationStatus:  "disabled",
		ConsistencyGroupID: "consistency-group-id",
		Multiattach:        true,
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestCreateSourceFields(t *testing.T) {
	tests := []struct {
		name string
		opts volumes.CreateOpts
		body string
	}{
		{
			name: "snapshot",
			opts: volumes.CreateOpts{
				Size:             40,
				AvailabilityZone: "az-dc-1",
				SnapshotID:       "snapshot-id",
			},
			body: `{"volume":{"availability_zone":"az-dc-1","size":40,"snapshot_id":"snapshot-id"}}`,
		},
		{
			name: "source replica",
			opts: volumes.CreateOpts{
				Size:             40,
				AvailabilityZone: "az-dc-1",
				SourceReplica:    "replica-id",
			},
			body: `{"volume":{"availability_zone":"az-dc-1","size":40,"source_replica":"replica-id"}}`,
		},
		{
			name: "source volume",
			opts: volumes.CreateOpts{
				Size:             40,
				AvailabilityZone: "az-dc-1",
				SourceVolID:      "source-volume-id",
			},
			body: `{"volume":{"availability_zone":"az-dc-1","size":40,"source_volid":"source-volume-id"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()

			th.Mux.HandleFunc("/volumes", func(w http.ResponseWriter, r *http.Request) {
				th.TestMethod(t, r, "POST")
				th.TestJSONRequest(t, r, tt.body)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusAccepted)
				_, _ = fmt.Fprint(w, `{"volume":{"id":"volume-id"}}`)
			})

			actual, err := volumes.Create(client.ServiceClient(), tt.opts)
			th.AssertNoErr(t, err)
			th.AssertEquals(t, "volume-id", actual.ID)
		})
	}
}

func TestCreateMissingSize(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	_, err := volumes.Create(client.ServiceClient(), volumes.CreateOpts{})
	if err == nil {
		t.Fatal("expected an error for a missing size")
	}
	if !strings.Contains(err.Error(), "Size") {
		t.Fatalf("expected size validation error, got %v", err)
	}
}

func TestCreateErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "create failed", http.StatusInternalServerError)
	})

	_, err := volumes.Create(client.ServiceClient(), volumes.CreateOpts{Size: 40})
	if err == nil {
		t.Fatal("expected an error response")
	}
}

func TestCreateInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		_, _ = fmt.Fprint(w, `{"volume":`)
	})

	_, err := volumes.Create(client.ServiceClient(), volumes.CreateOpts{Size: 40})
	if err == nil {
		t.Fatal("expected a response extraction error")
	}
}

func TestCreateResponseMeaningfulZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, _ = fmt.Fprint(w, `{
			"volume": {
				"attachments": [],
				"encrypted": false,
				"multiattach": false,
				"size": 0,
				"updated_at": null
			}
		}`)
	})

	actual, err := volumes.Create(client.ServiceClient(), volumes.CreateOpts{Size: 40})
	th.AssertNoErr(t, err)

	expected := &volumes.Volume{
		Attachments: []volumes.Attachment{},
	}
	th.AssertDeepEquals(t, expected, actual)
}
