package testing

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v2/volumes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListRequestOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		expected := map[string]string{
			"all_tenants":       "true",
			"project_id":        "project-id",
			"name":              "volume-name",
			"status":            "available",
			"metadata":          "{'environment':'production'}",
			"availability_zone": "eu-de-01",
			"sort":              "name:asc,status:desc",
			"sort_key":          "created_at",
			"sort_dir":          "desc",
			"limit":             "50",
			"offset":            "2",
		}

		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Query().Get("marker") {
		case "volume-marker":
			expected["marker"] = "volume-marker"
			th.TestFormValues(t, r, expected)
			q := r.URL.Query()
			q.Set("marker", "next")
			_, _ = fmt.Fprintf(w, `{
				"volumes": [{"id":"volume-id","name":"volume-name"}],
				"volumes_links": [{"href":%q,"rel":"next"}]
			}`, th.Server.URL+"/volumes/detail?"+q.Encode())
		case "next":
			expected["marker"] = "next"
			th.TestFormValues(t, r, expected)
			_, _ = w.Write([]byte(`{
				"volumes": [{"id":"next-volume-id","name":"volume-name"}],
				"volumes_links": null
			}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	actual, err := volumes.List(client.ServiceClient(), volumes.ListOpts{
		AllTenants:       true,
		TenantID:         "project-id",
		Name:             "volume-name",
		Status:           "available",
		Metadata:         map[string]string{"environment": "production"},
		AvailabilityZone: "eu-de-01",
		Sort:             "name:asc,status:desc",
		SortKey:          "created_at",
		SortDir:          "desc",
		Limit:            50,
		Offset:           2,
		Marker:           "volume-marker",
	})
	th.AssertNoErr(t, err)

	expected := &volumes.ListResponse{
		Volumes: []volumes.Volume{
			{
				ID:   "volume-id",
				Name: "volume-name",
			},
			{
				ID:   "next-volume-id",
				Name: "volume-name",
			},
		},
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestListEscapesQueryValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestFormValues(t, r, map[string]string{
			"name":     "volume&status=error",
			"metadata": "{'key&admin':'value=1&project_id=other'}",
			"marker":   "id=1&all_tenants=true",
		})
		if strings.Contains(r.URL.RawQuery, "&all_tenants=") ||
			strings.Contains(r.URL.RawQuery, "&project_id=") ||
			strings.Contains(r.URL.RawQuery, "&status=") {
			t.Fatalf("query value injected an extra parameter: %s", r.URL.RawQuery)
		}
		_, _ = w.Write([]byte(`{"volumes":[],"volumes_links":[]}`))
	})

	_, err := volumes.List(client.ServiceClient(), volumes.ListOpts{
		Name:     "volume&status=error",
		Metadata: map[string]string{"key&admin": "value=1&project_id=other"},
		Marker:   "id=1&all_tenants=true",
	})
	th.AssertNoErr(t, err)
}

func TestListZeroOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestFormValues(t, r, map[string]string{})
		if r.URL.RawQuery != "" {
			t.Fatalf("expected no query string, got %q", r.URL.RawQuery)
		}
		if _, ok := r.URL.Query()["all_tenants"]; ok {
			t.Fatal("zero AllTenants option was sent")
		}
		if _, ok := r.URL.Query()["project_id"]; ok {
			t.Fatal("zero TenantID option was sent")
		}
		_, _ = w.Write([]byte(`{"volumes":[],"volumes_links":null}`))
	})

	actual, err := volumes.List(client.ServiceClient(), volumes.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &volumes.ListResponse{
		Volumes: []volumes.Volume{},
	}, actual)
}

func TestListTenantOptionsAreIndependent(t *testing.T) {
	tests := []struct {
		name     string
		opts     volumes.ListOpts
		expected map[string]string
	}{
		{
			name: "all tenants only",
			opts: volumes.ListOpts{AllTenants: true},
			expected: map[string]string{
				"all_tenants": "true",
			},
		},
		{
			name: "tenant only",
			opts: volumes.ListOpts{TenantID: "project-id"},
			expected: map[string]string{
				"project_id": "project-id",
			},
		},
		{
			name: "all tenants and tenant",
			opts: volumes.ListOpts{
				AllTenants: true,
				TenantID:   "project-id",
			},
			expected: map[string]string{
				"all_tenants": "true",
				"project_id":  "project-id",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()

			th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, r *http.Request) {
				th.TestFormValues(t, r, test.expected)
				_, _ = w.Write([]byte(`{"volumes":[]}`))
			})

			_, err := volumes.List(client.ServiceClient(), test.opts)
			th.AssertNoErr(t, err)
		})
	}
}

func TestListResponseFixture(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	actual, err := volumes.List(client.ServiceClient(), volumes.ListOpts{})
	th.AssertNoErr(t, err)

	expected := &volumes.ListResponse{
		Volumes: []volumes.Volume{
			{
				ID:   "289da7f8-6440-407c-9fb4-7db01ec49164",
				Name: "vol-001",
				Attachments: []volumes.Attachment{{
					ServerID:     "83ec2e3b-4321-422b-8706-a84185f52a0a",
					AttachmentID: "05551600-a936-4d4a-ba42-79a037c1-c91a",
					AttachedAt:   time.Date(2016, 8, 6, 14, 48, 20, 0, time.UTC),
					HostName:     "foobar",
					VolumeID:     "d6cacb1a-8b59-4c88-ad90-d70ebb82bb75",
					Device:       "/dev/vdc",
					ID:           "d6cacb1a-8b59-4c88-ad90-d70ebb82bb75",
				}},
				AvailabilityZone:  "nova",
				Bootable:          "false",
				CreatedAt:         time.Date(2015, 9, 17, 3, 35, 3, 0, time.UTC),
				Metadata:          map[string]string{"foo": "bar"},
				ReplicationStatus: "disabled",
				Size:              75,
				Status:            "available",
				UserID:            "ff1ce52c03ab433aaba9108c2e3ef541",
				TenantID:          "304dc00909ac4d0da6c62d816bcb3459",
				VolumeType:        "lvmdriver-1",
			},
			{
				ID:                "96c3bda7-c82a-4f50-be73-ca7621794835",
				Name:              "vol-002",
				Attachments:       []volumes.Attachment{},
				AvailabilityZone:  "nova",
				Bootable:          "false",
				CreatedAt:         time.Date(2015, 9, 17, 3, 32, 29, 0, time.UTC),
				Metadata:          map[string]string{},
				ReplicationStatus: "disabled",
				Size:              75,
				Status:            "available",
				UserID:            "ff1ce52c03ab433aaba9108c2e3ef541",
				TenantID:          "304dc00909ac4d0da6c62d816bcb3459",
				VolumeType:        "lvmdriver-1",
			},
		},
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestIDFromNameAggregatesPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)

		switch r.URL.Query().Get("marker") {
		case "":
			th.TestFormValues(t, r, map[string]string{
				"limit": "1000",
				"name":  "target",
			})
			_, _ = w.Write([]byte(`{
				"volumes": [{"id":"first","name":"other"}],
				"volumes_links": [{"href":"` + th.Server.URL + `/volumes/detail?limit=1000&marker=first&name=target","rel":"next"}]
			}`))
		case "first":
			th.TestFormValues(t, r, map[string]string{
				"limit":  "1000",
				"marker": "first",
				"name":   "target",
			})
			_, _ = w.Write([]byte(`{
				"volumes": [{"id":"second","name":"target"}],
				"volumes_links": null
			}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	actual, err := volumes.IDFromName(client.ServiceClient(), "target")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "second", actual)
}

func TestIDFromNameDetectsDuplicatesAcrossPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Query().Get("marker") {
		case "":
			_, _ = w.Write([]byte(`{
				"volumes": [{"id":"first","name":"duplicate"}],
				"volumes_links": [{"href":"` + th.Server.URL + `/volumes/detail?limit=1000&marker=first&name=duplicate","rel":"next"}]
			}`))
		case "first":
			_, _ = w.Write([]byte(`{
				"volumes": [{"id":"second","name":"duplicate"}],
				"volumes_links": null
			}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	_, err := volumes.IDFromName(client.ServiceClient(), "duplicate")
	actual, ok := err.(golangsdk.ErrMultipleResourcesFound)
	if !ok {
		t.Fatalf("expected ErrMultipleResourcesFound, got %T: %v", err, err)
	}
	th.AssertEquals(t, 2, actual.Count)
}

func TestListErrorResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "list failed", http.StatusInternalServerError)
	})

	actual, err := volumes.List(client.ServiceClient(), volumes.ListOpts{})
	if err == nil {
		t.Fatal("expected an error response")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}

func TestListMalformedResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"volumes":`))
	})

	actual, err := volumes.List(client.ServiceClient(), volumes.ListOpts{})
	if err == nil {
		t.Fatal("expected a response extraction error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}

func TestListResponseMeaningfulZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{
			"volumes": [{
				"attachments": [],
				"encrypted": false,
				"metadata": {},
				"multiattach": false,
				"size": 0,
				"updated_at": null
			}],
			"volumes_links": []
		}`))
	})

	actual, err := volumes.List(client.ServiceClient(), volumes.ListOpts{})
	th.AssertNoErr(t, err)

	expected := &volumes.ListResponse{
		Volumes: []volumes.Volume{{
			Attachments: []volumes.Attachment{},
			Metadata:    map[string]string{},
		}},
	}
	th.AssertDeepEquals(t, expected, actual)
}
