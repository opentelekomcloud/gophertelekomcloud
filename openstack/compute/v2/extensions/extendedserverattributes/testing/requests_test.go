package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/extendedserverattributes"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestServerWithUsageExt(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/d650a0ce-17c3-497d-961a-43c4af80998a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		_, _ = fmt.Fprint(w, ServerWithAttributesExtResult)
	})

	var serverWithAttributesExt struct {
		servers.Server
		extendedserverattributes.ServerAttributesExt
	}

	err := servers.GetInto(fake.ServiceClient(), "d650a0ce-17c3-497d-961a-43c4af80998a", &serverWithAttributesExt)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, serverWithAttributesExt.Host, "compute01")
	th.AssertEquals(t, serverWithAttributesExt.InstanceName, "instance-00000001")
	th.AssertEquals(t, serverWithAttributesExt.HypervisorHostname, "compute01")
}
