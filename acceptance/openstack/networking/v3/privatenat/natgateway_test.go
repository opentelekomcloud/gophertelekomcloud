package privatenat

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v3/privatenat/natgateway"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPrivateNatGatewayLifecycle(t *testing.T) {
	client, err := clients.NewNatV3Client()
	th.AssertNoErr(t, err)

	natGateway := createPrivateNatGateway(t, client)
	t.Cleanup(func() {
		deletePrivateNatGateway(t, client, natGateway.Id)
	})

	getResponse, err := natgateway.Get(client, natGateway.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, natGateway.Name, getResponse.Gateway.Name)

	listResponse, err := natgateway.List(client, natgateway.ListGatewaysQueryParams{
		Id: []string{natGateway.Id},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(listResponse.Gateways))

	updateResponse, err := natgateway.Update(client, natGateway.Id, natgateway.UpdateGatewayOpts{
		Description: "updated",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "updated", updateResponse.Gateway.Description)
}

func createPrivateNatGateway(t *testing.T, client *golangsdk.ServiceClient) *natgateway.PrivateNATGateway {
	t.Logf("Attempting to create Private Nat Gateway")
	natGatewayName := tools.RandomString("create-private-nat-", 8)

	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	if networkID == "" {
		t.Skip("OS_NETWORK_ID is missing but test requires using existing network")
	}

	createNatGatewayOpts := natgateway.CreateGatewayOpts{
		Name:        natGatewayName,
		Description: "created",
		Spec:        "Small",
		DownlinkVpcs: []natgateway.DownlinkVpcOption{
			{
				VirSubnetID: networkID,
			},
		},
	}

	createResp, err := natgateway.Create(client, createNatGatewayOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "created", createResp.Gateway.Description)

	t.Logf("Created Nat Gateway: %s", createResp.Gateway.Id)

	return &createResp.Gateway
}

func deletePrivateNatGateway(t *testing.T, client *golangsdk.ServiceClient, natGatewayID string) {
	t.Logf("Attempting to delete Private Nat Gateway: %s", natGatewayID)

	th.AssertNoErr(t, natgateway.Delete(client, natGatewayID))

	t.Logf("Private Nat Gateway is deleted: %s", natGatewayID)
}
