package bandwidths_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/bandwidths"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRemoveEip(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths/bandwidth-id/remove", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"bandwidth":{"publicip_info":[{"publicip_id":"publicip-id"}],"charge_mode":"traffic","size":10}}`)
		w.WriteHeader(http.StatusOK)
	})

	err := bandwidths.RemoveEip(serviceClient(), "bandwidth-id", bandwidths.RemoveEipOpts{
		PublicipInfo: []bandwidths.RemovePublicIPInfo{{PublicipId: "publicip-id"}},
		ChargeMode:   "traffic",
		Size:         10,
	})
	th.AssertNoErr(t, err)
}

func TestRemoveEipError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths/bandwidth-id/remove", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0001","error_msg":"Bad request."}`))
	})

	err := bandwidths.RemoveEip(serviceClient(), "bandwidth-id", bandwidths.RemoveEipOpts{
		PublicipInfo: []bandwidths.RemovePublicIPInfo{{PublicipId: "publicip-id"}},
		ChargeMode:   "traffic",
		Size:         10,
	})
	if err == nil {
		t.Fatal("expected request error")
	}
}
