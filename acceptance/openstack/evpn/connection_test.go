package evpn

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evpn/v5/connection"
	cgw "github.com/opentelekomcloud/gophertelekomcloud/openstack/evpn/v5/customer-gateway"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evpn/v5/gateway"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/eips"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/subnets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestConnectionList(t *testing.T) {
	client, err := clients.NewEVPNClient()
	th.AssertNoErr(t, err)
	conn, err := connection.List(client, connection.ListOpts{})
	th.AssertNoErr(t, err)
	for _, c := range conn {
		tools.PrintResource(t, c)
	}
}

func TestConnectionLifecycle(t *testing.T) {
	subnetId := os.Getenv("OS_SUBNET_ID")
	vpcId := os.Getenv("OS_VPC_ID")

	if subnetId == "" && vpcId == "" {
		t.Skip("`OS_SUBNET_ID` and `OS_VPC_ID` and needs to be defined for test")
	}

	client, err := clients.NewEVPNClient()
	th.AssertNoErr(t, err)

	clientNetV2, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	subnet, err := subnets.Get(clientNetV2, subnetId).Extract()
	th.AssertNoErr(t, err)

	gw := createEvpnGateway(t, subnet, vpcId, client)
	customerGw := createEvpnCustomerGateway(t, client)

	name := tools.RandomString("acc_evpn_connection_", 3)
	cOpts := connection.CreateOpts{
		Name:        name,
		VgwId:       gw.ID,
		VgwIp:       gw.Eip1.ID,
		CgwId:       customerGw.ID,
		PeerSubnets: []string{"192.168.55.0/24"},
		Tags: []tags.ResourceTag{
			{
				Key:   "CGW",
				Value: "Value",
			},
		},
		Psk: "abcde1Z!",
	}
	t.Logf("Attempting to CREATE Enterprise VPN Connection: %s", name)
	conn, err := connection.Create(client, cOpts)
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, WaitForConnectionActive(client, conn.ID, 800))

	t.Cleanup(func() {
		t.Logf("Attempting to DELETE Enterprise VPN Connection: %s", gw.ID)
		th.AssertNoErr(t, connection.Delete(client, conn.ID))
		th.AssertNoErr(t, WaitForConnectionDeleted(client, conn.ID, 800))

	})

	t.Logf("Attempting to GET Enterprise VPN Connection: %s", conn.ID)
	connGet, err := connection.Get(client, conn.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, connGet.Name)
	th.AssertEquals(t, "CGW", connGet.Tags[0].Key)
	th.AssertEquals(t, "Value", connGet.Tags[0].Value)

	t.Logf("Attempting to UPDATE Enterprise VPN Connection: %s", conn.ID)
	connUpdate, err := connection.Update(client, connection.UpdateOpts{
		ConnectionID: conn.ID,
		Name:         name + "updated",
		IkePolicy: &connection.IkePolicy{
			LifetimeSeconds: pointerto.Int(3600),
		},
		IpSecPolicy: &connection.IpSecPolicy{
			LifetimeSeconds: pointerto.Int(3500),
		},
	})
	th.AssertNoErr(t, WaitForConnectionActive(client, connUpdate.ID, 800))
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name+"updated", connUpdate.Name)
	th.AssertEquals(t, 3500, *connUpdate.IpSecPolicy.LifetimeSeconds)
	th.AssertEquals(t, 3600, *connUpdate.IkePolicy.LifetimeSeconds)

	t.Logf("Attempting to LIST Enterprise VPN Connections")
	connList, err := connection.List(client, connection.ListOpts{
		VgwId: gw.ID,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(connList))
}

func createEvpnCustomerGateway(t *testing.T, client *golangsdk.ServiceClient) *cgw.CustomerGateway {
	name := tools.RandomString("acc_evpn_cgw_", 5)
	cOpts := cgw.CreateOpts{
		Name:    name,
		IdType:  "ip",
		IdValue: "10.1.2.10",
		BgpAsn:  pointerto.Int(65000),
		Tags: []tags.ResourceTag{
			{
				Key:   "NewKey",
				Value: "NewValue",
			},
		},
	}
	t.Logf("Attempting to CREATE Enterprise VPN Customer Gateway: %s", name)
	gw, err := cgw.Create(client, cOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, gw.Name)
	t.Cleanup(func() {
		t.Logf("Attempting to DELETE Enterprise VPN Customer Gateway: %s", gw.ID)
		th.AssertNoErr(t, cgw.Delete(client, gw.ID))
	})
	return gw
}
func createEvpnGateway(t *testing.T, subnet *subnets.Subnet, vpcId string, client *golangsdk.ServiceClient) *gateway.Gateway {
	clientNetV1, err := clients.NewNetworkV1Client()
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
	fullGw, err := gateway.Get(client, gw.ID)
	th.AssertNoErr(t, err)

	return fullGw
}

func WaitForConnectionDeleted(c *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := connection.Get(c, id)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, fmt.Errorf("error retriving connection status: %w", err)
		}
		return false, nil
	})
}

func WaitForConnectionActive(c *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := connection.Get(c, id)
		if err != nil {
			return false, err
		}

		if current.Status == "ACTIVE" || current.Status == "DOWN" {
			return true, nil
		}
		return false, nil
	})
}
