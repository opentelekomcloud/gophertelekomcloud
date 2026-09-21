package listeners_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/listeners", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"listener":{"loadbalancer_id":"loadbalancer-id","name":"listener-test","protocol":"HTTPS","protocol_port":443,"insert_headers":{"X-Forwarded-ELB-IP":true,"X-Forwarded-Host":null,"X-Forwarded-Proto":true},"protection_status":"consoleProtection","protection_reason":"managed","access_log_customized_headers_config":{"enable":false,"include_headers":["X-Test"]}}}`)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"listener":` + listenerJSON + `}`))
	})
	actual, err := listeners.Create(serviceClient(), listeners.CreateOpts{
		LoadbalancerID: "loadbalancer-id", Name: "listener-test", Protocol: listeners.ProtocolHTTPS,
		ProtocolPort: 443, InsertHeaders: &listeners.InsertHeaders{
			ForwardedELBIP: pointerto.Bool(true), ForwardedProto: pointerto.Bool(true),
		},
		ProtectionStatus: "consoleProtection", ProtectionReason: "managed",
		AccessLogCustomizedHeadersConfig: &listeners.AccessLogCustomizedHeadersOpts{
			Enable: pointerto.Bool(false), IncludeHeaders: []string{"X-Test"},
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "listener-id", actual.ID)
}
