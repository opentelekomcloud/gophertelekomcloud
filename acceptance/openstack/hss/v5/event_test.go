package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/hss/v5/event"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/hss/v5/host"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/hss/v5/quota"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEventsLifecycle(t *testing.T) {
	// if os.Getenv("RUN_HSS_LIFECYCLE") == "" {
	// 	t.Skip("too slow to run in zuul")
	// }
	client, err := clients.NewHssClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Create member for Server group")
	ecsClient, err := clients.NewComputeV2Client()
	ecs := openstack.CreateServer(t, ecsClient,
		tools.RandomString("hss-group-member-", 3),
		"Standard_Debian_11_latest",
		"s2.large.2",
	)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete Server: %s", ecs.ID)
		th.AssertNoErr(t, servers.Delete(ecsClient, ecs.ID).ExtractErr())
	})

	err = golangsdk.WaitFor(1000, func() (bool, error) {
		h, err := host.ListHost(client, host.ListHostOpts{HostID: ecs.ID})
		if err != nil {
			return false, err
		}

		if len(h) == 1 {
			return true, nil
		}

		return false, nil
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Create Server group")
	name := tools.RandomString("hss-group-", 3)
	err = host.Create(client, host.CreateOpts{
		Name: name,
		HostIds: []string{
			ecs.ID,
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Obtain Server group")
	getResp, err := host.List(client, host.ListOpts{
		Name: name,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, getResp[0].Name)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Server group")
		th.AssertNoErr(t, host.Delete(client, host.DeleteOpts{GroupID: getResp[0].ID}))
	})

	t.Logf("Attempting to Change server Protection Status to null")
	_, err = host.ChangeProtectionStatus(client, host.ProtectionOpts{
		Version: "hss.version.null",
		HostIds: []string{
			ecs.ID,
		},
		Tags: []tags.ResourceTag{
			{
				Key:   "muh",
				Value: "kuh",
			},
			{
				Key:   "muh2",
				Value: "kuh2",
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Change server Protection Status to premium")
	_, err = host.ChangeProtectionStatus(client, host.ProtectionOpts{
		Version:      "hss.version.premium",
		ChargingMode: "on_demand",
		HostIds: []string{
			ecs.ID,
		},
		Tags: []tags.ResourceTag{
			{
				Key:   "muh",
				Value: "kuh",
			},
			{
				Key:   "muh2",
				Value: "kuh2",
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to get used quota details")
	q, err := quota.List(client, quota.ListOpts{
		HostName: ecs.Name,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "used", q[0].UsedStatus)
	th.AssertEquals(t, "hss.version.premium", q[0].Version)

	t.Logf("Attempting to get host events")
	listEventsResp, err := event.List(client, event.ListOpts{Category: "host"})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listEventsResp)

	t.Logf("Attempting to get host events")
	listWhitelistsResp, err := event.ListAlarmWhitelist(client, event.ListAlarmWhitelistOpts{})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listWhitelistsResp)
}
