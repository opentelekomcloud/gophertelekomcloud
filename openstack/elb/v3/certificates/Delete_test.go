package certificates_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/certificates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/certificates/certificate-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodDelete)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.AssertEquals(t, int64(0), r.ContentLength)
		w.WriteHeader(http.StatusNoContent)
	})
	th.AssertNoErr(t, certificates.Delete(serviceClient(), "certificate-id"))
}

func TestDeleteErrors(t *testing.T) {
	for _, status := range []int{http.StatusNotFound, http.StatusConflict} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/certificates/certificate-id", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
			})
			if err := certificates.Delete(serviceClient(), "certificate-id"); err == nil {
				t.Fatal("expected deletion error")
			}
		})
	}
}

func TestDeleteInvalidID(t *testing.T) {
	if err := certificates.Delete(serviceClient(), "../listeners/listener-id"); err == nil {
		t.Fatal("expected invalid ID error")
	}
}
