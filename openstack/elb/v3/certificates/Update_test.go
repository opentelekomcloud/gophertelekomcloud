package certificates_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/certificates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/certificates/certificate-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{"certificate":{"name":"test-certificate","description":"","domain":"example.com,*.example.org","private_key":"test-key","certificate":"test-certificate-content"}}`)
		_, _ = w.Write([]byte(`{"certificate":` + certificateJSON + `}`))
	})
	actual, err := certificates.Update(serviceClient(), "certificate-id", certificates.UpdateOpts{
		Name: "test-certificate", Description: pointerto.String(""), Domain: "example.com,*.example.org",
		PrivateKey: "test-key", Certificate: "test-certificate-content",
	})
	th.AssertNoErr(t, err)
	assertCertificate(t, actual)
}

func TestUpdateEmptyOptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/certificates/certificate-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{"certificate":{}}`)
		_, _ = w.Write([]byte(`{"certificate":` + certificateJSON + `}`))
	})
	_, err := certificates.Update(serviceClient(), "certificate-id", certificates.UpdateOpts{})
	th.AssertNoErr(t, err)
}

func TestUpdateErrors(t *testing.T) {
	for _, status := range []int{http.StatusBadRequest, http.StatusConflict, http.StatusOK} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/certificates/certificate-id", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				_, _ = w.Write([]byte(`{"certificate":`))
			})
			actual, err := certificates.Update(serviceClient(), "certificate-id", certificates.UpdateOpts{Name: "updated"})
			if err == nil || actual != nil {
				t.Fatalf("expected error and nil certificate, got %#v, %v", actual, err)
			}
		})
	}
}

func TestUpdateInvalidID(t *testing.T) {
	actual, err := certificates.Update(serviceClient(), "../listeners/listener-id", certificates.UpdateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
