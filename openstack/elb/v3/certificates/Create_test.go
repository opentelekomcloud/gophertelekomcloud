package certificates_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/certificates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{"certificate":{"name":"test-certificate","description":"test description","type":"server","domain":"example.com,*.example.org","private_key":"test-key","certificate":"test-certificate-content","project_id":"project-id","admin_state_up":false,"enterprise_project_id":"enterprise-id"}}`)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"certificate":` + certificateJSON + `}`))
	})
	actual, err := certificates.Create(serviceClient(), certificates.CreateOpts{
		Name: "test-certificate", Description: "test description", Type: "server",
		Domain: "example.com,*.example.org", PrivateKey: "test-key", Certificate: "test-certificate-content",
		ProjectID: "project-id", AdminStateUp: pointerto.Bool(false), EnterpriseProjectID: "enterprise-id",
	})
	th.AssertNoErr(t, err)
	assertCertificate(t, actual)
}

func TestCreateMinimal(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{"certificate":{"certificate":"test-certificate-content"}}`)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"certificate":` + certificateJSON + `}`))
	})
	_, err := certificates.Create(serviceClient(), certificates.CreateOpts{Certificate: "test-certificate-content"})
	th.AssertNoErr(t, err)
}

func TestCreateClientCertificate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{"certificate":{"type":"client","certificate":"test-ca-content"}}`)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"certificate":{"id":"ca-id","type":"client","certificate":"test-ca-content","domain":null}}`))
	})
	actual, err := certificates.Create(serviceClient(), certificates.CreateOpts{
		Type: "client", Certificate: "test-ca-content",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "client", actual.Type)
	th.AssertEquals(t, "", actual.PrivateKey)
}

func TestCreateRequiredCertificate(t *testing.T) {
	actual, err := certificates.Create(serviceClient(), certificates.CreateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected required certificate error, got %#v, %v", actual, err)
	}
}

func TestCreateErrors(t *testing.T) {
	for _, status := range []int{http.StatusBadRequest, http.StatusCreated, http.StatusOK} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				_, _ = w.Write([]byte(`{"certificate":`))
			})
			actual, err := certificates.Create(serviceClient(), certificates.CreateOpts{Certificate: "test-content"})
			if err == nil || actual != nil {
				t.Fatalf("expected error and nil certificate, got %#v, %v", actual, err)
			}
		})
	}
}
