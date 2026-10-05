package zones_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/zones"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDisableDNSSEC(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/zones/zone-id/disable-dnssec", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(dnssecDisabledJSON))
	})
	actual, err := zones.DisableDNSSEC(serviceClient(), "zone-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "DISABLE", actual.Status)
	th.AssertEquals(t, 0, actual.KeyTag)
}
