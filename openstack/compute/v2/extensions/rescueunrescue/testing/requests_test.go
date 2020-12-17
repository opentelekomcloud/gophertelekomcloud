package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/rescueunrescue"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestRescue(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/3f54d05f-3430-4d80-aa07-63e6af9e2488/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, RescueRequest)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, RescueResult)
	})

	s, err := rescueunrescue.Rescue(fake.ServiceClient(), "3f54d05f-3430-4d80-aa07-63e6af9e2488", rescueunrescue.RescueOpts{
		AdminPass:      "aUPtawPzE9NU",
		RescueImageRef: "115e5c5b-72f0-4a0a-9067-60706545248c",
	}).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, "aUPtawPzE9NU", s)
}

func TestUnrescue(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/3f54d05f-3430-4d80-aa07-63e6af9e2488/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, UnrescueRequest)

		w.WriteHeader(http.StatusAccepted)
	})

	err := rescueunrescue.Unrescue(fake.ServiceClient(), "3f54d05f-3430-4d80-aa07-63e6af9e2488").ExtractErr()
	th.AssertNoErr(t, err)
}
