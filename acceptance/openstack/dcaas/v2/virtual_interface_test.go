package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/virtual-interface"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVirtualInterfaceLifecycle(t *testing.T) {
	client, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)

	// Create a virtual interface
	name := tools.RandomString("test-virtual-interface", 5)
	createOpts := virtual_interface.CreateOpts{
		Name:              name,
		DirectConnectID:   "b07d42dc-6137-4af3-a93b-853d879ae268",
		VgwID:             "d27d5bd2-97b3-4bd8-b7e5-189a71c14846",
		Type:              "private",
		ServiceType:       "vpc",
		VLAN:              100,
		Bandwidth:         100,
		LocalGatewayV4IP:  "16.16.16.1/30",
		RemoteGatewayV4IP: "16.16.16.2/30",
		RouteMode:         "static",
		RemoteEPGroupID:   "31dd8536-1ac7-4a38-b2fc-178a69f11b11",
	}

	created, err := virtual_interface.Create(client, createOpts)
	th.AssertNoErr(t, err)

	_, err = virtual_interface.Get(client, created.ID)
	th.AssertNoErr(t, err)

	_ = virtual_interface.Update(client, created.ID, virtual_interface.UpdateOpts{
		Name:        tools.RandomString(name, 3),
		Description: "New description",
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = virtual_interface.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
