package extensions

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/snatrules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSnatRuleLifeCycle(t *testing.T) {
	client, err := clients.NewNatV2Client()
	th.AssertNoErr(t, err)

	natGateway := createNatGateway(t, client)
	defer deleteNatGateway(t, client, natGateway.ID)

	elasticIp := createEip(t)
	defer deleteEip(t, elasticIp.ID)

	t.Logf("Attempting to create SNAT rule")
	createOpts := snatrules.CreateOpts{
		NatGatewayID: natGateway.ID,
		NetworkID:    natGateway.InternalNetworkID,
		FloatingIPID: elasticIp.ID,
		SourceType:   0,
	}
	snatRule, err := snatrules.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created SNAT rule: %s", snatRule.ID)

	defer func() {
		t.Logf("Attempting to delete SNAT rule: %s", snatRule.ID)
		err = snatrules.Delete(client, snatRule.ID).Err
		th.AssertNoErr(t, err)
		t.Logf("Deleted SNAT rule: %s", snatRule.ID)
	}()

	newSnatRule, err := snatrules.Get(client, snatRule.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.NatGatewayID, newSnatRule.NatGatewayID)
}
