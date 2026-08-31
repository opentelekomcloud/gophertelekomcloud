package subnets_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/subnets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListAllOptionsAndPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		expected := map[string]string{
			"limit":  "2",
			"vpc_id": "vpc-id",
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
				`{"subnets":[%s,%s]}`,
				strings.ReplaceAll(subnetJSON, `"id": "subnet-id"`, `"id": "subnet-1"`),
				strings.ReplaceAll(subnetJSON, `"id": "subnet-id"`, `"id": "subnet-2"`),
			)
		case "subnet-2":
			_, _ = fmt.Fprintf(
				w,
				`{"subnets":[%s]}`,
				strings.ReplaceAll(subnetJSON, `"id": "subnet-id"`, `"id": "subnet-3"`),
			)
		case "subnet-3":
			_, _ = w.Write([]byte(`{"subnets":[]}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	limit := 2
	actual, err := subnets.List(serviceClient(), subnets.ListOpts{
		Limit: &limit,
		VpcID: "vpc-id",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, len(actual))
	th.AssertEquals(t, "subnet-1", actual[0].ID)
	th.AssertEquals(t, "subnet-3", actual[2].ID)
	th.AssertEquals(t, "192.168.20.0/24", actual[0].CIDR)
}

func TestListZeroLimit(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets", func(w http.ResponseWriter, r *http.Request) {
		if actual := r.URL.Query().Get("limit"); actual != "0" {
			t.Fatalf("unexpected limit %q", actual)
		}
		_, _ = w.Write([]byte(`{"subnets":[]}`))
	})

	limit := 0
	actual, err := subnets.List(serviceClient(), subnets.ListOpts{Limit: &limit})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(actual))
}

func TestListNoOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Query().Get("marker") {
		case "":
			if actual := r.URL.RawQuery; actual != "" {
				t.Fatalf("unexpected query %q", actual)
			}
			_, _ = w.Write([]byte(`{"subnets":[` + subnetJSON + `]}`))
		case "subnet-id":
			_, _ = w.Write([]byte(`{"subnets":[]}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	actual, err := subnets.List(serviceClient(), subnets.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(actual))
	th.AssertEquals(t, "subnet-id", actual[0].ID)
	th.AssertEquals(t, "vpc-id", actual[0].VpcID)
}

func TestListInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"subnets":`))
	})

	actual, err := subnets.List(serviceClient(), subnets.ListOpts{})
	if err == nil {
		t.Fatal("expected response extraction error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
