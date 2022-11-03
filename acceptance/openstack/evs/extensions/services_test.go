package extensions

import (
	"testing"

	services2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/extensions/services"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestServicesList(t *testing.T) {
	clients.RequireAdmin(t)

	blockClient, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	allPages, err := services2.List(blockClient, services2.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)

	allServices, err := services2.ExtractServices(allPages)
	th.AssertNoErr(t, err)

	for _, service := range allServices {
		tools.PrintResource(t, service)
	}
}
