package bandwidths_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/bandwidths"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths/bandwidth-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"bandwidth":` + bandwidthJSON + `}`))
	})

	actual, err := bandwidths.Get(serviceClient(), "bandwidth-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "bandwidth-id", actual.ID)
	th.AssertEquals(t, "bandwidth-6f78", actual.Name)
	th.AssertEquals(t, 5, actual.Size)
	th.AssertEquals(t, "PER", actual.ShareType)
	th.AssertEquals(t, "NORMAL", actual.Status)
	th.AssertEquals(t, "center", actual.PublicBorderGroup)
	th.AssertEquals(t, 1, len(actual.PublicipInfo))
	th.AssertEquals(t, "publicip-id", actual.PublicipInfo[0].PublicipId)
	th.AssertEquals(t, 4, actual.PublicipInfo[0].IPVersion)
}

func TestGetZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths/bandwidth-id", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"bandwidth":{"id":"bandwidth-id","size":0}}`))
	})

	actual, err := bandwidths.Get(serviceClient(), "bandwidth-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "bandwidth-id", actual.ID)
	th.AssertEquals(t, 0, actual.Size)
	th.AssertEquals(t, "", actual.Name)
	if actual.PublicipInfo != nil {
		t.Fatalf("expected nil PublicipInfo, got %#v", actual.PublicipInfo)
	}
}

func TestGetError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/bandwidths/bandwidth-id", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0602","error_msg":"Bandwidth does not exist."}`))
	})

	actual, err := bandwidths.Get(serviceClient(), "bandwidth-id")
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
