package bandwidths_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/bandwidths"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths/bandwidth-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		_, _ = w.Write([]byte(`{"bandwidth":` + bandwidthJSON + `}`))
	})

	actual, err := bandwidths.Update(serviceClient(), "bandwidth-id", bandwidths.UpdateOpts{
		Name: "bandwidth-6f78",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "bandwidth-id", actual.ID)
	th.AssertEquals(t, "bandwidth-6f78", actual.Name)
}

func TestUpdateError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths/bandwidth-id", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0001","error_msg":"Bad request."}`))
	})

	actual, err := bandwidths.Update(serviceClient(), "bandwidth-id", bandwidths.UpdateOpts{Size: 10})
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
