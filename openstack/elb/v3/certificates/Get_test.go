package certificates_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/certificates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/certificates/certificate-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		_, _ = w.Write([]byte(`{"certificate":` + certificateJSON + `}`))
	})
	actual, err := certificates.Get(serviceClient(), "certificate-id")
	th.AssertNoErr(t, err)
	assertCertificate(t, actual)
}

func TestGetNullDomain(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/certificates/certificate-id", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"certificate":{"id":"certificate-id","domain":null}}`))
	})
	actual, err := certificates.Get(serviceClient(), "certificate-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", actual.Domain)
}

func TestGetErrors(t *testing.T) {
	for _, status := range []int{http.StatusNotFound, http.StatusOK} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/certificates/certificate-id", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				_, _ = w.Write([]byte(`{"certificate":`))
			})
			actual, err := certificates.Get(serviceClient(), "certificate-id")
			if err == nil || actual != nil {
				t.Fatalf("expected error and nil certificate, got %#v, %v", actual, err)
			}
		})
	}
}

func TestGetInvalidID(t *testing.T) {
	actual, err := certificates.Get(serviceClient(), "../listeners/listener-id")
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
