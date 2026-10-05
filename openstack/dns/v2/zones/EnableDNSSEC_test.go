package zones_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/zones"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEnableDNSSEC(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/zones/zone-id/enable-dnssec", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(dnssecEnabledJSON))
	})
	actual, err := zones.EnableDNSSEC(serviceClient(), "zone-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "ENABLE", actual.Status)
	th.AssertEquals(t, 257, actual.Flag)
	th.AssertEquals(t, "/+AJ2gscXUGVea9ofxx5AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==", actual.KSKPublicKey)
}
