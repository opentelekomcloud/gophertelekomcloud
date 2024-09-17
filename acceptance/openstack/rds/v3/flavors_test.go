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
		DatabaseName: "PostgreSQL",
		VersionName:  "14",
	}

	rdsFlavors, err := flavors.ListFlavors(client, listOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, rdsFlavors)

	dataStores, err := flavors.ListDatastores(client, "PostgreSQL")
	th.AssertNoErr(t, err)
	tools.PrintResource(t, dataStores)

	storageTypes, err := flavors.ListStorageTypes(client, flavors.ListStorageTypesOpts{
		DatabaseName: "PostgreSQL",
		VersionName:  "15",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, storageTypes)
}
