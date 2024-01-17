package v2

import (
	"os"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	dc_endpoint_group "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/dc-endpoint-group"
	virtual_gateway "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/virtual-gateway"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/virtual-interface"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVirtualInterfaceListing(t *testing.T) {
	client, err := clients.NewDCaaSV2Client()
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
	client, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)
	dcId := os.Getenv("DIRECT_CONNECT_ID")
	if dcId == "" {
		// this direct connect available only in test_terraform project
		dcId = "ea254fe6-e16b-4d9c-88de-543d6bb91a28"
	}

	// Create endpoint group
	localEgName := strings.ToLower(tools.RandomString("test-acc-dc-eg-", 5))
	localEgCreateOpts := dc_endpoint_group.CreateOpts{
		TenantId:  client.ProjectID,
		Name:      localEgName,
		Endpoints: []string{"100.2.0.0/24", "100.3.0.0/24"},
		Type:      "cidr",
	}

	localEg, err := dc_endpoint_group.Create(client, localEgCreateOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = dc_endpoint_group.Delete(client, localEg.ID)
		th.AssertNoErr(t, err)
	})

	remoteEgName := strings.ToLower(tools.RandomString("test-acc-dc-eg-", 5))
	remoteEgCreateOpts := dc_endpoint_group.CreateOpts{
		TenantId:  client.ProjectID,
		Name:      remoteEgName,
		Endpoints: []string{"172.16.0.0/24"},
		Type:      "cidr",
	}

	remoteEg, err := dc_endpoint_group.Create(client, remoteEgCreateOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = dc_endpoint_group.Delete(client, remoteEg.ID)
		th.AssertNoErr(t, err)
	})

	// Create a virtual gateway
	vgName := strings.ToLower(tools.RandomString("test-virtual-gateway", 5))
	vgCreateOpts := virtual_gateway.CreateOpts{
		Name:                 vgName,
		Type:                 "default",
		VpcId:                os.Getenv("OS_VPC_ID"),
		Description:          "test-virtual-interface",
		LocalEndpointGroupId: localEg.ID,
	}

	vg, err := virtual_gateway.Create(client, vgCreateOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = virtual_gateway.Delete(client, vg.ID)
		th.AssertNoErr(t, err)
	})

	// Create a virtual interface
	name := tools.RandomString("test-virtual-interface-", 5)
	createOpts := virtual_interface.CreateOpts{
		Name:              name,
		DirectConnectID:   dcId,
		VgwID:             vg.ID,
		Type:              "private",
		ServiceType:       "vpc",
		VLAN:              100,
		Bandwidth:         5,
		LocalGatewayV4IP:  "16.16.16.1/30",
		RemoteGatewayV4IP: "16.16.16.2/30",
		RouteMode:         "static",
		RemoteEPGroupID:   remoteEg.ID,
	}

	created, err := virtual_interface.Create(client, createOpts)
	th.AssertNoErr(t, err)

	vi, err := virtual_interface.Get(client, created.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vi.RemoteEPGroupID, remoteEg.ID)

	err = virtual_interface.Update(client, created.ID, virtual_interface.UpdateOpts{
		Name:        name + "-updated",
		Description: "New description",
	})
	th.AssertNoErr(t, err)

	// Cleanup
	t.Cleanup(func() {
		err = virtual_interface.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
