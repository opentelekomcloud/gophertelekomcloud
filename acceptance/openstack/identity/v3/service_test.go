//go:build acceptance
// +build acceptance

package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/services"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestServicesList(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	listOpts := services.ListOpts{
		ServiceType: "identity",
	}

	allPages, err := services.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allServices, err := services.ExtractServices(allPages)
	th.AssertNoErr(t, err)

	for _, service := range allServices {
		tools.PrintResource(t, service)
	}

}
