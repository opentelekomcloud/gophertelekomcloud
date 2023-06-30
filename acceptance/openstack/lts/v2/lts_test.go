package v2

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/streams"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/transfers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
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

	sname := tools.RandomString("test-stream-", 3)
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

func TestLtsTransferLifecycle(t *testing.T) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-group-", 3)
	createOpts := groups.CreateOpts{
		LogGroupName: name,
		TTLInDays:    7,
	}

	logId, err := groups.CreateLogGroup(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = groups.DeleteLogGroup(client, logId)
		th.AssertNoErr(t, err)
	})

	sname := tools.RandomString("test-stream-", 3)
	streamId, err := streams.CreateLogStream(client, streams.CreateOpts{
		GroupId:       logId,
		LogStreamName: sname,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = streams.DeleteLogStream(client, streams.DeleteOpts{
			GroupId:  logId,
			StreamId: streamId,
		})
		th.AssertNoErr(t, err)
	})

	obsClient, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("obs-test", 5))

	_, err = obsClient.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		_, err = obsClient.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	switchOn := false
	createTransferOpts := transfers.CreateLogDumpObsOpts{
		LogGroupId: logId,
		LogStreamIds: []string{
			streamId,
		},
		ObsBucketName: bucketName,
		Type:          "cycle",
		StorageFormat: "RAW",
		SwitchOn:      &switchOn,
		PrefixName:    "test",
		DirPrefixName: "dir-test",
		Period:        3,
		PeriodUnit:    "hour",
	}
	logDumpId, err := transfers.CreateLogDumpObs(client, createTransferOpts)
	th.AssertNoErr(t, err)
	t.Logf("Obs log dump created, id: %s", logDumpId)

	// GET API is currently not working

	listLogs, err := transfers.ListTransfers(client, transfers.ListTransfersOpts{})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listLogs)
	if len(listLogs) < 1 {
		t.Error("Log dump wasn't found")
	}

	err = transfers.DeleteTransfer(client, logDumpId)
	th.AssertNoErr(t, err)
}
