package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dis/v2/apps"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dis/v2/checkpoints"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dis/v2/data"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dis/v2/streams"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDIS(t *testing.T) {
	client, err := clients.NewDisV2Client()
	th.AssertNoErr(t, err)

	appName := tools.RandomString("app-create-test-", 3)
	err = apps.CreateApp(client, apps.CreateAppOpts{
		AppName: appName,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = apps.DeleteApp(client, appName)
		th.AssertNoErr(t, err)
	})

	app, err := apps.DescribeApp(client, appName)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, app)

	streamName := tools.RandomString("stream-create-test-", 3)
	err = streams.CreateStream(client, streams.CreateStreamOpts{
		StreamName:     streamName,
		PartitionCount: 3,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = streams.DeleteStream(client, streamName)
		th.AssertNoErr(t, err)
	})

	stream, err := streams.DescribeStream(client, streams.DescribeStreamOpts{
		StreamName: streamName,
	})
	th.AssertNoErr(t, err)

	err = streams.CreatePolicyRule(client, streams.CreatePolicyRuleOpts{
		StreamName:    streamName,
		StreamId:      stream.StreamId,
		PrincipalName: client.DomainID,
		ActionType:    "putRecords",
		Effect:        "effect",
	})
	th.AssertNoErr(t, err)

	rule, err := streams.GetPolicyRule(client, streamName)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, rule)

	err = streams.UpdatePartitionCount(client, streams.UpdatePartitionCountOpts{
		StreamName:           streamName,
		TargetPartitionCount: 4,
	})
	th.AssertNoErr(t, err)

	listStreams, err := streams.ListStreams(client, streams.ListStreamsOpts{})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listStreams)

	err = checkpoints.CommitCheckpoint(client, checkpoints.CommitCheckpointOpts{
		AppName:        appName,
		CheckpointType: "LAST_READ",
		StreamName:     streamName,
		PartitionId:    "0",
		SequenceNumber: "0",
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = checkpoints.DeleteCheckpoint(client, checkpoints.DeleteCheckpointOpts{
			AppName:        appName,
			StreamName:     streamName,
			PartitionId:    "0",
			CheckpointType: "LAST_READ",
		})
		th.AssertNoErr(t, err)
	})

	checkpoint, err := checkpoints.GetCheckpoint(client, checkpoints.GetCheckpointOpts{
		AppName:        appName,
		StreamName:     streamName,
		PartitionId:    "0",
		CheckpointType: "LAST_READ",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, checkpoint)

	_, err = data.PutRecords(client, data.PutRecordsOpts{
		StreamName: streamName,
		StreamId:   stream.StreamId,
		Records: []data.PutRecordsRequestEntry{
			{
				Data: "ZGF0YQ==",
			},
		},
	})
	th.AssertNoErr(t, err)

	cursor, err := data.GetCursor(client, data.GetCursorOpts{
		StreamName:  streamName,
		PartitionId: "0",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, cursor)

	records, err := data.GetRecords(client, data.GetRecordsOpts{
		PartitionCursor: cursor.PartitionCursor,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, records)
}
