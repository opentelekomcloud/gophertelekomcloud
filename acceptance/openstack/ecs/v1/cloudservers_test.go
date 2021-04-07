package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCloudServerLifecycle(t *testing.T) {
	client, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	// Get ECSv1 createOpts
	createOpts := openstack.GetCloudServerCreateOpts(t)

	// Check ECSv1 createOpts
	openstack.DryRunCloudServerConfig(t, client, createOpts)
	t.Logf("CreateOpts are ok for creating a cloudServer")

	// Create ECSv1 instance
	ecs := openstack.CreateCloudServer(t, client, createOpts)
	defer openstack.DeleteCloudServer(t, client, ecs.ID)

	tools.PrintResource(t, ecs)
}
