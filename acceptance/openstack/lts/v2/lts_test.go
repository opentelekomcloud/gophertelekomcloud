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

	t.Logf("Attempting to List Log Groups")
	got, err := groups.List(client)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(got) > 0)
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
}
