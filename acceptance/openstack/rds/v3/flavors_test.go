package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/flavors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFlavorsList(t *testing.T) {
	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	listOpts := flavors.ListOpts{
		VersionName: "10",
	}
	allFlavorPages, err := flavors.ListFlavors(client, listOpts, "PostgreSQL").AllPages()
	th.AssertNoErr(t, err)

	rdsFlavors, err := flavors.ExtractDbFlavors(allFlavorPages)
	th.AssertNoErr(t, err)

	for _, rds := range rdsFlavors {
		tools.PrintResource(t, rds)
	}
}
