package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/flavors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFlavorsList(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	listOpts := flavors.ListOpts{}
	flavorsPages, err := flavors.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	flavorsList, err := flavors.ExtractFlavors(flavorsPages)
	th.AssertNoErr(t, err)

	zeroFlavor, err := flavors.Get(client, flavorsList[0].ID)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, zeroFlavor, &flavorsList[0])
}
