package v2

import (
	"log"
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
	log.Printf("Create DIS App, Name: %s", appName)
	err = apps.CreateApp(client, apps.CreateAppOpts{
		AppName: appName,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		log.Printf("Delete DIS App, Name: %s", appName)
		err = apps.DeleteApp(client, appName)
		th.AssertNoErr(t, err)
	})

	log.Printf("Get DIS App, Name: %s", appName)
	app, err := apps.GetApp(client, appName)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, app)

	log.Print("List DIS Apps")
	listApps, err := apps.ListApps(client, apps.ListAppOpts{})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listApps)

	streamName := tools.RandomString("stream-create-test-", 3)
	log.Printf("Create DIS Stream, Name: %s", streamName)
	err = streams.CreateStream(client, streams.CreateStreamOpts{
		StreamName:     streamName,
		PartitionCount: 3,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		log.Printf("Delete DIS Stream, Name: %s", streamName)
		err = streams.DeleteStream(client, streamName)
		th.AssertNoErr(t, err)
	})

	log.Printf("Get DIS App status, Name: %s", appName)
	appStatus, err := apps.GetAppStatus(client, apps.GetAppStatusOpts{
		AppName:        appName,
		StreamName:     streamName,
		CheckpointType: "LAST_READ",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, appStatus)

	log.Printf("Get DIS Stream, Name: %s", streamName)
	getStream, err := streams.GetStream(client, streams.GetStreamOpts{
		StreamName: streamName,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getStream.StreamName, streamName)

	log.Printf("Update DIS Stream partitions count, Name: %s", streamName)
	err = streams.UpdatePartitionCount(client, streams.UpdatePartitionCountOpts{
		StreamName:           streamName,
		TargetPartitionCount: 4,
	})
	th.AssertNoErr(t, err)

	// "Bad request with: [PUT https://dis.eu-de.otc.t-systems.com/v2/5045c215010c440d91b2f7dce1f3753b/streams/stream-create-test-jmn],
	// error message: {\"errorCode\":\"DIS.4200\",\"message\":\"Invalid request. [Invalid target_partition_count null.]\"}"
	// https://jira.tsi-dev.otc-service.com/browse/BM-2472
	// err = streams.UpdateStream(client, streams.UpdateStreamOpts{
	// 	StreamName: streamName,
	// 	DataType:   "JSON",
	// })
	// th.AssertNoErr(t, err)

	// getStreamUpdated, err := streams.GetStream(client, streams.GetStreamOpts{
	// 	StreamName: streamName,
	// })
	// th.AssertNoErr(t, err)
	// th.AssertEquals(t, getStreamUpdated.StreamName, streamName)
	// th.AssertEquals(t, getStreamUpdated.DataType, "JSON")

	// url not found: https://jira.tsi-dev.otc-service.com/browse/BM-2474
	// log.Printf("Create DIS Stream Policy Rule, Name: %s", streamName)
	// err = streams.CreatePolicyRule(client, streams.CreatePolicyRuleOpts{
	// 	StreamName:    streamName,
	// 	StreamId:      getStream.StreamId,
	// 	PrincipalName: client.DomainID,
	// 	ActionType:    "putRecords",
	// 	Effect:        "effect",
	// })
	// th.AssertNoErr(t, err)
	//
	// log.Printf("Get DIS Stream Policy Rule, Name: %s", streamName)
	// rule, err := streams.GetPolicyRule(client, streamName)
	// th.AssertNoErr(t, err)
	// tools.PrintResource(t, rule)

	log.Print("List DIS Streams")
	listStreams, err := streams.ListStreams(client, streams.ListStreamsOpts{})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listStreams)

	log.Printf("Commit DIS App Checkpoint, Name: %s", appName)
	err = checkpoints.CommitCheckpoint(client, checkpoints.CommitCheckpointOpts{
		AppName:        appName,
		CheckpointType: "LAST_READ",
		StreamName:     streamName,
		PartitionId:    "0",
		SequenceNumber: "0",
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		log.Printf("Delete DIS App Checkpoint, Name: %s", appName)
		err = checkpoints.DeleteCheckpoint(client, checkpoints.DeleteCheckpointOpts{
			AppName:        appName,
			StreamName:     streamName,
			PartitionId:    "0",
			CheckpointType: "LAST_READ",
		})
		th.AssertNoErr(t, err)
	})

	log.Printf("Get DIS App Checkpoint, Name: %s", appName)
	checkpoint, err := checkpoints.GetCheckpoint(client, checkpoints.GetCheckpointOpts{
		AppName:        appName,
		StreamName:     streamName,
		PartitionId:    "0",
		CheckpointType: "LAST_READ",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, checkpoint)

	log.Printf("Create DIS Stream Data records, Name: %s", streamName)
	_, err = data.PutRecords(client, data.PutRecordsOpts{
		StreamName: streamName,
		StreamId:   getStream.StreamId,
		Records: []data.PutRecordsRequestEntry{
			{
				Data: "ZGF0YQ==",
			},
		},
	})
	th.AssertNoErr(t, err)

	log.Printf("Create DIS Stream cursor, Name: %s", streamName)
	cursor, err := data.GetCursor(client, data.GetCursorOpts{
		StreamName:  streamName,
		PartitionId: "0",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, cursor)

	log.Printf("Get DIS Stream records, Name: %s", streamName)
	records, err := data.GetRecords(client, data.GetRecordsOpts{
		PartitionCursor: cursor.PartitionCursor,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, records)
}
