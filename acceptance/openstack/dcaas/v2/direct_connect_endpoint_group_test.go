package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	dc_endpoint_group "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/dc-endpoint-group"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDirectConnectEndpointGroupLifecycle(t *testing.T) {
	// Create a direct connect endpoint group
	client, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)

	createOpts := dc_endpoint_group.CreateOpts{
		Name:      "test-direct-connect-endpoint-group",
		Endpoints: []string{"10.2.0.0/24", "10.3.0.0/24"},
		Type:      "cidr",
	}

	created, err := dc_endpoint_group.Create(client, createOpts)
	th.AssertNoErr(t, err)

	// Get a direct connect endpoint group
	_, err = dc_endpoint_group.Get(client, created.ID)
	th.AssertNoErr(t, err)

	// List direct connect endpoint groups
	listOpts := dc_endpoint_group.ListOpts{
		ID: created.ID,
	}
	_, err = dc_endpoint_group.List(client, listOpts)
	th.AssertNoErr(t, err)

	// Cleanup
	t.Cleanup(func() {
		err = dc_endpoint_group.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
