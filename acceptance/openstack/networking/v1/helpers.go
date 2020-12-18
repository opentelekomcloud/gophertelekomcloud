package v1

import (
	"fmt"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/subnets"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func WaitForSubnetToDelete(client *golangsdk.ServiceClient, subnetID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := subnets.Get(client, subnetID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}

func WaitForSubnetToActive(client *golangsdk.ServiceClient, subnetID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		n, err := subnets.Get(client, subnetID).Extract()
		if err != nil {
			return false, err
		}

		if n.Status == "ACTIVE" {
			return true, nil
		}

		// If subnet status is other than Active, send error
		if n.Status == "DOWN" || n.Status == "ERROR" {
			return false, fmt.Errorf("subnet status: '%s'", n.Status)
		}

		return false, nil
	})
}

// CreateNetwork - create VPC+Subnet and returns active subnet instace
func CreateNetwork(t *testing.T, prefix, az string) *subnets.Subnet {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)
	vpc, err := vpcs.Create(client, vpcs.CreateOpts{
		Name: tools.RandomString(prefix, 4),
		CIDR: "192.168.0.0/16",
	}).Extract()
	th.AssertNoErr(t, err)
	subnet, err := subnets.Create(client, subnets.CreateOpts{
		Name:             tools.RandomString(prefix, 4),
		CIDR:             "192.168.0.0/24",
		DnsList:          []string{"1.1.1.1", "8.8.8.8"},
		GatewayIP:        "192.168.0.1",
		EnableDHCP:       true,
		AvailabilityZone: az,
		VPC_ID:           vpc.ID,
	}).Extract()
	th.AssertNoErr(t, err)

	err = WaitForSubnetToActive(client, subnet.ID, 300)
	th.AssertNoErr(t, err)

	return subnet
}

// DeleteNetwork - remove subnet and network
func DeleteNetwork(t *testing.T, subnet *subnets.Subnet) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	err = subnets.Delete(client, subnet.VPC_ID, subnet.ID).ExtractErr()
	th.AssertNoErr(t, err)
	err = WaitForSubnetToDelete(client, subnet.ID, 300)
	th.AssertNoErr(t, err)

	err = vpcs.Delete(client, subnet.VPC_ID).ExtractErr()
	th.AssertNoErr(t, err)
}
