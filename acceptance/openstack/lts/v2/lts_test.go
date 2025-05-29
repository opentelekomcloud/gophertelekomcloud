package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	rt "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/streams"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsLifecycle(t *testing.T) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	clientV1, err := clients.NewLtsV1Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-group-", 3)
	createOpts := groups.CreateOpts{
		LogGroupName: name,
		TTLInDays:    7,
	}
	t.Logf("Attempting to Create Log Group")
	created, err := groups.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Group")
		err = groups.Delete(client, created)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Update Log Group")
	group, err := groups.Update(client, groups.UpdateLogGroupOpts{
		LogGroupId: created,
		TTLInDays:  3,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, group.TTLInDays)

	t.Logf("Attempting to Add tag to Log Groups")
	err = tags.Manage(clientV1, "groups", created, tags.TagOpts{
		Action: "create",
		IsOpen: true,
		Tags: []rt.ResourceTag{
			{
				Key:   "TestKey",
				Value: "TestValue",
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to List Log Groups")
	got, err := groups.List(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, got)

	t.Logf("Attempting to Remove tag from Log Groups")
	err = tags.Manage(clientV1, "groups", created, tags.TagOpts{
		Action: "delete",
		IsOpen: true,
		Tags: []rt.ResourceTag{
			{
				Key:   "TestKey",
				Value: "TestValue",
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to List Log Groups")
	gotNew, err := groups.List(client)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(gotNew) > 0)
	th.AssertEquals(t, 1, len(gotNew[0].Tag))
	tools.PrintResource(t, got)

	t.Logf("Attempting to Create Log Stream")
	sname := tools.RandomString("test-stream-", 3)
	stream, err := streams.Create(client, streams.CreateOpts{
		GroupId:       created,
		LogStreamName: sname,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Stream")
		err = streams.Delete(client, streams.DeleteOpts{
			GroupId:  created,
			StreamId: stream,
		})
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Update Log Stream")
	update, err := streams.Update(client, streams.UpdateLogStreamOpts{
		GroupId:   created,
		StreamId:  stream,
		TTLInDays: 6,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, stream, update.LogStreamId)

	t.Logf("Attempting to List Log Streams")
	slist, err := streams.List(client, created)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(slist) > 0)
	tools.PrintResource(t, slist)

	t.Logf("Attempting to List of Log Streams filtered by Stream Name")
	byName, err := streams.ListStreams(client, streams.ListStreamsOpts{
		StreamName: update.LogStreamName,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(byName))
	tools.PrintResource(t, byName)
}
