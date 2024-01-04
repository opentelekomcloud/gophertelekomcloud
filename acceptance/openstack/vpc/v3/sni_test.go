package v3

import (
	"fmt"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/subnets"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/networks"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/ports"
	subnetsV2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/subnets"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v3/sni"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSNIV3Listing(t *testing.T) {
	client, err := clients.NewVPCV3Client()
	th.AssertNoErr(t, err)

	listOpts := sni.ListOpts{}
	snisList, err := sni.List(client, listOpts)
	th.AssertNoErr(t, err)

	for _, rtb := range snisList {
		tools.PrintResource(t, rtb)
	}

	sniCount, err := sni.GetCount(client)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, sniCount.SubNetworkInterfaces, 0)
}

func TestPortListing(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	listOpts := ports.ListOpts{}
	portsPages, err := ports.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	var allPorts []ports.Port

	err = ports.ExtractPortsInto(portsPages, &allPorts)
	th.AssertNoErr(t, err)

	for _, pts := range allPorts {
		tools.PrintResource(t, pts)
	}
	portId := "9ce9a5c2-7af0-4cf7-9adf-469e8f918721"
	ports.Delete(client, portId)
}

func TestSNIV3Lifecycle(t *testing.T) {
	clientV1, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	clientV2, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	clientV3, err := clients.NewVPCV3Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("vpc-sni-acc-", 3)
	createOpts := vpcs.CreateOpts{
		Name:        name,
		Description: "some interesting description",
		CIDR:        "192.168.0.0/16",
	}

	vpc, err := vpcs.Create(clientV1, createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vpc.EnableSharedSnat, false)
	th.AssertEquals(t, vpc.Description, "some interesting description")
	th.AssertEquals(t, vpc.Name, name)

	t.Cleanup(func() {
		err = vpcs.Delete(clientV1, vpc.ID).ExtractErr()
		th.AssertNoErr(t, err)
	})

	networkName := tools.RandomString("network-sni-acc-", 3)
	createNetworkOpts := NetworkCreateOpts{
		CreateOpts: networks.CreateOpts{
			Name: networkName,
		},
	}

	t.Logf("Attempting to create network: %s", networkName)
	n, err := networks.Create(clientV2, createNetworkOpts).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Logf("Attempting to delete network: %s", networkName)
		networks.Delete(clientV2, n.ID)
	})

	subnetName := tools.RandomString("network-sni-acc-", 3)
	createSubnetOpts := SubnetCreateOpts{
		CreateOpts: subnetsV2.CreateOpts{
			NetworkID:  n.ID,
			CIDR:       "192.168.20.0/24",
			Name:       subnetName,
			EnableDHCP: nil,
		},
	}
	s, err := subnetsV2.Create(clientV2, createSubnetOpts).Extract()
	th.AssertNoErr(t, err)

	routerName := tools.RandomString("router-sni-acc-", 3)
	createRouterOpts := RouterCreateOpts{
		CreateOpts: routers.CreateOpts{
			Name: routerName,
		},
	}
	t.Logf("Attempting to create router: %s", routerName)
	router, err := routers.Create(clientV2, createRouterOpts).Extract()
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete router: %s", routerName)
		err = routers.Delete(clientV2, router.ID).ExtractErr()
		th.AssertNoErr(t, err)
	})
	createIntOpts := routers.AddInterfaceOpts{
		SubnetID: s.ID,
	}

	t.Logf("Attempting to create network interface")
	ri, err := routers.AddInterface(clientV2, router.ID, createIntOpts).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Logf("Attempting to delete network interface: %s", routerName)
		removeIntOpts := routers.RemoveInterfaceOpts{
			SubnetID: s.ID,
		}
		_, err = routers.RemoveInterface(clientV2, router.ID, removeIntOpts).Extract()
		th.AssertNoErr(t, err)
	})
	// subnet := createSubnet(t, clientV1, n.ID)
	// t.Cleanup(func() {
	// 	deleteSubnet(t, clientV1, subnet.VpcID, subnet.ID)
	// })

	// createPortOpts := ports.CreateOpts{
	// 	Name:      "acc-sni-private-port",
	// 	NetworkID: subnet.NetworkID,
	// }
	// t.Logf("Attempting to create port: %s", createPortOpts.Name)
	// port, err := ports.Create(clientV2, createPortOpts).Extract()
	// th.AssertNoErr(t, err)
	//
	// t.Cleanup(func() {
	// 	t.Logf("Attempting to delete port: %s", port.Name)
	// 	err = ports.Delete(clientV2, port.ID).ExtractErr()
	// 	th.AssertNoErr(t, err)
	// })

	sniOpts := sni.CreateOpts{
		Sni: &sni.CreateSubNetworkInterfaceOption{
			PortId:     ri.ID,
			SubnetId:   s.ID,
			Ipv6Enable: pointerto.Bool(true),
		},
	}

	subInterface, err := sni.Create(clientV3, sniOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, subInterface.VpcId, n.ID)

	t.Cleanup(func() {
		err := sni.Delete(clientV3, subInterface.Id)
		th.AssertNoErr(t, err)
	})

	// cidrOpts := VpcV3.CidrOpts{
	// 	Vpc: &VpcV3.AddExtendCidrOption{
	// 		ExtendCidrs: []string{"23.8.0.0/16"}},
	// }
	// vpcSecCidr, err := VpcV3.AddSecondaryCidr(clientV3, vpc.ID, cidrOpts)
	// th.AssertNoErr(t, err)
	// th.AssertEquals(t, vpcSecCidr.SecondaryCidrs[0], "23.8.0.0/16")
	//
	// vpcV3Get, err := VpcV3.Get(clientV3, vpc.ID)
	// th.AssertNoErr(t, err)
	// th.AssertEquals(t, vpcV3Get.SecondaryCidrs[0], "23.8.0.0/16")
	//
	// vpcSecCidrRm, err := VpcV3.RemoveSecondaryCidr(clientV3, vpc.ID, cidrOpts)
	// th.AssertNoErr(t, err)
	// th.AssertEquals(t, len(vpcSecCidrRm.SecondaryCidrs), 0)
}

