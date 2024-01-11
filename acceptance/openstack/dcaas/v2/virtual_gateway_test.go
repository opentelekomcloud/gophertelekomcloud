package v2

import (
	"os"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	dc_endpoint_group "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/dc-endpoint-group"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/virtual-gateway"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVirtualGatewayListing(t *testing.T) {
	client, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)

	opts := virtual_gateway.ListOpts{}
	vgwList, err := virtual_gateway.List(client, opts)
	th.AssertNoErr(t, err)

	for _, vgw := range vgwList {
		tools.PrintResource(t, vgw)
	}
}

func TestVirtualGatewayLifecycle(t *testing.T) {
	if os.Getenv("OS_VPC_ID") == "" {
		t.Skip("OS_VPC_ID necessary for this test")
	}
	client, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)

	egName := strings.ToLower(tools.RandomString("test-acc-dc-eg-", 5))
	egCreateOpts := dc_endpoint_group.CreateOpts{
		TenantId:  client.ProjectID,
		Name:      egName,
		Endpoints: []string{"10.2.0.0/24", "10.3.0.0/24"},
		Type:      "cidr",
	}

	eg, err := dc_endpoint_group.Create(client, egCreateOpts)
	th.AssertNoErr(t, err)

	// Cleanup
	t.Cleanup(func() {
		err = dc_endpoint_group.Delete(client, eg.ID)
		th.AssertNoErr(t, err)
	})

	// Create a virtual gateway
	name := strings.ToLower(tools.RandomString("test-virtual-gateway", 5))
	createOpts := virtual_gateway.CreateOpts{
		Name:                 name,
		Type:                 "default",
		VpcId:                os.Getenv("OS_VPC_ID"),
		Description:          "test-virtual-interface",
		LocalEndpointGroupId: eg.ID,
	}

	created, err := virtual_gateway.Create(client, createOpts)
	th.AssertNoErr(t, err)

	_, err = virtual_gateway.Get(client, created.ID)
	th.AssertNoErr(t, err)

	updateOpts := virtual_gateway.UpdateOpts{
		Name:        tools.RandomString(name, 3),
		Description: "test-virtual-interface-updated",
	}
	_ = virtual_gateway.Update(client, created.ID, updateOpts)
	th.AssertNoErr(t, err)

	opts := virtual_gateway.ListOpts{}
	_, err = virtual_gateway.List(client, opts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = virtual_gateway.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
