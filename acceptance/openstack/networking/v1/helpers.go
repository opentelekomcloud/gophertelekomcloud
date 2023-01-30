package v1

import (
	"fmt"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/eips"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/subnets"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// CreateNetwork - create VPC+Subnet and returns active subnet instance
func CreateNetwork(t *testing.T, prefix, az string) *subnets.Subnet {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)
	vpc, err := vpcs.Create(client, vpcs.CreateOpts{
		Name: tools.RandomString(prefix, 4),
		CIDR: "192.168.0.0/16",
	}).Extract()
	th.AssertNoErr(t, err)
	enableDHCP := true
	subnet, err := subnets.Create(client, subnets.CreateOpts{
		Name:             tools.RandomString(prefix, 4),
		CIDR:             "192.168.0.0/24",
		DNSList:          []string{"1.1.1.1", "8.8.8.8"},
		GatewayIP:        "192.168.0.1",
		EnableDHCP:       &enableDHCP,
		AvailabilityZone: az,
		VpcID:            vpc.ID,
	}).Extract()
	th.AssertNoErr(t, err)

	err = waitForSubnetToBeActive(client, subnet.ID, 300)
	th.AssertNoErr(t, err)

	return subnet
}

// DeleteNetwork - remove subnet and network
func DeleteNetwork(t *testing.T, subnet *subnets.Subnet) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	err = subnets.Delete(client, subnet.VpcID, subnet.ID).ExtractErr()
	th.AssertNoErr(t, err)
	err = waitForSubnetToBeDeleted(client, subnet.ID, 300)
	th.AssertNoErr(t, err)

	err = vpcs.Delete(client, subnet.VpcID).ExtractErr()
	th.AssertNoErr(t, err)
}

func CreateEip(t *testing.T, client *golangsdk.ServiceClient, bandwidthSize int) *eips.PublicIp {
	t.Logf("Attempting to create eip/bandwidth")
	eipCreateOpts := eips.ApplyOpts{
		IP: eips.PublicIpOpts{
			Type: "5_bgp",
		},
		Bandwidth: eips.BandwidthOpts{
			ShareType: "PER",
			Name:      tools.RandomString("acc-band-", 3),
			Size:      bandwidthSize,
		},
	}

	eip, err := eips.Apply(client, eipCreateOpts).Extract()
	th.AssertNoErr(t, err)

	// wait to be DOWN
	t.Logf("Waiting for eip %s to be active", eip.ID)
	err = waitForEipToActive(client, eip.ID, 600)
	th.AssertNoErr(t, err)

	newEip, err := eips.Get(client, eip.ID).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Created eip/bandwidth: %s", newEip.ID)

	return newEip
}

func DeleteEip(t *testing.T, client *golangsdk.ServiceClient, eipID string) {
	t.Logf("Attempting to delete eip/bandwidth: %s", eipID)

	err := eips.Delete(client, eipID).ExtractErr()
	th.AssertNoErr(t, err)

	// wait to be deleted
	t.Logf("Waitting for eip %s to be deleted", eipID)

	err = waitForEipToDelete(client, eipID, 600)
	th.AssertNoErr(t, err)

	t.Logf("Deleted eip/bandwidth: %s", eipID)
}

func waitForEipToActive(client *golangsdk.ServiceClient, eipID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		eip, err := eips.Get(client, eipID).Extract()
		if err != nil {
			return false, err
		}
		if eip.Status == "DOWN" {
			return true, nil
		}

		return false, nil
	})
}

func waitForEipToDelete(client *golangsdk.ServiceClient, eipID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := eips.Get(client, eipID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}

func createSubnet(t *testing.T, client *golangsdk.ServiceClient, vpcID string) *subnets.Subnet {
	enableDHCP := true
	createSubnetOpts := subnets.CreateOpts{
		Name:        tools.RandomString("acc-subnet-", 3),
		Description: "some description",
		CIDR:        "192.168.20.0/24",
		GatewayIP:   "192.168.20.1",
		EnableDHCP:  &enableDHCP,
		VpcID:       vpcID,
	}
	t.Logf("Attempting to create subnet: %s", createSubnetOpts.Name)

	subnet, err := subnets.Create(client, createSubnetOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, subnet.Description, createSubnetOpts.Description)

	// wait to be active
	t.Logf("Waitting for subnet %s to be active", subnet.ID)
	err = waitForSubnetToBeActive(client, subnet.ID, 600)
	th.AssertNoErr(t, err)
	t.Logf("Created subnet: %v", subnet.ID)

	return subnet
}

func deleteSubnet(t *testing.T, client *golangsdk.ServiceClient, vpcID string, id string) {
	t.Logf("Attempting to delete subnet: %s", id)

	err := subnets.Delete(client, vpcID, id).ExtractErr()
	th.AssertNoErr(t, err)

	t.Logf("Waiting for subnet %s to be deleted", id)
	err = waitForSubnetToBeDeleted(client, id, 60)
	th.AssertNoErr(t, err)

	t.Logf("Deleted subnet: %s", id)
}

func waitForSubnetToBeDeleted(client *golangsdk.ServiceClient, subnetID string, secs int) error {
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

func waitForSubnetToBeActive(client *golangsdk.ServiceClient, subnetID string, secs int) error {
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

func createVpc(t *testing.T, client *golangsdk.ServiceClient) *vpcs.Vpc {
	createOpts := vpcs.CreateOpts{
		Name: tools.RandomString("acc-vpc-", 3),
		CIDR: "192.168.20.0/24",
	}

	t.Logf("Attempting to create vpc: %s", createOpts.Name)

	vpc, err := vpcs.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created vpc: %s", vpc.ID)

	return vpc
}

func deleteVpc(t *testing.T, client *golangsdk.ServiceClient, vpcID string) {
	t.Logf("Attempting to delete vpc: %s", vpcID)

	err := vpcs.Delete(client, vpcID).ExtractErr()
	th.AssertNoErr(t, err)

	t.Logf("Deleted vpc: %s", vpcID)
}

func createEipTags(t *testing.T, client *golangsdk.ServiceClient, eipID string, eipTags []tags.ResourceTag) {
	err := tags.Create(client, "publicips", eipID, eipTags).ExtractErr()
	th.AssertNoErr(t, err)
}

func deleteEipTags(t *testing.T, client *golangsdk.ServiceClient, eipID string, eipTags []tags.ResourceTag) {
	err := tags.Delete(client, "publicips", eipID, eipTags).ExtractErr()
	th.AssertNoErr(t, err)
}
