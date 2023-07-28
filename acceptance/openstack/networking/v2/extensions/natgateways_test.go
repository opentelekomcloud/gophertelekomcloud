package extensions

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/natgateways"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestNatGatewaysList(t *testing.T) {
	client, err := clients.NewNatV2Client()
	th.AssertNoErr(t, err)

	listOpts := natgateways.ListOpts{}
	allPages, err := natgateways.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	natGateways, err := natgateways.ExtractNatGateways(allPages)
	th.AssertNoErr(t, err)

	for _, natGateway := range natGateways {
		tools.PrintResource(t, natGateway)
	}
}

func TestNatGatewaysLifeCycle(t *testing.T) {
	client, err := clients.NewNatV2Client()
	th.AssertNoErr(t, err)

	// Create Nat Gateway
	natGateway := createNatGateway(t, client)
	defer func() {
		deleteNatGateway(t, client, natGateway.ID)
	}()

	tools.PrintResource(t, natGateway)

	newNatGw := updateNatGateway(t, client, natGateway.ID)
	tools.PrintResource(t, newNatGw)
}

func createNatGateway(t *testing.T, client *golangsdk.ServiceClient) *natgateways.NatGateway {
	t.Logf("Attempting to create Nat Gateway")
	natGatewayName := tools.RandomString("create-nat-", 8)

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	if vpcID == "" || networkID == "" {
		t.Skip("OS_VPC_ID or OS_NETWORK_ID is missing but test requires using existing network")
	}
	natSmallSpec := "1"

	createNatGatewayOpts := natgateways.CreateOpts{
		Name:              natGatewayName,
		Description:       "some nat gateway for acceptance test",
		Spec:              natSmallSpec,
		RouterID:          vpcID,
		InternalNetworkID: networkID,
	}

	natGateway, err := natgateways.Create(client, createNatGatewayOpts).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Created Nat Gateway: %s", natGateway.ID)

	return &natGateway
}

func deleteNatGateway(t *testing.T, client *golangsdk.ServiceClient, natGatewayID string) {
	t.Logf("Attempting to delete Nat Gateway: %s", natGatewayID)

	err := natgateways.Delete(client, natGatewayID).Err
	th.AssertNoErr(t, err)

	t.Logf("Nat Gateway is deleted: %s", natGatewayID)
}

func updateNatGateway(t *testing.T, client *golangsdk.ServiceClient, natGatewayID string) *natgateways.NatGateway {
	t.Logf("Attempting to update Nat Gateway")

	natGatewayNewName := tools.RandomString("update-nat-", 8)

	updateOpts := natgateways.UpdateOpts{
		Name: natGatewayNewName,
	}

	err := natgateways.Update(client, natGatewayID, updateOpts).Err
	th.AssertNoErr(t, err)

	t.Logf("Nat Gateway successfully updated: %s", natGatewayID)

	gateway, err := natgateways.Get(client, natGatewayID).Extract()
	th.AssertNoErr(t, err)

	return &gateway
}
