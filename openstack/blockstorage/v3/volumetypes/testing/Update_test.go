package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v3/volumetypes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateResponse(t)

	isPublic := false
	actual, err := volumetypes.Update(client.ServiceClient(), volumeTypeID, volumetypes.UpdateOpts{
		Name:        "vol-type-002",
		Description: "volume type 0002",
		IsPublic:    &isPublic,
	})
	th.AssertNoErr(t, err)

	expected := &volumetypes.VolumeType{
		ID:          volumeTypeID,
		Name:        "vol-type-002",
		Description: "volume type 0002",
		ExtraSpecs:  map[string]string{"capabilities": "gpu"},
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestUpdateZeroOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/"+volumeTypeID, func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{"volume_type":{}}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"volume_type":{"id":"`+volumeTypeID+`"}}`)
	})

	actual, err := volumetypes.Update(client.ServiceClient(), volumeTypeID, volumetypes.UpdateOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, volumeTypeID, actual.ID)
}

func TestUpdateRejectsUnexpectedSuccessStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/"+volumeTypeID, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	})

	actual, err := volumetypes.Update(client.ServiceClient(), volumeTypeID, volumetypes.UpdateOpts{Name: "vol-type-002"})
	if err == nil {
		t.Fatal("expected an error for an unexpected success status")
	}
	if actual != nil {
		t.Fatalf("expected nil volume type, got %#v", actual)
	}
}

func TestUpdateErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/"+volumeTypeID, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "update failed", http.StatusInternalServerError)
	})

	actual, err := volumetypes.Update(client.ServiceClient(), volumeTypeID, volumetypes.UpdateOpts{Name: "vol-type-002"})
	if err == nil {
		t.Fatal("expected an error response")
	}
	if actual != nil {
		t.Fatalf("expected nil volume type, got %#v", actual)
	}
}

func TestUpdateInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/"+volumeTypeID, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"volume_type":`)
	})

	actual, err := volumetypes.Update(client.ServiceClient(), volumeTypeID, volumetypes.UpdateOpts{Name: "vol-type-002"})
	if err == nil {
		t.Fatal("expected a response extraction error")
	}
	if actual == nil {
		t.Fatal("expected a zero volume type with the extraction error")
	}
}

func TestUpdateResponseMeaningfulZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/"+volumeTypeID, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
			"volume_type": {
				"description": null,
				"extra_specs": {},
				"is_public": false,
				"os-volume-type-access:is_public": false,
				"qos_specs_id": null
			}
		}`)
	})

	actual, err := volumetypes.Update(client.ServiceClient(), volumeTypeID, volumetypes.UpdateOpts{Name: "vol-type-002"})
	th.AssertNoErr(t, err)

	expected := &volumetypes.VolumeType{
		ExtraSpecs: map[string]string{},
	}
	th.AssertDeepEquals(t, expected, actual)
}
