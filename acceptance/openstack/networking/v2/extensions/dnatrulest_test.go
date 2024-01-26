package extensions

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/dnatrules"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/portsecurity"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/ports"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDnatRuleLifeCycle(t *testing.T) {
	dnatLong := os.Getenv("OS_DNAT_RULES_LONG")

	client, err := clients.NewNatV2Client()
	th.AssertNoErr(t, err)

	natGateway := createNatGateway(t, client)
	t.Cleanup(func() {
		deleteNatGateway(t, client, natGateway.ID)
	})

	elasticIp := createEip(t)
	t.Cleanup(func() {
		deleteEip(t, elasticIp.ID)
	})

	t.Logf("Attempting to create DNAT rule, Direct Connect scenario")
	allServicePorts := 0
	createDCOpts := dnatrules.CreateOpts{
		NatGatewayID:        natGateway.ID,
		InternalServicePort: &allServicePorts,
		PrivateIp:           "192.168.1.100",
		FloatingIpID:        elasticIp.ID,
		ExternalServicePort: &allServicePorts,
		Protocol:            "any",
	}
	dnatRule, err := dnatrules.Create(client, createDCOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created DNAT DC rule: %s", dnatRule.ID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete DC DNAT rule: %s", dnatRule.ID)
		err = dnatrules.Delete(client, dnatRule.ID)
		th.AssertNoErr(t, err)
		t.Logf("Deleted DC DNAT rule: %s", dnatRule.ID)
	})

	if dnatLong != "" {
		clientNetwork, err := clients.NewNetworkV2Client()
		th.AssertNoErr(t, err)

		clientCompute, err := clients.NewComputeV1Client()
		th.AssertNoErr(t, err)

		// Get ECSv1 createOpts
		createEcsOpts := openstack.GetCloudServerCreateOpts(t)

		// Create ECSv1 instance
		ecs := openstack.CreateCloudServer(t, clientCompute, createEcsOpts)
		t.Cleanup(func() {
			openstack.DeleteCloudServer(t, clientCompute, ecs.ID)
		})

		type portWithExt struct {
			ports.Port
			portsecurity.PortSecurityExt
		}

		var allPorts []portWithExt

		allPages, err := ports.List(clientNetwork, ports.ListOpts{DeviceID: ecs.ID}).AllPages()
		th.AssertNoErr(t, err)

		err = ports.ExtractPortsInto(allPages, &allPorts)
		th.AssertNoErr(t, err)

		elasticIpSecond := createEip(t)
		t.Cleanup(func() {
			deleteEip(t, elasticIpSecond.ID)
		})

		t.Logf("Attempting to create DNAT rule, VPC scenario")
		createVPCOpts := dnatrules.CreateOpts{
			NatGatewayID:        natGateway.ID,
			InternalServicePort: &allServicePorts,
			FloatingIpID:        elasticIpSecond.ID,
			ExternalServicePort: &allServicePorts,
			Protocol:            "any",
			PortID:              allPorts[0].ID,
		}
		dnatRuleVpc, err := dnatrules.Create(client, createVPCOpts)
		th.AssertNoErr(t, err)
		t.Logf("Created DNAT VPC rule: %s", dnatRuleVpc.ID)

		t.Cleanup(func() {
			t.Logf("Attempting to delete VPC DNAT rule: %s", dnatRuleVpc.ID)
			err = dnatrules.Delete(client, dnatRuleVpc.ID)
			th.AssertNoErr(t, err)
			t.Logf("Deleted VPC DNAT rule: %s", dnatRuleVpc.ID)
		})
	}
	t.Logf("Attempting to get DNAT rule: %s", dnatRule.ID)
	newDnatRule, err := dnatrules.Get(client, dnatRule.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createDCOpts.NatGatewayID, newDnatRule.NatGatewayId)

	t.Logf("Attempting to list DNAT rules")
	listRules, err := dnatrules.List(client, dnatrules.ListOpts{})
	th.AssertNoErr(t, err)
	if dnatLong != "" {
		th.AssertEquals(t, len(listRules), 2)
	} else {
		th.AssertEquals(t, len(listRules), 1)
	}

}
