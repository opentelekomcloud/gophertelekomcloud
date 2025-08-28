package privatenat

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v3/privatenat/snatrules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPrivateNatSnatRulesLifecycle(t *testing.T) {
	client, err := clients.NewNatV3Client()
	th.AssertNoErr(t, err)
	transitIpId := clients.EnvOS.GetEnv("TRANSIT_IP_ID")
	networkId := clients.EnvOS.GetEnv("NETWORK_ID")
	if transitIpId == "" || networkId == "" {
		t.Skip("OS_TRANSIT_IP_ID or OS_NETWORK_ID is missing but test requires using existing network")
	}

	natGateway := createPrivateNatGateway(t, client)
	t.Cleanup(func() {
		deletePrivateNatGateway(t, client, natGateway.Id)
	})

	createOpts := snatrules.CreatePrivateSnatOpts{
		GatewayId:    natGateway.Id,
		VirSubnetId:  networkId,
		Description:  "created",
		TransitIpIds: []string{transitIpId},
	}
	snatCreateResp, err := snatrules.Create(client, createOpts)
	th.AssertNoErr(t, err)
	ruleId := snatCreateResp.SnatRule.Id
	t.Logf("Created SNAT rule: %s", ruleId)
	t.Cleanup(func() {
		t.Logf("Deleting SNAT rule: %s", ruleId)
		th.AssertNoErr(t, snatrules.Delete(client, ruleId))
		t.Logf("Deleted SNAT rule: %s", ruleId)
	})

	getResponse, err := snatrules.Get(client, ruleId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "created", getResponse.SnatRule.Description)

	listResponse, err := snatrules.List(client, snatrules.ListSnatRulesQueryParams{
		Id: []string{ruleId},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(listResponse.SnatRules))

	updateResponse, err := snatrules.Update(client, ruleId, snatrules.UpdatePrivateSnatOpts{
		Description: "updated",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "updated", updateResponse.SnatRule.Description)
}
