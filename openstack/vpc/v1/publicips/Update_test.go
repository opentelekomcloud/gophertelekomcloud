package publicips_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/publicips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdateBindPort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips/publicip-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{"publicip":{"port_id":"port-id","alias":"tom"}}`)
		_, _ = w.Write([]byte(`{"publicip":` + publicIPJSON + `}`))
	})

	actual, err := publicips.Update(serviceClient(), "publicip-id", publicips.UpdateOpts{
		PortId: "port-id",
		Alias:  "tom",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "publicip-id", actual.ID)
	th.AssertEquals(t, "port-id", actual.PortId)
	th.AssertEquals(t, "tom", actual.Alias)
}

func TestUpdateLegacyIPVersion(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips/publicip-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{"publicip":{"ip_version":6}}`)
		_, _ = w.Write([]byte(`{"publicip":{"id":"publicip-id","ip_version":6}}`))
	})

	actual, err := publicips.Update(serviceClient(), "publicip-id", publicips.UpdateOpts{IPVersion: 6})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "publicip-id", actual.ID)
	th.AssertEquals(t, 6, actual.IPVersion)
}

func TestUpdateEmptyOpts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips/publicip-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `{"publicip":{}}`)
		_, _ = w.Write([]byte(`{"publicip":` + publicIPJSON + `}`))
	})

	actual, err := publicips.Update(serviceClient(), "publicip-id", publicips.UpdateOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "publicip-id", actual.ID)
}

func TestUpdateError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips/publicip-id", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0601","error_msg":"Invalid EIP parameter values."}`))
	})

	actual, err := publicips.Update(serviceClient(), "publicip-id", publicips.UpdateOpts{})
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
