package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/ports"
	v1groups "github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/securitygroups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/subnets"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/vpcs"
	v2groups "github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/securitygroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSecurityGroupUpdatePorts(t *testing.T) {
	v1Client, err := clients.NewVPCV1Client()
	th.AssertNoErr(t, err)
	v2Client, err := clients.NewVPCV2Client()
	th.AssertNoErr(t, err)
	networkClient, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	group, err := v1groups.Create(v1Client, v1groups.CreateOpts{
		Name: tools.RandomString("security-group-ports-acc-", 3),
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, v1groups.Delete(v1Client, group.ID))
	})

	vpc := createSecurityGroupVPC(t, v1Client)
	t.Cleanup(func() {
		deleteSecurityGroupVPC(t, v1Client, vpc.ID)
	})
	subnet := createSecurityGroupSubnet(t, v1Client, vpc.ID)
	t.Cleanup(func() {
		deleteSecurityGroupSubnet(t, v1Client, subnet.VpcID, subnet.ID)
	})

	port, err := ports.Create(networkClient, ports.CreateOpts{
		NetworkID: subnet.NetworkID,
		Name:      tools.RandomString("security-group-port-acc-", 3),
	}).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, ports.Delete(networkClient, port.ID).ExtractErr())
	})

	failures, err := v2groups.UpdatePorts(v2Client, group.ID, v2groups.UpdatePortsOpts{
		Ports:  []v2groups.Port{{ID: port.ID}},
		Action: "add",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(failures))

	failures, err = v2groups.UpdatePorts(v2Client, group.ID, v2groups.UpdatePortsOpts{
		Ports:  []v2groups.Port{{ID: port.ID}},
		Action: "remove",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(failures))
}

func createSecurityGroupVPC(t *testing.T, client *golangsdk.ServiceClient) *vpcs.Vpc {
	t.Helper()
	vpc, err := vpcs.Create(client, vpcs.CreateOpts{
		Name: tools.RandomString("security-group-vpc-acc-", 3),
		CIDR: "192.168.0.0/16",
	})
	th.AssertNoErr(t, err)
	return vpc
}

func createSecurityGroupSubnet(t *testing.T, client *golangsdk.ServiceClient, vpcID string) *subnets.Subnet {
	t.Helper()
	subnet, err := subnets.Create(client, subnets.CreateOpts{
		Name:       tools.RandomString("security-group-subnet-acc-", 3),
		CIDR:       "192.168.30.0/24",
		GatewayIP:  "192.168.30.1",
		EnableDHCP: pointerto.Bool(true),
		VpcID:      vpcID,
	})
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, golangsdk.WaitFor(600, func() (bool, error) {
		current, err := subnets.Get(client, subnet.ID)
		if err != nil {
			return false, err
		}
		return current.Status == "ACTIVE", nil
	}))
	return subnet
}

func deleteSecurityGroupSubnet(t *testing.T, client *golangsdk.ServiceClient, vpcID, subnetID string) {
	t.Helper()
	th.AssertNoErr(t, subnets.Delete(client, vpcID, subnetID))
	th.AssertNoErr(t, golangsdk.WaitFor(600, func() (bool, error) {
		_, err := subnets.Get(client, subnetID)
		if _, ok := err.(golangsdk.ErrDefault404); ok {
			return true, nil
		}
		return false, err
	}))
}

func deleteSecurityGroupVPC(t *testing.T, client *golangsdk.ServiceClient, vpcID string) {
	t.Helper()
	th.AssertNoErr(t, golangsdk.WaitFor(600, func() (bool, error) {
		err := vpcs.Delete(client, vpcID)
		if err == nil {
			return true, nil
		}
		if _, ok := err.(golangsdk.ErrDefault404); ok {
			return true, nil
		}
		if _, ok := err.(golangsdk.ErrDefault409); ok {
			return false, nil
		}
		return false, err
	}))
}
