package v3

import (
	"os"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	direct_connect "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/direct-connect"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v3/virtual-gateway"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v3/virtual-interface"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVirtualInterfaceListing(t *testing.T) {
	client, err := clients.NewDCaaSV3Client()
	th.AssertNoErr(t, err)

	// List virtual interfaces
	opts := virtual_interface.ListOpts{}
	viList, err := virtual_interface.List(client, opts)
	th.AssertNoErr(t, err)

	for _, vi := range viList {
		tools.PrintResource(t, vi)
	}
}

func TestVirtualInterfaceLifecycle(t *testing.T) {
	if os.Getenv("RUN_DCAAS_VIRTUAL_INTERFACE") == "" {
		t.Skip("DIRECT_CONNECT_ID necessary for this test or run it only in test_terraform")
	}
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" {
		t.Skip("OS_VPC_ID necessary for this test")
	}
	clientV2, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)
	clientV3, err := clients.NewDCaaSV3Client()
	th.AssertNoErr(t, err)

	name := strings.ToLower(tools.RandomString("acc-direct-connect", 5))
	createOpts := direct_connect.CreateOpts{
		Name:      name,
		PortType:  "1G",
		Bandwidth: 100,
		Location:  "Biere",
		Provider:  "OTC",
	}

	dc, err := direct_connect.Create(clientV2, createOpts)
	th.AssertNoErr(t, err)

	clientNet, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)
	vpc, err := vpcs.Get(clientNet, vpcID).Extract()
	th.AssertNoErr(t, err)

	// Create a virtual gateway
	nameVgw := strings.ToLower(tools.RandomString("acc-virtual-gateway-v3-", 5))
	vgOpts := virtual_gateway.CreateOpts{
		Name:         nameVgw,
		VpcId:        vpcID,
		Description:  "acc-virtual-gateway-v3",
		LocalEpGroup: []string{vpc.CIDR},
		Tags: []tags.ResourceTag{
			{
				Key:   "TestKey",
				Value: "TestValue",
			},
		},
	}

	vg, err := virtual_gateway.Create(clientV3, vgOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = virtual_gateway.Delete(clientV3, vg.ID)
		th.AssertNoErr(t, err)
	})

	// Create a virtual interface
	nameVi := tools.RandomString("acc-virtual-interface-", 5)
	viOpts := virtual_interface.CreateOpts{
		Name:              nameVi,
		DirectConnectID:   dc.ID,
		VgwId:             vg.ID,
		Type:              "private",
		ServiceType:       "vpc",
		VLAN:              100,
		Bandwidth:         5,
		LocalGatewayV4IP:  "16.16.16.1/30",
		RemoteGatewayV4IP: "16.16.16.2/30",
		RouteMode:         "static",
		RemoteEpGroup:     []string{vpc.CIDR},
	}

	created, err := virtual_interface.Create(clientV3, viOpts)
	th.AssertNoErr(t, err)

	vi, err := virtual_interface.Get(clientV3, created.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vi.RemoteEpGroup[0], vpc.CIDR)

	updated, err := virtual_interface.Update(clientV3, created.ID, virtual_interface.UpdateOpts{
		Name:        name + "-updated",
		Description: pointerto.String("New description"),
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name+"-updated", updated.Name)

	// Cleanup
	t.Cleanup(func() {
		err = virtual_interface.Delete(clientV3, created.ID)
		th.AssertNoErr(t, err)
	})
}
