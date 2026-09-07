package privateips_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/privateips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestListAllOptionsAndPagination(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	requestCount := 0
	th.Mux.HandleFunc("/project-id/subnets/network-id/privateips", func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		th.TestMethod(t, r, http.MethodGet)
		if actual := r.URL.Query().Get("limit"); actual != "2" {
			t.Fatalf("unexpected limit %q", actual)
		}

		switch r.URL.Query().Get("marker") {
		case "start-id":
			_, _ = fmt.Fprintf(
				w,
				`{"privateips":[%s,%s]}`,
				strings.ReplaceAll(privateIPJSON, `"private-ip-id"`, `"private-ip-1"`),
				strings.ReplaceAll(privateIPJSON, `"private-ip-id"`, `"private-ip-2"`),
			)
		case "private-ip-2":
			_, _ = fmt.Fprintf(
				w,
				`{"privateips":[%s]}`,
				strings.ReplaceAll(privateIPJSON, `"private-ip-id"`, `"private-ip-3"`),
			)
		case "private-ip-3":
			_, _ = w.Write([]byte(`{"privateips":[]}`))
		default:
			t.Fatalf("unexpected marker %q", r.URL.Query().Get("marker"))
		}
	})

	limit := 2
	actual, err := privateips.List(serviceClient(), "network-id", privateips.ListOpts{
		Marker: "start-id",
		Limit:  &limit,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, len(actual))
	th.AssertEquals(t, "private-ip-1", actual[0].ID)
	th.AssertEquals(t, "private-ip-3", actual[2].ID)
	th.AssertEquals(t, "192.168.20.10", actual[0].IPAddress)
	th.AssertEquals(t, 3, requestCount)
}

func TestListZeroLimit(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets/network-id/privateips", func(w http.ResponseWriter, r *http.Request) {
		if actual := r.URL.Query().Get("limit"); actual != "0" {
			t.Fatalf("unexpected limit %q", actual)
		}
		_, _ = w.Write([]byte(`{"privateips":[]}`))
	})

	limit := 0
	actual, err := privateips.List(serviceClient(), "network-id", privateips.ListOpts{Limit: &limit})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(actual))
}

func TestListNoOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets/network-id/privateips", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "" {
			t.Fatalf("unexpected query %q", r.URL.RawQuery)
		}
		_, _ = w.Write([]byte(`{"privateips":[]}`))
	})

	actual, err := privateips.List(serviceClient(), "network-id", privateips.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(actual))
}

func TestListInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets/network-id/privateips", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"privateips":`))
	})

	actual, err := privateips.List(serviceClient(), "network-id", privateips.ListOpts{})
	if err == nil {
		t.Fatal("expected response extraction error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}

func TestListError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/subnets/network-id/privateips", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0207","error_msg":"Invalid private IP parameter."}`))
	})

	actual, err := privateips.List(serviceClient(), "network-id", privateips.ListOpts{})
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
