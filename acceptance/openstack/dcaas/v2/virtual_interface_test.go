package v2

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/virtual-interface"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVirtualInterfaceLifecycle(t *testing.T) {
	if os.Getenv("RUN_DCAAS_VIRTUAL_INTERFACE") == "" {
		t.Skip("unstable test")
	}
	client, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)

	// Create a virtual interface
	name := tools.RandomString("test-virtual-interface", 5)
	createOpts := virtual_interface.CreateOpts{
		Name:              name,
		DirectConnectID:   clients.EnvOS.GetEnv("DIRECT_CONNECT_ID"),
		VgwID:             clients.EnvOS.GetEnv("VIRTUAL_GATEWAY_ID"),
		Type:              "private",
		ServiceType:       "vpc",
		VLAN:              100,
		Bandwidth:         100,
		LocalGatewayV4IP:  "16.16.16.1/30",
		RemoteGatewayV4IP: "16.16.16.2/30",
		RouteMode:         "static",
		RemoteEPGroupID:   clients.EnvOS.GetEnv("REMOTE_ENDPOINT_GROUP_ID"),
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

	// List virtual interfaces
	_, err = virtual_interface.List(client, created.ID)
	th.AssertNoErr(t, err)

	// Cleanup
	t.Cleanup(func() {
		err = virtual_interface.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
