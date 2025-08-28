package privatenat

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v3/privatenat/dnatrules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPrivateNatDnatRulesLifecycle(t *testing.T) {
	client, err := clients.NewNatV3Client()
	th.AssertNoErr(t, err)
	transitIpId := clients.EnvOS.GetEnv("TRANSIT_IP_ID")
	networkInterfaceId := clients.EnvOS.GetEnv("NIC_ID")
	if transitIpId == "" || networkInterfaceId == "" {
		t.Skip("OS_TRANSIT_IP_ID or OS_NIC_ID is missing but test requires using existing network")
	}

	natGateway := createPrivateNatGateway(t, client)
	t.Cleanup(func() {
		deletePrivateNatGateway(t, client, natGateway.Id)
	})

	createOpts := dnatrules.CreatePrivateDnatOpts{
		Description:        "created",
		GatewayId:          natGateway.Id,
		TransitIpId:        transitIpId,
		NetworkInterfaceId: networkInterfaceId,
	}
	dnatCreateResp, err := dnatrules.Create(client, createOpts)
	th.AssertNoErr(t, err)
	ruleId := dnatCreateResp.DnatRule.Id
	t.Logf("Created DNAT rule: %s", ruleId)
	t.Cleanup(func() {
		t.Logf("Deleting DNAT rule: %s", ruleId)
		th.AssertNoErr(t, dnatrules.Delete(client, ruleId))
		t.Logf("Deleted DNAT rule: %s", ruleId)
	})

	getResponse, err := dnatrules.Get(client, ruleId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "created", getResponse.DnatRule.Description)

	listResponse, err := dnatrules.List(client, dnatrules.ListDnatRulesQueryParams{
		Id: []string{ruleId},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(listResponse.DnatRules))

	updateResponse, err := dnatrules.Update(client, ruleId, dnatrules.UpdatePrivateDnatOpts{
		Description: "updated",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "updated", updateResponse.DnatRule.Description)
}
