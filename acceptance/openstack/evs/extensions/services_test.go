package extensions

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/extensions/services"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestServicesList(t *testing.T) {
	clients.RequireAdmin(t)

	blockClient, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	allPages, err := services.List(blockClient, services.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)

	allServices, err := services.ExtractServices(allPages)
	th.AssertNoErr(t, err)

	for _, service := range allServices {
		tools.PrintResource(t, service)
	}
}
