package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/virtual-interface"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVirtualInterfaceLifecycle(t *testing.T) {
	client, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)

	// Create a virtual interface
	createOpts := virtual_interface.CreateOpts{
		Name:              "test-virtual-interface",
		DirectConnectID:   "test-direct-connect-uuid",
		VgwID:             "test-vgw-id",
		Type:              "private",
		ServiceType:       "vpc",
		VLAN:              100,
		Bandwidth:         100,
		LocalGatewayV4IP:  "192.168.0.1",
		RemoteGatewayV4IP: "192.168.1.1",
		RouteMode:         "static",
		RemoteEPGroupID:   "test-remote-ep-group-id",
	}

	created, err := virtual_interface.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = virtual_interface.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
