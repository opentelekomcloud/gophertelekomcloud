package v3

import (
	"os"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	hosted_connect "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v3/hosted-connect"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestHostedConnectLifecycle(t *testing.T) {
	t.Skip("The API does not exist or has not been published in the environment")
	hostingId := os.Getenv("DCAAS_HOSTING_ID")
	if hostingId == "" {
		t.Skip("DCAAS_HOSTING_ID must be set for test")
	}

	client, err := clients.NewDCaaSV3Client()
	th.AssertNoErr(t, err)

	// Create a hosted connect
	name := strings.ToLower(tools.RandomString("test-hosted-connect", 5))
	createOpts := hosted_connect.CreateOpts{
		Name:             name,
		Description:      "hosted",
		Bandwidth:        10,
		Vlan:             441,
		ResourceTenantId: client.ProjectID,
		HostingID:        hostingId,
	}
	t.Logf("Attempting to create DCaaSv3 hosted connect")
	created, err := hosted_connect.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delte DCaaSv3 hosted connect: %s", created.ID)

		err = hosted_connect.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
