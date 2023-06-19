package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestServerList(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	listOpts := servers.ListOpts{}
	allServerPages, err := servers.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	serversList, err := servers.ExtractServers(allServerPages)
	th.AssertNoErr(t, err)

	for _, server := range serversList {
		tools.PrintResource(t, server)
	}
}

func TestServerLifecycle(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	ecs, err := CreateServer(t, client)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		DeleteServer(t, client, ecs)
	})

	nicInfo, err := servers.GetNICs(client, ecs.ID).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, nicInfo)

	t.Logf("Attempting to update ECSv2: %s", ecs.ID)

	ecsName := tools.RandomString("update-ecs-", 3)
	_, err = servers.Update(client, ecs.ID, servers.UpdateOpts{
		Name: ecsName,
	}).Extract()
	th.AssertNoErr(t, err)

	t.Logf("ECSv2 successfully updated: %s", ecs.ID)
	th.AssertNoErr(t, err)

	newECS, err := servers.Get(client, ecs.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ecsName, newECS.Name)
}
