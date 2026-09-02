package bandwidths_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/bandwidths"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		if err := r.ParseForm(); err != nil {
			t.Fatalf("parse form failed: %s", err)
		}
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			_, _ = fmt.Fprintf(w, `{"bandwidths": [%s]}`, bandwidthJSON)
		case "bandwidth-id":
			_, _ = w.Write([]byte(`{"bandwidths": []}`))
		default:
			t.Fatalf("unexpected marker: %q", marker)
		}
	})

	actual, err := bandwidths.List(serviceClient(), bandwidths.ListOpts{Limit: 2})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(actual))
	th.AssertEquals(t, "bandwidth-id", actual[0].ID)
}

func TestListError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0001","error_msg":"Bad request."}`))
	})

	actual, err := bandwidths.List(serviceClient(), bandwidths.ListOpts{})
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
