package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	direct_connect "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/direct-connect"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDirectConnectLifecycle(t *testing.T) {
	client, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)

	// Create a direct connect
	createOpts := direct_connect.CreateOpts{
		Name:      "test-direct-connect",
		PortType:  "1G",
		Bandwidth: 1000,
		Location:  "Biere",
		Provider:  "OTC",
	}

	created, err := direct_connect.Create(client, createOpts)
	th.AssertNoErr(t, err)

	_, err = direct_connect.Get(client, created.ID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = direct_connect.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
