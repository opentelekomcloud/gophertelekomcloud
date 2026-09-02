package bandwidths_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/bandwidths"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths/bandwidth-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	err := bandwidths.Delete(serviceClient(), "bandwidth-id")
	th.AssertNoErr(t, err)
}

func TestDeleteError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths/bandwidth-id", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0602","error_msg":"Bandwidth does not exist."}`))
	})

	err := bandwidths.Delete(serviceClient(), "bandwidth-id")
	if err == nil {
		t.Fatal("expected request error")
	}
}
