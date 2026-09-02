package publicips_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/publicips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips/publicip-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	th.AssertNoErr(t, publicips.Delete(serviceClient(), "publicip-id"))
}

func TestDeleteError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips/publicip-id", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0601","error_msg":"EIP is in use and cannot be released."}`))
	})

	if err := publicips.Delete(serviceClient(), "publicip-id"); err == nil {
		t.Fatal("expected request error")
	}
}
