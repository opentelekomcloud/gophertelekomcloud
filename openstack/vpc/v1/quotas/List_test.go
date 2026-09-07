package quotas_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/quotas"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/quotas", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.AssertEquals(t, "publicIp", r.URL.Query().Get("type"))
		_, _ = w.Write([]byte(quotasJSON))
	})

	actual, err := quotas.List(serviceClient(), quotas.ListOpts{Type: "publicIp"})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(actual))
	th.AssertEquals(t, "publicIp", actual[0].Type)
	th.AssertEquals(t, 2, actual[0].Used)
	th.AssertEquals(t, 10, actual[0].Quota)
	th.AssertEquals(t, 0, actual[0].Min)
	th.AssertEquals(t, -1, actual[1].Quota)
}

func TestListNoOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/quotas", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "" {
			t.Fatalf("unexpected query %q", r.URL.RawQuery)
		}
		_, _ = w.Write([]byte(`{"quotas":{"resources":[]}}`))
	})

	actual, err := quotas.List(serviceClient(), quotas.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(actual))
}

func TestListInvalidResponse(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/quotas", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"quotas":`))
	})

	actual, err := quotas.List(serviceClient(), quotas.ListOpts{})
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

	th.Mux.HandleFunc("/project-id/quotas", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0002","error_msg":"Invalid parameter."}`))
	})

	actual, err := quotas.List(serviceClient(), quotas.ListOpts{})
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
