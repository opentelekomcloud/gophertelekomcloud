package privateips_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/privateips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/privateips/private-ip-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	th.AssertNoErr(t, privateips.Delete(serviceClient(), "private-ip-id"))
}

func TestDeleteError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/privateips/private-ip-id", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0208","error_msg":"Private IP is in use."}`))
	})

	if err := privateips.Delete(serviceClient(), "private-ip-id"); err == nil {
		t.Fatal("expected request error")
	}
}

func TestDeleteInvalidID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	if err := privateips.Delete(serviceClient(), "../vpcs/vpc-id"); err == nil {
		t.Fatal("expected invalid ID error")
	}
}
