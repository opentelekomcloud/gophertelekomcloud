package evpn

import (
	"fmt"
	"os"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evpn/v5/gateway"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/eips"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/subnets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGatewaysList(t *testing.T) {
	client, err := clients.NewEVPNClient()
	th.AssertNoErr(t, err)
	gateways, err := gateway.List(client)
	th.AssertNoErr(t, err)
	for _, gw := range gateways {
		tools.PrintResource(t, gw)
	}
}

func TestGatewaysAZsList(t *testing.T) {
	client, err := clients.NewEVPNClient()
	th.AssertNoErr(t, err)
	azs, err := gateway.ListAZs(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, azs)
}

func TestGatewayVPCLifecycle(t *testing.T) {
	t.Skip("unstable creation of gateway")
	subnetId := os.Getenv("OS_SUBNET_ID")
	vpcId := os.Getenv("OS_VPC_ID")

	if subnetId == "" && vpcId == "" {
		t.Skip("`OS_SUBNET_ID` and `OS_VPC_ID` and needs to be defined for test")
	}

	client, err := clients.NewEVPNClient()
	th.AssertNoErr(t, err)

	clientNetV1, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	clientNetV2, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	subnet, err := subnets.Get(clientNetV2, subnetId).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to CREATE EIP 1")
	eip1, err := createEip(clientNetV1, "eip1")
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Logf("Attempting to DELETE EIP 1: %s", eip1.ID)
		eips.Delete(clientNetV1, eip1.ID)
	})

	t.Logf("Attempting to CREATE EIP 2")
	eip2, err := createEip(clientNetV1, "eip2")
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Logf("Attempting to DELETE EIP 2: %s", eip2.ID)
		eips.Delete(clientNetV1, eip2.ID)
	})

	name := tools.RandomString("acc_evpn_gateway_", 5)
	cOpts := gateway.CreateOpts{
		Name:          name,
		VpcId:         vpcId,
		ConnectSubnet: subnet.NetworkID,
		LocalSubnets:  []string{subnet.CIDR},
		Eip1: &gateway.Eip{
			ID: eip1.ID,
		},
		Eip2: &gateway.Eip{
			ID: eip2.ID,
		},
	}
	// seems eip unavailable for a second after creation so need to wait,
	// otherwise fails with 500 error
	time.Sleep(1 * time.Second)
	t.Logf("Attempting to CREATE Enterprise VPN Gateway: %s", name)
	gw, err := gateway.Create(client, cOpts)
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, WaitForGatewayActive(client, gw.ID, 800))
	th.AssertEquals(t, name, gw.Name)
	t.Cleanup(func() {
		t.Logf("Attempting to DELETE Enterprise VPN Gateway: %s", gw.ID)
		th.AssertNoErr(t, gateway.Delete(client, gw.ID))
		th.AssertNoErr(t, WaitForGatewayDeleted(client, gw.ID, 800))
	})
	t.Logf("Attempting to GET Enterprise VPN Gateway: %s", gw.ID)
	gwGet, err := gateway.Get(client, gw.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, eip1.ID, gwGet.Eip1.ID)
	th.AssertEquals(t, eip2.ID, gwGet.Eip2.ID)

	t.Logf("Attempting to UPDATE Enterprise VPN Gateway: %s", gw.ID)
	gwUpdate, err := gateway.Update(client, gateway.UpdateOpts{
		GatewayID: gw.ID,
		Name:      name + "updated",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name+"updated", gwUpdate.Name)

	t.Logf("Attempting to LIST Enterprise VPN Gateways: %s", gw.ID)
	gws, err := gateway.List(client)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(gws))
}

func createEip(clientNet *golangsdk.ServiceClient, name string) (*eips.PublicIp, error) {
	eip, err := eips.Apply(clientNet, eips.ApplyOpts{
		IP: eips.PublicIpOpts{
			Name: "eip-" + name,
			Type: "5_bgp",
		},
		Bandwidth: eips.BandwidthOpts{
			Name:       "bandwidth-" + name,
			Size:       1,
			ShareType:  "PER",
			ChargeMode: "traffic"},
	}).Extract()
	if err != nil {
		return nil, err
	}
	return eip, nil
}

func WaitForGatewayActive(c *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := gateway.Get(c, id)
		if err != nil {
			return false, err
		}

		if current.Status == "ACTIVE" {
			return true, nil
		}
		return false, nil
	})
}

func WaitForGatewayDeleted(c *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := gateway.Get(c, id)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, fmt.Errorf("error retriving gateway status: %w", err)
		}
		return false, nil
	})
}
