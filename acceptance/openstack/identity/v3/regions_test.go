package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/regions"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRegionsList(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	listOpts := regions.ListOpts{
		ParentRegionID: "RegionOne",
	}

	allPages, err := regions.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allRegions, err := regions.ExtractRegions(allPages)
	th.AssertNoErr(t, err)

	for _, region := range allRegions {
		tools.PrintResource(t, region)
	}
}

func TestRegionsGet(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	allPages, err := regions.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allRegions, err := regions.ExtractRegions(allPages)
	th.AssertNoErr(t, err)

	region := allRegions[0]
	p, err := regions.Get(client, region.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, p)
}
