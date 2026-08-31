package vpcs_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListAllOptionsAndPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		expected := map[string]string{
			"id":                    "vpc-id",
			"limit":                 "2",
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
				`{"vpcs":[%s,%s]}`,
				strings.ReplaceAll(vpcJSON, "vpc-id", "vpc-1"),
				strings.ReplaceAll(vpcJSON, "vpc-id", "vpc-2"),
			)
		case "vpc-2":
			_, _ = fmt.Fprintf(
				w,
				`{"vpcs":[%s]}`,
				strings.ReplaceAll(vpcJSON, "vpc-id", "vpc-3"),
			)
		case "vpc-3":
			_, _ = w.Write([]byte(`{"vpcs":[]}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	limit := 2
	actual, err := vpcs.List(serviceClient(), vpcs.ListOpts{
		ID:                  "vpc-id",
		Limit:               &limit,
		EnterpriseProjectID: "0",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, len(actual))
	th.AssertEquals(t, "vpc-1", actual[0].ID)
	th.AssertEquals(t, "vpc-3", actual[2].ID)
	th.AssertEquals(t, "192.168.0.0/16", actual[0].CIDR)
}

func TestListZeroLimit(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs", func(w http.ResponseWriter, r *http.Request) {
		if actual := r.URL.Query().Get("limit"); actual != "0" {
			t.Fatalf("unexpected limit %q", actual)
		}
		_, _ = w.Write([]byte(`{"vpcs":[]}`))
	})

	limit := 0
	actual, err := vpcs.List(serviceClient(), vpcs.ListOpts{Limit: &limit})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(actual))
}

func TestListNoOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Query().Get("marker") {
		case "":
			if actual := r.URL.RawQuery; actual != "" {
				t.Fatalf("unexpected query %q", actual)
			}
			_, _ = w.Write([]byte(`{"vpcs":[` + vpcJSON + `]}`))
		case "vpc-id":
			_, _ = w.Write([]byte(`{"vpcs":[]}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	actual, err := vpcs.List(serviceClient(), vpcs.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(actual))
	th.AssertEquals(t, "vpc-id", actual[0].ID)
}

func TestListInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"vpcs":`))
	})

	actual, err := vpcs.List(serviceClient(), vpcs.ListOpts{})
	if err == nil {
		t.Fatal("expected response extraction error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
