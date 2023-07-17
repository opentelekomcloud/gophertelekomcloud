package v2

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/virtual-gateway"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVirtualGatewayLifecycle(t *testing.T) {
	client, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)

	// Create a virtual gateway
	name := strings.ToLower(tools.RandomString("test-virtual-gateway", 5))
	createOpts := virtual_gateway.CreateOpts{
		Name:                 name,
		Type:                 "default",
		VpcId:                clients.EnvOS.GetEnv("VPC_ID"),
		Description:          "test-virtual-interface",
		LocalEndpointGroupId: clients.EnvOS.GetEnv("LOCAL_ENDPOINT_GROUP_ID"),
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

	_, err = virtual_gateway.List(client, created.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = virtual_gateway.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
