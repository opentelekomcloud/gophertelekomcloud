package ports

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/networks"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/ports"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/subnets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPortLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	network := createNetwork(t, client)
	t.Cleanup(func() {
		deleteNetwork(t, client, network.ID)
	})

	subnet := createSubnet(t, client, network.ID)
	t.Cleanup(func() {
		deleteSubnet(t, client, subnet.ID)
	})

	adminStateUp := true
	createOpts := ports.CreateOpts{
		NetworkID:    network.ID,
		Name:         tools.RandomString("acc-port-create-", 3),
		AdminStateUp: &adminStateUp,
		FixedIps: []ports.FixedIp{
			{SubnetId: subnet.ID},
		},
	}

	t.Logf("Attempting to create port: %s", createOpts.Name)
	port, err := ports.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		deletePort(t, client, port.ID)
	})

	th.AssertEquals(t, createOpts.Name, port.Name)
	th.AssertEquals(t, createOpts.NetworkID, port.NetworkID)
	th.AssertEquals(t, *createOpts.AdminStateUp, port.AdminStateUp)
	th.AssertEquals(t, 1, len(port.FixedIPs))
	th.AssertEquals(t, subnet.ID, port.FixedIPs[0].SubnetId)
	t.Logf("Created port: %s", port.ID)

	updatedName := tools.RandomString("acc-port-update-", 3)
	adminStateDown := false
	updateOpts := ports.UpdateOpts{
		Name:         updatedName,
		AdminStateUp: &adminStateDown,
	}

	t.Logf("Attempting to update port: %s", port.ID)
	updatedPort, err := ports.Update(client, port.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, updatedPort.Name)
	th.AssertEquals(t, *updateOpts.AdminStateUp, updatedPort.AdminStateUp)

	gotPort, err := ports.Get(client, port.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, gotPort.Name)
	th.AssertEquals(t, *updateOpts.AdminStateUp, gotPort.AdminStateUp)
	th.AssertEquals(t, network.ID, gotPort.NetworkID)

	portPages, err := ports.List(client, ports.ListOpts{
		ID:        port.ID,
		NetworkID: network.ID,
	}).AllPages()
	th.AssertNoErr(t, err)

	portList, err := ports.ExtractPorts(portPages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(portList))
	th.AssertEquals(t, port.ID, portList[0].ID)
	th.AssertEquals(t, updateOpts.Name, portList[0].Name)
}

func createNetwork(t *testing.T, client *golangsdk.ServiceClient) *networks.Network {
	adminStateUp := true
	createOpts := networks.CreateOpts{
		Name:         tools.RandomString("acc-port-net-", 3),
		AdminStateUp: &adminStateUp,
	}

	t.Logf("Attempting to create network: %s", createOpts.Name)
	network, err := networks.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created network: %s", network.ID)

	return network
}

func deleteNetwork(t *testing.T, client *golangsdk.ServiceClient, networkID string) {
	t.Logf("Attempting to delete network: %s", networkID)
	err := networks.Delete(client, networkID).ExtractErr()
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, waitForNetworkDeleted(client, networkID, 120))
	t.Logf("Deleted network: %s", networkID)
}

func createSubnet(t *testing.T, client *golangsdk.ServiceClient, networkID string) *subnets.Subnet {
	gatewayIP := "192.168.42.1"
	enableDHCP := true
	createOpts := subnets.CreateOpts{
		NetworkID:      networkID,
		Name:           tools.RandomString("acc-port-subnet-", 3),
		CIDR:           "192.168.42.0/24",
		IPVersion:      golangsdk.IPv4,
		GatewayIP:      &gatewayIP,
		EnableDHCP:     &enableDHCP,
		DNSNameservers: []string{"1.1.1.1", "8.8.8.8"},
	}

	t.Logf("Attempting to create subnet: %s", createOpts.Name)
	subnet, err := subnets.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created subnet: %s", subnet.ID)

	return subnet
}

func deleteSubnet(t *testing.T, client *golangsdk.ServiceClient, subnetID string) {
	t.Logf("Attempting to delete subnet: %s", subnetID)
	err := subnets.Delete(client, subnetID).ExtractErr()
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, waitForSubnetDeleted(client, subnetID, 120))
	t.Logf("Deleted subnet: %s", subnetID)
}

func deletePort(t *testing.T, client *golangsdk.ServiceClient, portID string) {
	t.Logf("Attempting to delete port: %s", portID)
	err := ports.Delete(client, portID).ExtractErr()
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, waitForPortDeleted(client, portID, 120))
	t.Logf("Deleted port: %s", portID)
}

func waitForPortDeleted(client *golangsdk.ServiceClient, portID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := ports.Get(client, portID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}

func waitForSubnetDeleted(client *golangsdk.ServiceClient, subnetID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := subnets.Get(client, subnetID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}

func waitForNetworkDeleted(client *golangsdk.ServiceClient, networkID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := networks.Get(client, networkID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}
