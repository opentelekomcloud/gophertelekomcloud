package testing

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v2/volumes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestGetRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/"+volumeID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		body, err := io.ReadAll(r.Body)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, "", string(body))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"volume":{"id":"`+volumeID+`"}}`)
	})

	actual, err := volumes.Get(client.ServiceClient(), volumeID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, volumeID, actual.ID)
}

func TestGetResponseFixture(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	actual, err := volumes.Get(client.ServiceClient(), volumeID)
	th.AssertNoErr(t, err)

	expected := &volumes.Volume{
		ID:               volumeID,
		Status:           "available",
		Size:             75,
		AvailabilityZone: "nova",
		CreatedAt:        time.Date(2015, 9, 17, 3, 32, 29, 0, time.UTC),
		Attachments: []volumes.Attachment{
			{
				AttachedAt:   time.Date(2016, 8, 6, 14, 48, 20, 0, time.UTC),
				AttachmentID: "05551600-a936-4d4a-ba42-79a037c1-c91a",
				Device:       "/dev/vdc",
				HostName:     "foobar",
				ID:           "d6cacb1a-8b59-4c88-ad90-d70ebb82bb75",
				ServerID:     "83ec2e3b-4321-422b-8706-a84185f52a0a",
				VolumeID:     "d6cacb1a-8b59-4c88-ad90-d70ebb82bb75",
			},
		},
		Name:               "vol-001",
		VolumeType:         "lvmdriver-1",
		Metadata:           map[string]string{},
		UserID:             "ff1ce52c03ab433aaba9108c2e3ef541",
		TenantID:           "304dc00909ac4d0da6c62d816bcb3459",
		Bootable:           "false",
		ReplicationStatus:  "disabled",
		ConsistencyGroupID: "",
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestGetRejectsUnexpectedSuccessStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/"+volumeID, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	})

	actual, err := volumes.Get(client.ServiceClient(), volumeID)
	if err == nil {
		t.Fatal("expected an error for an unexpected success status")
	}
	if actual != nil {
		t.Fatalf("expected nil volume, got %#v", actual)
	}
}

func TestGetErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/"+volumeID, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "get failed", http.StatusInternalServerError)
	})

	actual, err := volumes.Get(client.ServiceClient(), volumeID)
	if err == nil {
		t.Fatal("expected an error response")
	}
	if actual != nil {
		t.Fatalf("expected nil volume, got %#v", actual)
	}
}

func TestGetInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/"+volumeID, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"volume":`)
	})

	actual, err := volumes.Get(client.ServiceClient(), volumeID)
	if err == nil {
		t.Fatal("expected a response extraction error")
	}
	if actual == nil {
		t.Fatal("expected a zero volume with the extraction error")
	}
}

func TestGetResponseMeaningfulZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/"+volumeID, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
			"volume": {
				"attachments": [],
				"encrypted": false,
				"metadata": {},
				"multiattach": false,
				"size": 0,
				"updated_at": null
			}
		}`)
	})

	actual, err := volumes.Get(client.ServiceClient(), volumeID)
	th.AssertNoErr(t, err)

	expected := &volumes.Volume{
		Attachments: []volumes.Attachment{},
		Metadata:    map[string]string{},
	}
	th.AssertDeepEquals(t, expected, actual)
}
