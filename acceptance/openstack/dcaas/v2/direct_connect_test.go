package v2

import (
	"fmt"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	direct_connect "github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v2/direct-connect"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDirectConnectLifecycle(t *testing.T) {
	client, err := clients.NewDCaaSV2Client()
	th.AssertNoErr(t, err)

	// Create a direct connect
	name := strings.ToLower(tools.RandomString("test-direct-connect", 5))
	createOpts := direct_connect.CreateOpts{
		Name:      name,
		PortType:  "1G",
		Bandwidth: 100,
		Location:  "Biere",
		Provider:  "OTC",
	}

	created, err := direct_connect.Create(client, createOpts)
	th.AssertNoErr(t, err)

	// Get a direct connect
	get, err := direct_connect.Get(client, created.ID)
	fmt.Println(get)
	th.AssertNoErr(t, err)

	// List direct connects
	listed, err := direct_connect.List(client, created.ID)
	fmt.Println(listed)

	th.AssertNoErr(t, err)

	// Update a direct connect
	updateOpts := direct_connect.UpdateOpts{
		Name:        tools.RandomString(name, 3),
		Description: "Updated description",
		Bandwidth:   200,
	}

	updated := direct_connect.Update(client, created.ID, updateOpts)
	fmt.Println(updated)

	// Cleanup
	t.Cleanup(func() {
		err = direct_connect.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
