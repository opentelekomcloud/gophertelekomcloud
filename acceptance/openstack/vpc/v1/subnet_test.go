package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/subnets"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createSubnetVPC(t *testing.T, client *golangsdk.ServiceClient) *vpcs.Vpc {
	t.Helper()

	vpc, err := vpcs.Create(client, vpcs.CreateOpts{
		Name: tools.RandomString("subnet-acc-", 3),
		CIDR: "192.168.0.0/16",
	})
	th.AssertNoErr(t, err)
	t.Logf("Created VPC: %s", vpc.ID)

	return vpc
}

func createTestSubnet(t *testing.T, client *golangsdk.ServiceClient, vpcID string) *subnets.Subnet {
	t.Helper()

	createOpts := subnets.CreateOpts{
		Name:        tools.RandomString("subnet-acc-", 3),
		Description: "some description",
		CIDR:        "192.168.20.0/24",
		GatewayIP:   "192.168.20.1",
		EnableDHCP:  pointerto.Bool(true),
		VpcID:       vpcID,
	}

	subnet, err := subnets.Create(client, createOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.Description, subnet.Description)

	t.Logf("Waiting for subnet %s to become active", subnet.ID)
	th.AssertNoErr(t, waitForSubnetActive(client, subnet.ID, 600))
	t.Logf("Created subnet: %s", subnet.ID)

	return subnet
}

func deleteTestSubnet(t *testing.T, client *golangsdk.ServiceClient, vpcID, id string) {
	t.Helper()

	th.AssertNoErr(t, subnets.Delete(client, vpcID, id))

	t.Logf("Waiting for subnet %s to be deleted", id)
	th.AssertNoErr(t, waitForSubnetDeleted(client, id, 600))
	t.Logf("Deleted subnet: %s", id)
}

func waitForSubnetActive(client *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		subnet, err := subnets.Get(client, id)
		if err != nil {
			return false, err
		}
		return subnet.Status == "ACTIVE", nil
	})
}

func waitForSubnetDeleted(client *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := subnets.Get(client, id)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}
		return false, nil
	})
}

func deleteTestVPC(t *testing.T, client *golangsdk.ServiceClient, id string) {
	t.Helper()

	err := golangsdk.WaitFor(600, func() (bool, error) {
		err := vpcs.Delete(client, id)
		if err == nil {
			return true, nil
		}
		if _, ok := err.(golangsdk.ErrDefault404); ok {
			return true, nil
		}
		if _, ok := err.(golangsdk.ErrDefault409); ok {
			return false, nil
		}
		return false, err
	})
	th.AssertNoErr(t, err)
}

func TestSubnetList(t *testing.T) {
	client, err := clients.NewVPCV1Client()
	th.AssertNoErr(t, err)

	vpc := createSubnetVPC(t, client)
	t.Cleanup(func() {
		deleteTestVPC(t, client, vpc.ID)
	})

	subnet := createTestSubnet(t, client, vpc.ID)
	t.Cleanup(func() {
		deleteTestSubnet(t, client, subnet.VpcID, subnet.ID)
	})

	list, err := subnets.List(client, subnets.ListOpts{VpcID: vpc.ID})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(list))
	th.AssertEquals(t, subnet.ID, list[0].ID)
	th.AssertEquals(t, vpc.ID, list[0].VpcID)
}

func TestSubnetLifecycle(t *testing.T) {
	client, err := clients.NewVPCV1Client()
	th.AssertNoErr(t, err)

	vpc := createSubnetVPC(t, client)
	t.Cleanup(func() {
		deleteTestVPC(t, client, vpc.ID)
	})

	subnet := createTestSubnet(t, client, vpc.ID)
	t.Cleanup(func() {
		deleteTestSubnet(t, client, subnet.VpcID, subnet.ID)
	})

	tools.PrintResource(t, subnet)

	emptyDescription := ""
	updateOpts := subnets.UpdateOpts{
		Name:        tools.RandomString("subnet-acc-update-", 3),
		Description: &emptyDescription,
		EnableIpv6:  pointerto.Bool(true),
	}

	t.Logf("Attempting to update subnet %s", subnet.ID)
	_, err = subnets.Update(client, subnet.VpcID, subnet.ID, updateOpts)
	th.AssertNoErr(t, err)

	updated, err := subnets.Get(client, subnet.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, updated.Name)
	th.AssertEquals(t, emptyDescription, updated.Description)
	th.AssertEquals(t, true, updated.EnableIpv6)
	th.AssertNotEquals(t, "", updated.SubnetIDV6)
}
