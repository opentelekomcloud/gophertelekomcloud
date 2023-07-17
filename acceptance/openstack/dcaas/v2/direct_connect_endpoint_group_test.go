package v2

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	dc_endpoint_group "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/dc-endpoint-group"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDirectConnectEndpointGroupLifecycle(t *testing.T) {
	// Create a direct connect endpoint group
	client, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)

	name := strings.ToLower(tools.RandomString("test-direct-connect-endpoint-group", 5))

	createOpts := dc_endpoint_group.CreateOpts{
		TenantId:  clients.EnvOS.GetEnv("OS_TENANT_ID"),
		Name:      name,
		Endpoints: []string{"10.2.0.0/24", "10.3.0.0/24"},
		Type:      "cidr",
	}

	created, err := dc_endpoint_group.Create(client, createOpts)
	th.AssertNoErr(t, err)

	// Get a direct connect endpoint group
	_, err = dc_endpoint_group.Get(client, created.ID)
	th.AssertNoErr(t, err)

	// Update a direct connect endpoint group
	updateOpts := dc_endpoint_group.UpdateOpts{
		Name:        tools.RandomString(name, 3),
		Description: "test-direct-connect-endpoint-group-updated",
	}
	_ = dc_endpoint_group.Update(client, created.ID, updateOpts)

	// List direct connect endpoint groups
	_, err = dc_endpoint_group.List(client, created.ID)
	th.AssertNoErr(t, err)

	// Cleanup
	t.Cleanup(func() {
		err = dc_endpoint_group.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
