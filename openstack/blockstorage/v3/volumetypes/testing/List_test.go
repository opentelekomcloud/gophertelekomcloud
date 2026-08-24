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

func TestListRequestOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	isPublic := false
	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestFormValues(t, r, map[string]string{
			"is_public": "false",
			"sort":      "name:asc,id:desc",
			"sort_key":  "created_at",
			"sort_dir":  "desc",
			"limit":     "50",
			"offset":    "2",
			"marker":    "volume-type-marker",
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"volume_types":[{"id":"volume-type-id"}]}`)
	})

	actual, err := volumetypes.List(client.ServiceClient(), volumetypes.ListOpts{
		IsPublic: &isPublic,
		Sort:     "name:asc,id:desc",
		SortKey:  "created_at",
		SortDir:  "desc",
		Limit:    50,
		Offset:   2,
		Marker:   "volume-type-marker",
	})
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []volumetypes.VolumeType{{ID: "volume-type-id"}}, actual)
}

func TestListEscapesQueryValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, r *http.Request) {
		th.TestFormValues(t, r, map[string]string{
			"sort":   "name&is_public=true",
			"marker": "id=1&limit=1000",
		})
		if strings.Contains(r.URL.RawQuery, "&is_public=") ||
			strings.Contains(r.URL.RawQuery, "&limit=") {
			t.Fatalf("query value injected an extra parameter: %s", r.URL.RawQuery)
		}
		_, _ = w.Write([]byte(`{"volume_types":[]}`))
	})

	_, err := volumetypes.List(client.ServiceClient(), volumetypes.ListOpts{
		Sort:   "name&is_public=true",
		Marker: "id=1&limit=1000",
	})
	th.AssertNoErr(t, err)
}

func TestListZeroOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "" {
			t.Fatalf("expected no query string, got %q", r.URL.RawQuery)
		}
		_, _ = w.Write([]byte(`{"volume_types":[],"volume_type_links":null}`))
	})

	actual, err := volumetypes.List(client.ServiceClient(), volumetypes.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []volumetypes.VolumeType{}, actual)
}

func TestListResponseFixtureAndAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	actual, err := volumetypes.List(client.ServiceClient(), volumetypes.ListOpts{})
	th.AssertNoErr(t, err)

	expected := []volumetypes.VolumeType{
		{
			ID:           "6685584b-1eac-4da6-b5c3-555430cf68ff",
			Name:         "SSD",
			ExtraSpecs:   map[string]string{"volume_backend_name": "lvmdriver-1"},
			IsPublic:     true,
			PublicAccess: true,
		},
		{
			ID:           "8eb69a46-df97-4e41-9586-9a40a7533803",
			Name:         "SATA",
			ExtraSpecs:   map[string]string{"volume_backend_name": "lvmdriver-1"},
			IsPublic:     true,
			PublicAccess: true,
		},
		{
			ID:         "0d05383a-6db1-4c73-9258-5f8f73f18462",
			Name:       "SAS",
			ExtraSpecs: map[string]string{},
		},
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestListRejectsUnexpectedSuccessStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	})

	actual, err := volumetypes.List(client.ServiceClient(), volumetypes.ListOpts{})
	if err == nil {
		t.Fatal("expected an error for an unexpected success status")
	}
	if actual != nil {
		t.Fatalf("expected nil volume types, got %#v", actual)
	}
}

func TestListErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "list failed", http.StatusInternalServerError)
	})

	actual, err := volumetypes.List(client.ServiceClient(), volumetypes.ListOpts{})
	if err == nil {
		t.Fatal("expected an error response")
	}
	if actual != nil {
		t.Fatalf("expected nil volume types, got %#v", actual)
	}
}

func TestListInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"volume_types":`))
	})

	actual, err := volumetypes.List(client.ServiceClient(), volumetypes.ListOpts{})
	if err == nil {
		t.Fatal("expected a response extraction error")
	}
	if actual != nil {
		t.Fatalf("expected nil volume types, got %#v", actual)
	}
}

func TestListResponseMeaningfulZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{
			"volume_types": [{
				"description": null,
				"extra_specs": {},
				"is_public": false,
				"os-volume-type-access:is_public": false,
				"qos_specs_id": null
			}],
			"volume_type_links": []
		}`))
	})

	actual, err := volumetypes.List(client.ServiceClient(), volumetypes.ListOpts{})
	th.AssertNoErr(t, err)

	expected := []volumetypes.VolumeType{{
		ExtraSpecs: map[string]string{},
	}}
	th.AssertDeepEquals(t, expected, actual)
}
