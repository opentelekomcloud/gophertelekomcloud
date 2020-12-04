package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/flavors"
)

func TestDdsFlavorsList(t *testing.T) {
	client, err := clients.NewDdsV3Client()
	if err != nil {
		t.Fatalf("Unable to create a DDSv3 client: %s", err)
	}

	listFlavorOpts := flavors.ListOpts{
		Region: clients.EnvOS.GetEnv("OS_REGION_NAME"),
	}
	allPages, err := flavors.List(client, listFlavorOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to get all pages: %s", err)
	}
	flavorsList, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		t.Fatalf("Unable to extract DDS flavors: %s", err)
	}
	tools.PrintResource(t, flavorsList)
}
