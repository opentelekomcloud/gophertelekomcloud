package zones_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/zones"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGetDNSSEC(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/zones/zone-id/dnssec", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(dnssecEnabledJSON))
	})
	actual, err := zones.GetDNSSEC(serviceClient(), "zone-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "ENABLE", actual.Status)
	th.AssertEquals(t, 55451, actual.KeyTag)
	th.AssertEquals(t, 257, actual.Flag)
	th.AssertEquals(t, 13, actual.SignatureType)
	th.AssertEquals(t, "example.org. 3600 IN DS 55451 13 2 be51f11d7e09131ebe9c0000000000000000000000000000000000000000abcd", actual.DSRecord)
}

func TestGetDNSSECDisabled(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/zones/zone-id/dnssec", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(dnssecDisabledJSON))
	})
	actual, err := zones.GetDNSSEC(serviceClient(), "zone-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "DISABLE", actual.Status)
	th.AssertEquals(t, "", actual.DSRecord)
}

func TestGetDNSSECInvalidID(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	actual, err := zones.GetDNSSEC(serviceClient(), "../zones/other-id")
	if err == nil || actual != nil {
		t.Fatalf("expected nil response and invalid ID error, got %#v, %v", actual, err)
	}
}
