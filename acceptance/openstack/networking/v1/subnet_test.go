package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/subnets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSubnetList(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	vpc := createVpc(t, client)
	defer deleteVpc(t, client, vpc.ID)

	subnet := createSubnet(t, client, vpc.ID)
	defer deleteSubnet(t, client, subnet.VpcID, subnet.ID)

	listOpts := subnets.ListOpts{
		VpcID: vpc.ID,
	}

	filteredSubnets, err := subnets.List(client, listOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(filteredSubnets))
	th.AssertEquals(t, vpc.ID, filteredSubnets[0].VpcID)
}

func TestSubnetsLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	vpc := createVpc(t, client)
	defer deleteVpc(t, client, vpc.ID)

	subnet := createSubnet(t, client, vpc.ID)
	defer deleteSubnet(t, client, subnet.VpcID, subnet.ID)

	tools.PrintResource(t, subnet)

	// Update a subnet
	emptyDescription := ""
	updateOpts := &subnets.UpdateOpts{
		Name:        tools.RandomString("acc-subnet-", 3),
		Description: &emptyDescription,
		EnableIpv6:  pointerto.Bool(true),
	}
	t.Logf("Attempting to update name of subnet to %s", updateOpts.Name)
	_, err = subnets.Update(client, subnet.VpcID, subnet.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	// Query a subnet
	newSubnet, err := subnets.Get(client, subnet.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, newSubnet.Name)
	th.AssertEquals(t, emptyDescription, newSubnet.Description)
	th.AssertEquals(t, true, newSubnet.EnableIpv6)
}
