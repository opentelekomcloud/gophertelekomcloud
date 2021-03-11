package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCloudServerLifecycle(t *testing.T) {
	client, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	// Get ECSv1 createOpts
	createOpts := getCloudServerCreateOpts(t)
	createOpts.DryRun = true

	// Check ECSv1 createOpts
	err = dryRunCloudServerConfig(t, client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("CreateOpts are true for creating a cloudServer")
	createOpts.DryRun = false

	// Create ECSv1 instance
	ecs := createCloudServer(t, client, createOpts)
	defer deleteCloudServer(t, client, ecs.ID)

	tools.PrintResource(t, ecs)
}
