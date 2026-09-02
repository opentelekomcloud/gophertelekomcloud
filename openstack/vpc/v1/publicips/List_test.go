package publicips_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/publicips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListAllOptionsAndPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		expected := map[string]string{
			"limit":                 "2",
			"ip_version":            "4",
			"enterprise_project_id": "0",
		}
		for key, value := range expected {
			if actual := r.URL.Query().Get(key); actual != value {
				t.Fatalf("unexpected %s: %q", key, actual)
			}
		}

		switch r.URL.Query().Get("marker") {
		case "":
			_, _ = fmt.Fprintf(
				w,
				`{"publicips":[%s,%s]}`,
				strings.ReplaceAll(publicIPJSON, `"id": "publicip-id"`, `"id": "publicip-1"`),
				strings.ReplaceAll(publicIPJSON, `"id": "publicip-id"`, `"id": "publicip-2"`),
			)
		case "publicip-2":
			_, _ = fmt.Fprintf(
				w,
				`{"publicips":[%s]}`,
				strings.ReplaceAll(publicIPJSON, `"id": "publicip-id"`, `"id": "publicip-3"`),
			)
		case "publicip-3":
			_, _ = w.Write([]byte(`{"publicips":[]}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	limit := 2
	actual, err := publicips.List(serviceClient(), publicips.ListOpts{
		Limit:               limit,
		IPVersion:           4,
		EnterpriseProjectId: "0",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, len(actual))
	th.AssertEquals(t, "publicip-1", actual[0].ID)
	th.AssertEquals(t, "publicip-3", actual[2].ID)
	th.AssertEquals(t, "161.17.17.12", actual[0].PublicIpAddress)
}

func TestListEmptyResult(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips", func(w http.ResponseWriter, r *http.Request) {
		if actual := r.URL.RawQuery; actual != "" {
			t.Fatalf("unexpected query %q", actual)
		}
		_, _ = w.Write([]byte(`{"publicips":[]}`))
	})

	actual, err := publicips.List(serviceClient(), publicips.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(actual))
}

func TestListNoOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Query().Get("marker") {
		case "":
			if actual := r.URL.RawQuery; actual != "" {
				t.Fatalf("unexpected query %q", actual)
			}
			_, _ = w.Write([]byte(`{"publicips":[` + publicIPJSON + `]}`))
		case "publicip-id":
			_, _ = w.Write([]byte(`{"publicips":[]}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	actual, err := publicips.List(serviceClient(), publicips.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(actual))
	th.AssertEquals(t, "publicip-id", actual[0].ID)
	th.AssertEquals(t, "5_bgp", actual[0].Type)
}

func TestListInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"publicips":`))
	})

	actual, err := publicips.List(serviceClient(), publicips.ListOpts{})
	if err == nil {
		t.Fatal("expected response extraction error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
