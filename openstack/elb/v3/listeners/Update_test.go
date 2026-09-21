package listeners_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/listeners/listener-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestJSONRequest(t, r, `{"listener":{"description":"","name":"updated","protection_status":"nonProtection","access_log_customized_headers_config":{"enable":false}}}`)
		_, _ = w.Write([]byte(`{"listener":` + listenerJSON + `}`))
	})
	actual, err := listeners.Update(serviceClient(), "listener-id", listeners.UpdateOpts{
		Description: pointerto.String(""), Name: pointerto.String("updated"),
		ProtectionStatus:                 "nonProtection",
		AccessLogCustomizedHeadersConfig: &listeners.AccessLogCustomizedHeadersOpts{Enable: pointerto.Bool(false)},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "listener-id", actual.ID)
}

func TestUpdateInvalidID(t *testing.T) {
	actual, err := listeners.Update(serviceClient(), "../pools/pool-id", listeners.UpdateOpts{})
	if err == nil || actual != nil {
		t.Fatalf("expected invalid ID error, got %#v, %v", actual, err)
	}
}
