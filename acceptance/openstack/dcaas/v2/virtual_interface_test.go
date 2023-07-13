package v2

import (
	"fmt"
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

	got, err := virtual_interface.Get(client, created.ID)
	th.AssertNoErr(t, err)

	fmt.Print(got)

	t.Cleanup(func() {
		err = virtual_interface.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
