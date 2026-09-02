package bandwidths_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/bandwidths"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAddEip(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths/bandwidth-id/insert", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"bandwidth":{"publicip_info":[{"publicip_id":"publicip-id","publicip_type":"5_bgp"}]}}`)
		_, _ = w.Write([]byte(`{"bandwidth":` + sharedBandwidthJSON + `}`))
	})

	actual, err := bandwidths.AddEip(serviceClient(), "bandwidth-id", bandwidths.AddEipOpts{
		PublicipInfo: []bandwidths.InsertPublicIPInfo{
			{PublicipId: "publicip-id", PublicipType: "5_bgp"},
		},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "bandwidth-id", actual.ID)
	th.AssertEquals(t, 1, len(actual.PublicipInfo))
	th.AssertEquals(t, "publicip-id", actual.PublicipInfo[0].PublicipId)
}

func TestAddEipError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths/bandwidth-id/insert", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0001","error_msg":"Bad request."}`))
	})

	actual, err := bandwidths.AddEip(serviceClient(), "bandwidth-id", bandwidths.AddEipOpts{
		PublicipInfo: []bandwidths.InsertPublicIPInfo{{PublicipId: "publicip-id"}},
	})
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
