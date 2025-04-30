package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/log"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/streams"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLogsLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	loadbalancerID := createLoadBalancer(t, client)
	t.Cleanup(func() {
		t.Logf("Attempting to delete ELBv3 Loadbalancer: %s", loadbalancerID)
		deleteLoadbalancer(t, client, loadbalancerID)
		t.Logf("Deleted ELBv3 Loadbalancer: %s", loadbalancerID)
	})

	clientV2, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-group-", 3)
	t.Logf("Attempting to Create Log Group")
	group, err := groups.Create(clientV2, groups.CreateOpts{
		LogGroupName: name,
		TTLInDays:    7,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Group")
		err = groups.Delete(clientV2, group)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Create Log Stream")
	sname := tools.RandomString("test-stream-", 3)
	stream, err := streams.Create(clientV2, streams.CreateOpts{
		GroupId:       group,
		LogStreamName: sname,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Stream")
		err = streams.Delete(clientV2, streams.DeleteOpts{
			GroupId:  group,
			StreamId: stream,
		})
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Create Loadbalancer Log")
	logtank, err := log.Create(client, log.CreateOpts{
		LoadbalancerId: loadbalancerID,
		LogGroupId:     group,
		LogStreamId:    stream,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Loadbalancer Log")
		err = log.Delete(client, logtank.Logtank.ID)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to List Loadbalancer Log")
	listLog, err := log.List(client, log.ListOpts{
		LoadbalancerId: []string{loadbalancerID},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(listLog))
	th.AssertEquals(t, stream, listLog[0].LogStreamId)

	t.Logf("Attempting to Create Second Log Stream")
	snameSec := tools.RandomString("test-second-stream-", 3)
	streamSec, err := streams.Create(clientV2, streams.CreateOpts{
		GroupId:       group,
		LogStreamName: snameSec,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Second Log Stream")
		err = streams.Delete(clientV2, streams.DeleteOpts{
			GroupId:  group,
			StreamId: streamSec,
		})
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Update Loadbalancer Log")
	err = log.Update(client, logtank.Logtank.ID, log.UpdateOpts{
		LogStreamId: streamSec,
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get Loadbalancer Log")
	get, err := log.Get(client, logtank.Logtank.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, streamSec, get.Logtank.LogStreamId)
}