func createSubnet(t *testing.T, client *golangsdk.ServiceClient, vpcID string) *subnets.Subnet {
	createSubnetOpts := subnets.CreateOpts{
		Name:        tools.RandomString("acc-sni-subnet-", 3),
		Description: "some description",
		CIDR:        "192.168.20.0/24",
		GatewayIP:   "192.168.20.1",
		EnableDHCP:  pointerto.Bool(true),
		VpcID:       vpcID,
	}
	t.Logf("Attempting to create subnet: %s", createSubnetOpts.Name)

	subnet, err := subnets.Create(client, createSubnetOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, subnet.Description, createSubnetOpts.Description)

	// wait to be active
	t.Logf("Waitting for subnet %s to be active", subnet.ID)
	err = waitForSubnetToBeActive(client, subnet.ID, 600)
	th.AssertNoErr(t, err)
	t.Logf("Created subnet: %v", subnet.ID)

	return subnet
}

func waitForSubnetToBeActive(client *golangsdk.ServiceClient, subnetID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		n, err := subnets.Get(client, subnetID).Extract()
		if err != nil {
			return false, err
		}

		if n.Status == "ACTIVE" {
			return true, nil
		}

		// If subnet status is other than Active, send error
		if n.Status == "DOWN" || n.Status == "ERROR" {
			return false, fmt.Errorf("subnet status: '%s'", n.Status)
		}

		return false, nil
	})
}

func deleteSubnet(t *testing.T, client *golangsdk.ServiceClient, vpcID string, id string) {
	t.Logf("Attempting to delete subnet: %s", id)

	err := subnets.Delete(client, vpcID, id).ExtractErr()
	th.AssertNoErr(t, err)

	t.Logf("Waiting for subnet %s to be deleted", id)
	err = waitForSubnetToBeDeleted(client, id, 60)
	th.AssertNoErr(t, err)

	t.Logf("Deleted subnet: %s", id)
}

func waitForSubnetToBeDeleted(client *golangsdk.ServiceClient, subnetID string, secs int) error {
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

// NetworkCreateOpts represents the attributes used when creating a new network.
type NetworkCreateOpts struct {
	networks.CreateOpts
	ValueSpecs map[string]string `json:"value_specs,omitempty"`
}

// SubnetCreateOpts represents the attributes used when creating a new subnet.
type SubnetCreateOpts struct {
	subnetsV2.CreateOpts
	ValueSpecs map[string]string `json:"value_specs,omitempty"`
}

// RouterCreateOpts represents the attributes used when creating a new router.
type RouterCreateOpts struct {
	routers.CreateOpts
	ValueSpecs map[string]string `json:"value_specs,omitempty"`
}
