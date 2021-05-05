package ports

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/networks"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateNetwork(t *testing.T, client *golangsdk.ServiceClient) *networks.Network {
	createName := tools.RandomString("network-", 3)
	createOpts := networks.CreateOpts{
		Name: createName,
	}

	network, err := networks.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	return network
}

func DeleteNetwork(t *testing.T, client *golangsdk.ServiceClient, networkID string) {
	err := networks.Delete(client, networkID).ExtractErr()
	th.AssertNoErr(t, err)
}
