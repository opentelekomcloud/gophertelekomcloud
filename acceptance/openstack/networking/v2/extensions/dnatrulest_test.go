package extensions

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/dnatrules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDnatRuleLifeCycle(t *testing.T) {
	client, err := clients.NewNatV2Client()
	th.AssertNoErr(t, err)

	natGateway := createNatGateway(t, client)
	defer deleteNatGateway(t, client, natGateway.ID)

	elasticIp := createEip(t)
	defer deleteEip(t, elasticIp.ID)

	t.Logf("Attempting to create DNAT rule")
	allServicePorts := 0
	createOpts := dnatrules.CreateOpts{
		NatGatewayID:        natGateway.ID,
		InternalServicePort: &allServicePorts,
		PrivateIp:           "192.168.1.100",
		FloatingIpID:        elasticIp.ID,
		ExternalServicePort: &allServicePorts,
		Protocol:            "any",
	}
	dnatRule, err := dnatrules.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created DNAT rule: %s", dnatRule.ID)

	defer func() {
		t.Logf("Attempting to delete DNAT rule: %s", dnatRule.ID)
		err = dnatrules.Delete(client, dnatRule.ID).Err
		th.AssertNoErr(t, err)
		t.Logf("Deleted DNAT rule: %s", dnatRule.ID)
	}()

	newDnatRule, err := dnatrules.Get(client, dnatRule.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.NatGatewayID, newDnatRule.NatGatewayID)
}
