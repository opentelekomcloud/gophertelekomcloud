package extensions

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/natgateways"
)

func TestNatGatewaysList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a NetworkingV2 client: %s", err)
	}

	listOpts := natgateways.ListOpts{}
	allPages, err := natgateways.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to fetch Nat Gateways pages: %s", err)
	}
	natGateways, err := natgateways.ExtractNatGateways(allPages)
	if err != nil {
		t.Fatalf("Unable to extract Nat Gateways pages: %s", err)
	}
	for _, natGateway := range natGateways {
		tools.PrintResource(t, natGateway)
	}
}

func TestNatGatewaysLifeCycle(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a NetworkingV2 client: %s", err)
	}

	// Create Nat Gateway
	natGateway, err := createNatGateway(t, client)
	if err != nil {
		t.Fatalf("Unable to create Nat Gateway: %s", err)
	}
	defer deleteNatGateway(t, client, natGateway.ID)

	tools.PrintResource(t, natGateway)

	err = updateNatGateway(t, client, natGateway.ID)
	if err != nil {
		t.Fatalf("Unable to update Nat Gateway: %s", err)
	}
	tools.PrintResource(t, natGateway)

	newNatGateway, err := natgateways.Get(client, natGateway.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get Nat Gateway: %s", err)
	}
	tools.PrintResource(t, newNatGateway)
}

func createNatGateway(t *testing.T, client *golangsdk.ServiceClient) (*natgateways.NatGateway, error) {
	natGatewayName := tools.RandomString("create-nat-", 8)

	routerID := clients.EnvOS.GetEnv("ROUTER_ID")
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	if routerID == "" || networkID == "" {
		t.Skip("OS_ROUTER_ID or OS_NETWORK_ID env vars is missing but test requires using existing network")
	}
	natSmallSpec := "1"

	createNatGatewayOpts := natgateways.CreateOpts{
		Name:              natGatewayName,
		Description:       "some nat gateway for acceptance test",
		Spec:              natSmallSpec,
		RouterID:          routerID,
		InternalNetworkID: networkID,
	}

	natGateway, err := natgateways.Create(client, createNatGatewayOpts).Extract()
	if err != nil {
		return nil, err
	}
	t.Logf("Created Nat Gateway: %s", natGateway.ID)

	return &natGateway, nil
}

func deleteNatGateway(t *testing.T, client *golangsdk.ServiceClient, natGatewayID string) {
	t.Logf("Attempting to delete Nat Gateway: %s", natGatewayID)

	if err := natgateways.Delete(client, natGatewayID).Err; err != nil {
		t.Fatalf("Unable to delete Nat Gateway: %s", err)
	}

	t.Logf("Nat Gateway is deleted: %s", natGatewayID)
}

func updateNatGateway(t *testing.T, client *golangsdk.ServiceClient, natGatewayID string) error {
	t.Logf("Attempting to update Nat Gateway")

	natGatewayNewName := tools.RandomString("update-nat-", 8)

	updateOpts := natgateways.UpdateOpts{
		Name: natGatewayNewName,
	}

	if err := natgateways.Update(client, natGatewayID, updateOpts).Err; err != nil {
		return err
	}
	t.Logf("Nat Gateway successfully updated: %s", natGatewayID)
	return nil
}
