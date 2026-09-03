package bandwidths_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/bandwidths"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestJSONRequest(t, r, `{"bandwidth":{"name":"bandwidth123","size":10}}`)
		_, _ = w.Write([]byte(`{"bandwidth":` + sharedBandwidthJSON + `}`))
	})

	actual, err := bandwidths.Create(serviceClient(), bandwidths.CreateOpts{
		Name: "bandwidth123",
		Size: 10,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "bandwidth-id", actual.ID)
	th.AssertEquals(t, "WHOLE", actual.ShareType)
	th.AssertEquals(t, 1, len(actual.PublicipInfo))
	th.AssertEquals(t, "publicip-id", actual.PublicipInfo[0].PublicipId)
}

func TestCreateZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"bandwidth":{"id":"bandwidth-id","size":0}}`))
	})

	actual, err := bandwidths.Create(serviceClient(), bandwidths.CreateOpts{Name: "bandwidth123", Size: 1})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "bandwidth-id", actual.ID)
	th.AssertEquals(t, 0, actual.Size)
	if actual.PublicipInfo != nil {
		t.Fatalf("expected nil PublicipInfo, got %#v", actual.PublicipInfo)
	}
}

func TestCreateError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0001","error_msg":"Bad request."}`))
	})

	actual, err := bandwidths.Create(serviceClient(), bandwidths.CreateOpts{Name: "bandwidth123", Size: 1})
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
