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
	t.Skip("This API only available in eu-ch2 region for now")
	hostingId := os.Getenv("DCAAS_HOSTING_ID")
	if hostingId == "" {
		// hostingId = "45d7cbf9-b78e-4273-9a16-68f772b6c71d"
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

	t.Logf("Attempting to retrieve list of DCaaSv3 hosted connects")
	l, err := hosted_connect.List(client, hosted_connect.ListOpts{ID: []string{created.ID}})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, l[0].Name)

	t.Logf("Attempting to retrieve DCaaSv3 hosted connect: %s", created.ID)
	c, err := hosted_connect.Get(client, created.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, c.Name)

	t.Logf("Attempting to update DCaaSv3 hosted connect: %s", created.ID)
	u, err := hosted_connect.Update(client, created.ID, hosted_connect.UpdateOpts{
		Description: "update",
		Bandwidth:   20,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "update", u.Description)
	th.AssertEquals(t, 20, u.Bandwidth)

	t.Cleanup(func() {
		t.Logf("Attempting to delete DCaaSv3 hosted connect: %s", created.ID)

		err = hosted_connect.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
