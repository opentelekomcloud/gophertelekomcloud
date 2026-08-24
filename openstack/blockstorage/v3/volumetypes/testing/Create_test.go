package testing

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v3/volumetypes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateResponse(t)

	isPublic := true
	actual, err := volumetypes.Create(client.ServiceClient(), volumetypes.CreateOpts{
		Name:        "test_type",
		Description: "test_type_desc",
		IsPublic:    &isPublic,
		ExtraSpecs:  map[string]string{"capabilities": "gpu"},
	})
	th.AssertNoErr(t, err)

	expected := &volumetypes.VolumeType{
		ID:           "6d0ff92a-0007-4780-9ece-acfe5876966a",
		Name:         "test_type",
		Description:  "test_type_desc",
		ExtraSpecs:   map[string]string{"capabilities": "gpu"},
		IsPublic:     true,
		PublicAccess: true,
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestCreateOptionalZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	isPublic := false
	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{
			"volume_type": {
				"name": "private_type",
				"os-volume-type-access:is_public": false
			}
		}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"volume_type":{"name":"private_type"}}`)
	})

	actual, err := volumetypes.Create(client.ServiceClient(), volumetypes.CreateOpts{
		Name:     "private_type",
		IsPublic: &isPublic,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "private_type", actual.Name)
}

func TestCreateMissingNameDoesNotRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	requested := false
	th.Mux.HandleFunc("/types", func(http.ResponseWriter, *http.Request) {
		requested = true
	})

	actual, err := volumetypes.Create(client.ServiceClient(), volumetypes.CreateOpts{})
	if err == nil {
		t.Fatal("expected an error for a missing name")
	}
	if !strings.Contains(err.Error(), "Name") {
		t.Fatalf("expected name validation error, got %v", err)
	}
	if actual != nil {
		t.Fatalf("expected nil volume type, got %#v", actual)
	}
	if requested {
		t.Fatal("expected no request for invalid options")
	}
}

func TestCreateRejectsUnexpectedSuccessStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})

	actual, err := volumetypes.Create(client.ServiceClient(), volumetypes.CreateOpts{Name: "test_type"})
	if err == nil {
		t.Fatal("expected an error for an unexpected success status")
	}
	if actual != nil {
		t.Fatalf("expected nil volume type, got %#v", actual)
	}
}

func TestCreateErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "create failed", http.StatusInternalServerError)
	})

	actual, err := volumetypes.Create(client.ServiceClient(), volumetypes.CreateOpts{Name: "test_type"})
	if err == nil {
		t.Fatal("expected an error response")
	}
	if actual != nil {
		t.Fatalf("expected nil volume type, got %#v", actual)
	}
}

func TestCreateInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"volume_type":`)
	})

	actual, err := volumetypes.Create(client.ServiceClient(), volumetypes.CreateOpts{Name: "test_type"})
	if err == nil {
		t.Fatal("expected a response extraction error")
	}
	if actual == nil {
		t.Fatal("expected a zero volume type with the extraction error")
	}
}

func TestCreateResponseMeaningfulZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, _ *http.Request) {
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

	actual, err := volumetypes.Create(client.ServiceClient(), volumetypes.CreateOpts{Name: "test_type"})
	th.AssertNoErr(t, err)

	expected := &volumetypes.VolumeType{
		ExtraSpecs: map[string]string{},
	}
	th.AssertDeepEquals(t, expected, actual)
}
