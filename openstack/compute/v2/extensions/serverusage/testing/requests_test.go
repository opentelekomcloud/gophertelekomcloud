package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/serverusage"
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

		_, _ = fmt.Fprint(w, ServerWithUsageExtResult)
	})

	type serverUsageExt struct {
		servers.Server
		serverusage.UsageExt
	}
	var serverWithUsageExt serverUsageExt
	err := servers.GetInto(fake.ServiceClient(), "d650a0ce-17c3-497d-961a-43c4af80998a", &serverWithUsageExt)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, serverWithUsageExt.LaunchedAt, time.Date(2018, 07, 27, 9, 15, 55, 0, time.UTC))
	th.AssertEquals(t, serverWithUsageExt.TerminatedAt, time.Time{})
	th.AssertEquals(t, serverWithUsageExt.Created, time.Date(2018, 07, 27, 9, 15, 48, 0, time.UTC))
	th.AssertEquals(t, serverWithUsageExt.Updated, time.Date(2018, 07, 27, 9, 15, 55, 0, time.UTC))
	th.AssertEquals(t, serverWithUsageExt.ID, "d650a0ce-17c3-497d-961a-43c4af80998a")
	th.AssertEquals(t, serverWithUsageExt.Name, "test_instance")
	th.AssertEquals(t, serverWithUsageExt.Status, "ACTIVE")
	th.AssertEquals(t, serverWithUsageExt.UserID, "0f2f3822679e4b3ea073e5d1c6ed5f02")
	th.AssertEquals(t, serverWithUsageExt.TenantID, "424e7cf0243c468ca61732ba45973b3e")
}
