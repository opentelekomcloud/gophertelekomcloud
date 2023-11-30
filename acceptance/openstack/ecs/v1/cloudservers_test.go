package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
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

	tagsList := []tags.ResourceTag{
		{
			Key:   "TestKey",
			Value: "TestValue",
		},
		{
			Key:   "empty",
			Value: "",
		},
	}
	err = tags.Create(client, "cloudservers", ecs.ID, tagsList).ExtractErr()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, ecs)
}
