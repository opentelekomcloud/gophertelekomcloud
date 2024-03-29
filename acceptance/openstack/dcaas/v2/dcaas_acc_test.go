package v2

import (
	"os"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	dc_endpoint_group "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/dc-endpoint-group"
	direct_connect "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/direct-connect"
	virtual_gateway "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/virtual-gateway"
	virtual_interface "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/virtual-interface"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDCAASLifecycle(t *testing.T) {
	if os.Getenv("RUN_DCAAS_VIRTUAL_INTERFACE") == "" {
		t.Skip("unstable test")
	}

	// Create client
	client, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)

	// Create a direct connect
	name := strings.ToLower(tools.RandomString("test-direct-connect-", 5))
	createOpts := direct_connect.CreateOpts{
		Name:      name,
		PortType:  "1G",
		Bandwidth: 100,
		Location:  "Biere",
		Provider:  "OTC",
	}

	dc, err := direct_connect.Create(client, createOpts)
	th.AssertNoErr(t, err)

	// Get a direct connect
	_, err = direct_connect.Get(client, dc.ID)
	th.AssertNoErr(t, err)

	// List direct connects
	_, err = direct_connect.List(client, dc.ID)
	th.AssertNoErr(t, err)

	// Update a direct connect
	updateOpts := direct_connect.UpdateOpts{
		Name:        tools.RandomString(name, 3),
		Description: "Updated description",
	}
	_ = direct_connect.Update(client, dc.ID, updateOpts)
	th.AssertNoErr(t, err)

	// Create a direct connect endpoint group
	DCegName := strings.ToLower(tools.RandomString("test-direct-connect-endpoint-group-", 5))

	TenantId := clients.EnvOS.GetEnv("TENANT_ID")

	DCegCreateOpts := dc_endpoint_group.CreateOpts{
		TenantId:  TenantId,
		Name:      DCegName,
		Endpoints: []string{"10.2.0.0/24", "10.3.0.0/24"},
		Type:      "cidr",
	}

	dceg, err := dc_endpoint_group.Create(client, DCegCreateOpts)
	th.AssertNoErr(t, err)

	// Create a virtual gateway
	VgName := strings.ToLower(tools.RandomString("test-virtual-gateway-", 5))
	VgCreateOpts := virtual_gateway.CreateOpts{
		Name:                 VgName,
		VpcId:                clients.EnvOS.GetEnv("VPC_ID"),
		LocalEndpointGroupId: dceg.ID,
		Type:                 "default",
	}

	vg, err := virtual_gateway.Create(client, VgCreateOpts)
	th.AssertNoErr(t, err)

	// Create a virtual interface
	ViName := strings.ToLower(tools.RandomString("test-virtual-interface-", 5))
	ViCreateOpts := virtual_interface.CreateOpts{
		Name:              ViName,
		DirectConnectID:   dc.ID,
		VgwID:             vg.ID,
		Type:              "private",
		ServiceType:       "vpc",
		VLAN:              2511,
		Bandwidth:         100,
		LocalGatewayV4IP:  "16.16.16.1/30",
		RemoteGatewayV4IP: "16.16.16.2/30",
		RouteMode:         "static",
		RemoteEPGroupID:   dceg.ID,
		AdminStateUp:      true,
	}

	vi, err := virtual_interface.Create(client, ViCreateOpts)
	if err != nil {
		t.Fatalf("Unable to create virtual interface: %v", err)
	}
	th.AssertNoErr(t, err)

	// Cleanup
	t.Cleanup(func() {
		err = virtual_interface.Delete(client, vi.ID)
		err = virtual_gateway.Delete(client, vg.ID)
		err = dc_endpoint_group.Delete(client, dceg.ID)
		err = direct_connect.Delete(client, dc.ID)
		th.AssertNoErr(t, err)
	})
}
