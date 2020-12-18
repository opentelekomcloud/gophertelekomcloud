package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/flavors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDdsFlavorsList(t *testing.T) {
	client, err := clients.NewDdsV3Client()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	listFlavorOpts := flavors.ListOpts{
		Region: cc.RegionName,
	}
	allPages, err := flavors.List(client, listFlavorOpts).AllPages()
	th.AssertNoErr(t, err)
	flavorsList, err := flavors.ExtractFlavors(allPages)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, flavorsList)
}
