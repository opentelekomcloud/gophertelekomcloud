package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/maintainwindows"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestMaintainWindowsGet(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	res, err := maintainwindows.Get(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, res)
	th.AssertEquals(t, true, len(res.MaintainWindows) > 0)

	defaults := 0
	for _, w := range res.MaintainWindows {
		th.AssertEquals(t, true, w.ID > 0)
		th.AssertEquals(t, true, w.Begin != "")
		th.AssertEquals(t, true, w.End != "")
		if w.Default {
			defaults++
		}
	}
	th.AssertEquals(t, 1, defaults)
}
