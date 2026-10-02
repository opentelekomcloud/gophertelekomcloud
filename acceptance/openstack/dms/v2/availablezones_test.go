package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/availablezones"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAvailableZonesGet(t *testing.T) {
	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	az, err := availablezones.Get(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, az)
	th.AssertEquals(t, cc.RegionName, az.RegionID)
	th.AssertEquals(t, true, len(az.AvailableZones) > 0)
	for _, zone := range az.AvailableZones {
		th.AssertEquals(t, true, zone.ID != "")
		th.AssertEquals(t, true, zone.Code != "")
	}
}
