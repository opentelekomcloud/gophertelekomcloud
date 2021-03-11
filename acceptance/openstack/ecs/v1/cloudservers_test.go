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

	createOpts := getCloudServerCreateOpts(t)
	createOpts.DryRun = true

	err = dryRunCloudServerConfig(t, client, createOpts)
	th.AssertNoErr(t, err)
	createOpts.DryRun = false

	// Create ECSv1 instance
	ecs := createCloudServer(t, client, createOpts)
	defer deleteCloudServer(t, client, ecs.ID)

	tools.PrintResource(t, ecs)
}
