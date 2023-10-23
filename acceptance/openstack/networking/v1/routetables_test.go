package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/routetables"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/subnets"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRouteTablesList(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	listOpts := routetables.ListOpts{}
	routeTablesList, err := routetables.List(client, listOpts)
	th.AssertNoErr(t, err)

	for _, rtb := range routeTablesList {
		tools.PrintResource(t, rtb)
	}
}

func TestRouteTablesLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	// Create first vpc
	createOpts1 := vpcs.CreateOpts{
		Name: tools.RandomString("acc-vpc-1-", 3),
		CIDR: "192.168.0.0/16",
	}

	t.Logf("Attempting to create first vpc: %s", createOpts1.Name)

	vpc1, err := vpcs.Create(client, createOpts1).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created first vpc: %s", vpc1.ID)

	createSubnetOpts1 := subnets.CreateOpts{
		Name:        tools.RandomString("acc-subnet-1-1-", 3),
		Description: "some description 1",
		CIDR:        "192.168.0.0/24",
		GatewayIP:   "192.168.0.1",
		VpcID:       vpc1.ID,
	}
	t.Logf("Attempting to create first subnet: %s", createSubnetOpts1.Name)

	subnet1, err := subnets.Create(client, createSubnetOpts1).Extract()
	th.AssertNoErr(t, err)

	createSubnetOpts2 := subnets.CreateOpts{
		Name:        tools.RandomString("acc-subnet-2-1-", 3),
		Description: "some description 2",
		CIDR:        "192.168.10.0/24",
		GatewayIP:   "192.168.10.1",
		VpcID:       vpc1.ID,
	}
	t.Logf("Attempting to create second subnet: %s", createSubnetOpts2.Name)

	subnet2, err := subnets.Create(client, createSubnetOpts2).Extract()
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		deleteSubnet(t, client, subnet1.VpcID, subnet1.ID)
		deleteSubnet(t, client, subnet2.VpcID, subnet2.ID)
		deleteVpc(t, client, vpc1.ID)
	})

	// Create second vpc
	createOpts2 := vpcs.CreateOpts{
		Name: tools.RandomString("acc-vpc-2-", 3),
		CIDR: "172.16.0.0/16",
	}

	t.Logf("Attempting to create first vpc: %s", createOpts2.Name)

	vpc2, err := vpcs.Create(client, createOpts2).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created second vpc: %s", vpc2.ID)

	createSubnetOpts3 := subnets.CreateOpts{
		Name:        tools.RandomString("acc-subnet-1-2-", 3),
		Description: "some description 3",
		CIDR:        "172.16.10.0/24",
		GatewayIP:   "172.16.10.1",
		VpcID:       vpc2.ID,
	}
	t.Logf("Attempting to create second subnet: %s", createSubnetOpts3.Name)

	subnet3, err := subnets.Create(client, createSubnetOpts3).Extract()
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		deleteSubnet(t, client, subnet3.VpcID, subnet3.ID)
		deleteVpc(t, client, vpc2.ID)
	})

	createRtbOpts := routetables.CreateOpts{
		Name:        tools.RandomString("acc-rtb-", 3),
		Description: "route table",
		VpcID:       vpc1.ID,
	}
	rtb, err := routetables.Create(client, createRtbOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		routetables.Delete(client, rtb.ID)
	})

	getRtb, err := routetables.Get(client, rtb.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getRtb.Description, "route table")

	// updateOpts := &vpcs.UpdateOpts{
	// 	Name: tools.RandomString("acc-vpc-upd-", 3),
	// }
	//
	// _, err = vpcs.Update(client, vpc.ID, updateOpts).Extract()
	// th.AssertNoErr(t, err)
	//
	// newVpc, err := vpcs.Get(client, vpc.ID).Extract()
	// th.AssertNoErr(t, err)
	//
	// tools.PrintResource(t, newVpc)
}
