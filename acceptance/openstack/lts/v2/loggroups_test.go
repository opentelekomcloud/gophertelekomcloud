package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/streams"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsLifecycle(t *testing.T) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-group-", 3)
	createOpts := groups.CreateOpts{
		LogGroupName: name,
		TTLInDays:    7,
	}

	created, err := groups.CreateLogGroup(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = groups.DeleteLogGroup(client, created)
		th.AssertNoErr(t, err)
	})

	group, err := groups.UpdateLogGroup(client, groups.UpdateLogGroupOpts{
		LogGroupId: created,
		TTLInDays:  3,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, group.TTLInDays)

	got, err := groups.ListLogGroups(client)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(got) > 0)
	tools.PrintResource(t, got)

	sname := tools.RandomString("test-group-", 3)
	stream, err := streams.CreateLogStream(client, streams.CreateOpts{
		GroupId:       created,
		LogStreamName: sname,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = streams.DeleteLogStream(client, streams.DeleteOpts{
			GroupId:  created,
			StreamId: stream,
		})
		th.AssertNoErr(t, err)
	})

	slist, err := streams.ListLogStream(client, created)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(slist) > 0)
	tools.PrintResource(t, slist)
}
