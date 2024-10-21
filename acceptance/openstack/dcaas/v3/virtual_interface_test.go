package v3

import (
	"os"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v3/virtual-gateway"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v3/virtual-interface"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVirtualInterfaceListing(t *testing.T) {
	t.Skip("This API only available in eu-ch2 region for now")
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
	t.Skip("This API only available in eu-ch2 region for now")
	dcID := os.Getenv("DIRECT_CONNECT_ID")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && dcID == "" {
		t.Skip("DIRECT_CONNECT_ID and OS_VPC_ID necessary for this test")
	}
	client, err := clients.NewDCaaSV3Client()
	th.AssertNoErr(t, err)

	clientNet, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)
	vpc, err := vpcs.Get(clientNet, vpcID).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create DCaaSv3 virtual gateway")
	name := strings.ToLower(tools.RandomString("acc-virtual-gateway-v3-", 5))
	createOpts := virtual_gateway.CreateOpts{
		Name:         name,
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
	vg, err := virtual_gateway.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = virtual_gateway.Delete(client, vg.ID)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to create DCaaSv3 virtual interface")
	nameVi := tools.RandomString("acc-virtual-interface-", 5)
	viOpts := virtual_interface.CreateOpts{
		Name:              nameVi,
		DirectConnectID:   dcID,
		VgwId:             vg.ID,
		Type:              "private",
		ServiceType:       "vpc",
		VLAN:              100,
		Bandwidth:         5,
		LocalGatewayV4IP:  "16.16.16.1/30",
		RemoteGatewayV4IP: "16.16.16.2/30",
		RouteMode:         "static",
		RemoteEpGroup:     []string{"16.16.16.0/30"},
	}
	created, err := virtual_interface.Create(client, viOpts)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to obtain DCaaSv3 virtual interface: %s", created.ID)
	vi, err := virtual_interface.Get(client, created.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "16.16.16.0/30", vi.RemoteEpGroup[0])

	t.Logf("Attempting to update DCaaSv3 virtual interface: %s", created.ID)
	updated, err := virtual_interface.Update(client, created.ID, virtual_interface.UpdateOpts{
		Name:        name + "-updated",
		Description: pointerto.String("New description"),
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name+"-updated", updated.Name)
	th.AssertEquals(t, "New description", updated.Description)

	t.Logf("Attempting to obtain list of DCaaSv3 virtual interfaces")
	viList, err := virtual_interface.List(client, virtual_interface.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(viList))

	t.Cleanup(func() {
		t.Logf("Attempting to delete DCaaSv3 virtual interface: %s", created.ID)
		err = virtual_interface.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
