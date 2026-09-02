package publicips_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/publicips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips/publicip-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"publicip":` + publicIPJSON + `}`))
	})

	actual, err := publicips.Get(serviceClient(), "publicip-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "publicip-id", actual.ID)
	th.AssertEquals(t, "192.168.10.5", actual.PrivateIpAddress)
	th.AssertEquals(t, "port-id", actual.PortId)
	th.AssertEquals(t, "bandwidth-id", actual.BandwidthId)
	th.AssertEquals(t, "PER", actual.BandwidthShareType)
	th.AssertEquals(t, "bandwidth-test", actual.BandwidthName)
	th.AssertEquals(t, "order-id", actual.Profile.OrderID)
	th.AssertEquals(t, "eu-de", actual.Profile.RegionID)
	th.AssertDeepEquals(t, []string{"share"}, actual.AllowShareBandwidthTypes)
	th.AssertDeepEquals(t, []string{"key=value"}, actual.Tags)
}

func TestGetZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips/publicip-id", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"publicip":{"id":"publicip-id","profile":{},"bandwidth_size":0}}`))
	})

	actual, err := publicips.Get(serviceClient(), "publicip-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "publicip-id", actual.ID)
	th.AssertEquals(t, 0, actual.BandwidthSize)
	th.AssertEquals(t, 0, actual.IPVersion)
	th.AssertEquals(t, "", actual.PortId)
	th.AssertEquals(t, publicips.Profile{}, actual.Profile)
	if actual.Tags != nil {
		t.Fatalf("expected nil Tags, got %#v", actual.Tags)
	}
}

func TestGetError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/project-id/publicips/publicip-id", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error_code":"VPC.0601","error_msg":"EIP does not exist."}`))
	})

	actual, err := publicips.Get(serviceClient(), "publicip-id")
	if err == nil {
		t.Fatal("expected request error")
	}
	if actual != nil {
		t.Fatalf("expected nil response, got %#v", actual)
	}
}
