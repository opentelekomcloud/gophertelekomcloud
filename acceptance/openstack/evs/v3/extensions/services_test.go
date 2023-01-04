package extensions

import (
	"testing"

	services2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/extensions/services"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestServicesList(t *testing.T) {
	t.Skip("The API does not exist or has not been published in the environment")

	blockClient, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	allServices, err := services2.List(blockClient, services2.ListOpts{})
	th.AssertNoErr(t, err)

	for _, service := range allServices {
		tools.PrintResource(t, service)
	}
}
