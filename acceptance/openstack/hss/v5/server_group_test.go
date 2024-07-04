package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	hss "github.com/opentelekomcloud/gophertelekomcloud/openstack/hss/v5/host"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestServerGroupList(t *testing.T) {
	client, err := clients.NewHssClient()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, client)
	listResp, err := hss.List(client, hss.ListOpts{})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}

func TestServerList(t *testing.T) {
	client, err := clients.NewHssClient()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, client)
	listResp, err := hss.ListHost(client, hss.ListHostOpts{})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}

func TestServerWorkflow(t *testing.T) {
	client, err := clients.NewHssClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Create member for Server group")
	ecsClient, err := clients.NewComputeV2Client()
	ecs := openstack.CreateServer(t, ecsClient,
		tools.RandomString("hss-group-member-", 3),
		"Standard_Debian_10_latest",
		"s2.large.2",
	)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Server group member")
		th.AssertNoErr(t, servers.Delete(ecsClient, ecs.ID).ExtractErr())
	})

	t.Logf("Attempting to Create Server group")
	name := tools.RandomString("hss-group-", 3)
	err = hss.Create(client, hss.CreateOpts{
		Name: name,
		HostIds: []string{
			ecs.ID,
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Obtain Server group")
	getResp, err := hss.List(client, hss.ListOpts{
		Name: name,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, getResp[0].Name)
	th.AssertEquals(t, ecs.ID, getResp[0].HostIds[0])
	tools.PrintResource(t, getResp)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Server group")
		th.AssertNoErr(t, hss.Delete(client, hss.DeleteOpts{GroupID: getResp[0].ID}))
	})

	t.Logf("Attempting to Update Server group")
	err = hss.Update(client, hss.UpdateOpts{
		Name: name + "update",
		ID:   getResp[0].ID,
		HostIds: []string{
			ecs.ID,
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Obtain Server group after update")
	getUpdResp, err := hss.List(client, hss.ListOpts{
		Name: name,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name+"update", getUpdResp[0].Name)
	th.AssertEquals(t, ecs.ID, getUpdResp[0].HostIds[0])
	tools.PrintResource(t, getUpdResp)

	// Internal Server Error
	// t.Logf("Attempting to Change server Protection Status")
	// status, err := hss.ChangeProtectionStatus(client, hss.ProtectionOpts{
	// 	Version: "hss.version.enterprise",
	// 	HostIds: []string{
	// 		ecs.ID,
	// 	},
	// 	Tags: []tags.ResourceTag{
	// 		{
	// 			Key:   "muh",
	// 			Value: "kuh",
	// 		},
	// 		{
	// 			Key:   "muh2",
	// 			Value: "kuh2",
	// 		},
	// 	},
	// })
	// th.AssertNoErr(t, err)
	// th.AssertEquals(t, "hss.version.enterprise", status.Version)
	// th.AssertEquals(t, ecs.ID, status.HostIds[0])
}
