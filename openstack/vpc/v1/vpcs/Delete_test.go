package vpcs_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs/vpc-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	th.AssertNoErr(t, vpcs.Delete(serviceClient(), "vpc-id"))
}

func TestDeleteError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/vpcs/vpc-id", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0109","error_msg":"Router is being used and cannot be deleted."}`))
	})

	if err := vpcs.Delete(serviceClient(), "vpc-id"); err == nil {
		t.Fatal("expected request error")
	}
}
