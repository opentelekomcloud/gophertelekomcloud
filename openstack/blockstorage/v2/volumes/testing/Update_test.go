package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v2/volumes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/"+volumeID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `{
			"volume": {
				"name": "vol-002",
				"description": "updated volume",
				"metadata": {
					"key": "value"
				}
			}
		}`)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
			"volume": {
				"id": "`+volumeID+`",
				"name": "vol-002",
				"description": "updated volume",
				"metadata": {
					"key": "value"
				},
				"status": "available"
			}
		}`)
	})

	actual, err := volumes.Update(client.ServiceClient(), volumeID, volumes.UpdateOpts{
		Name:        "vol-002",
		Description: "updated volume",
		Metadata: map[string]string{
			"key": "value",
		},
	})
	th.AssertNoErr(t, err)

	expected := &volumes.Volume{
		ID:          volumeID,
		Name:        "vol-002",
		Description: "updated volume",
		Metadata:    map[string]string{"key": "value"},
		Status:      "available",
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestUpdateRejectsUnexpectedSuccessStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/"+volumeID, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	})

	actual, err := volumes.Update(client.ServiceClient(), volumeID, volumes.UpdateOpts{Name: "vol-002"})
	if err == nil {
		t.Fatal("expected an error for an unexpected success status")
	}
	if actual != nil {
		t.Fatalf("expected nil volume, got %#v", actual)
	}
}

func TestUpdateErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/"+volumeID, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "update failed", http.StatusInternalServerError)
	})

	actual, err := volumes.Update(client.ServiceClient(), volumeID, volumes.UpdateOpts{Name: "vol-002"})
	if err == nil {
		t.Fatal("expected an error response")
	}
	if actual != nil {
		t.Fatalf("expected nil volume, got %#v", actual)
	}
}

func TestUpdateInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/"+volumeID, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"volume":`)
	})

	actual, err := volumes.Update(client.ServiceClient(), volumeID, volumes.UpdateOpts{Name: "vol-002"})
	if err == nil {
		t.Fatal("expected a response extraction error")
	}
	if actual == nil {
		t.Fatal("expected a zero volume with the extraction error")
	}
}

func TestUpdateResponseMeaningfulZeroValues(t *testing.T) {
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

	actual, err := volumes.Update(client.ServiceClient(), volumeID, volumes.UpdateOpts{Name: "vol-002"})
	th.AssertNoErr(t, err)

	expected := &volumes.Volume{
		Attachments: []volumes.Attachment{},
		Metadata:    map[string]string{},
	}
	th.AssertDeepEquals(t, expected, actual)
}
