package testing

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v3/volumetypes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestGetRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/"+volumeTypeID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		body, err := io.ReadAll(r.Body)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, "", string(body))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"volume_type":{"id":"`+volumeTypeID+`"}}`)
	})

	actual, err := volumetypes.Get(client.ServiceClient(), volumeTypeID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, volumeTypeID, actual.ID)
}

func TestGetResponseFixture(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	actual, err := volumetypes.Get(client.ServiceClient(), volumeTypeID)
	th.AssertNoErr(t, err)

	expected := &volumetypes.VolumeType{
		ID:           volumeTypeID,
		Name:         "vol-type-001",
		Description:  "volume type 001",
		ExtraSpecs:   map[string]string{"capabilities": "gpu"},
		IsPublic:     true,
		QosSpecID:    volumeTypeID,
		PublicAccess: true,
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestGetRejectsUnexpectedSuccessStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/"+volumeTypeID, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	})

	actual, err := volumetypes.Get(client.ServiceClient(), volumeTypeID)
	if err == nil {
		t.Fatal("expected an error for an unexpected success status")
	}
	if actual != nil {
		t.Fatalf("expected nil volume type, got %#v", actual)
	}
}

func TestGetErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/"+volumeTypeID, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "get failed", http.StatusInternalServerError)
	})

	actual, err := volumetypes.Get(client.ServiceClient(), volumeTypeID)
	if err == nil {
		t.Fatal("expected an error response")
	}
	if actual != nil {
		t.Fatalf("expected nil volume type, got %#v", actual)
	}
}

func TestGetInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types/"+volumeTypeID, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"volume_type":`)
	})

	actual, err := volumetypes.Get(client.ServiceClient(), volumeTypeID)
	if err == nil {
		t.Fatal("expected a response extraction error")
	}
	if actual == nil {
		t.Fatal("expected a zero volume type with the extraction error")
	}
}

func TestGetResponseMeaningfulZeroValues(t *testing.T) {
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

	actual, err := volumetypes.Get(client.ServiceClient(), volumeTypeID)
	th.AssertNoErr(t, err)

	expected := &volumetypes.VolumeType{
		ExtraSpecs: map[string]string{},
	}
	th.AssertDeepEquals(t, expected, actual)
}
