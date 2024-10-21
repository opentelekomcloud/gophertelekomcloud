package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rms/region"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRegionsList(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	regions, err := region.GetRegions(client, client.DomainID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(regions) > 0)
}
