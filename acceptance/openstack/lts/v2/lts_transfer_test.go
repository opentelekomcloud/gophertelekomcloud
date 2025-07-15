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

func TestLtsTransferLifecycle(t *testing.T) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-group-", 3)
	createOpts := groups.CreateOpts{
		LogGroupName: name,
		TTLInDays:    7,
	}
	t.Logf("Attempting to Create Log Group")
	logId, err := groups.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Group")
		err = groups.Delete(client, logId)
		th.AssertNoErr(t, err)
	})

	sname := tools.RandomString("test-stream-", 3)
	t.Logf("Attempting to Create Log Stream")
	streamId, err := streams.Create(client, streams.CreateOpts{
		GroupId:       logId,
		LogStreamName: sname,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Stream")
		err = streams.Delete(client, streams.DeleteOpts{
			GroupId:  logId,
			StreamId: streamId,
		})
		th.AssertNoErr(t, err)
	})

	obsClient, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("obs-test", 5))

	t.Logf("Attempting to Create OBS Bucket")
	_, err = obsClient.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		t.Logf("Attempting to Delete OBS Bucket")
		_, err = obsClient.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Create Log Transfer")
	createTransferOpts := transfers.CreateOpts{
		LogGroupId: logId,
		LogStreams: []transfers.LogStreams{
			{
				LogStreamId: streamId,
			},
		},
		LogTransferInfo: &transfers.LogTransferInfo{
			LogTransferType:   "OBS",
			LogTransferMode:   "cycle",
			LogStorageFormat:  "JSON",
			LogTransferStatus: "ENABLE",
			LogTransferDetail: &transfers.TransferDetail{
				ObsPeriod:          3,
				ObsPeriodUnit:      "hour",
				ObsBucketName:      bucketName,
				ObsDirPreFixName:   "dir-test",
				ObsPrefixName:      "test",
				ObsTimeZone:        "UTC+01:00",
				ObsTimeZoneId:      "Africa/Lagos",
				ObsEncryptedEnable: false,
			},
		},
	}
	logDump, err := transfers.Create(client, createTransferOpts)
	th.AssertNoErr(t, err)
	t.Logf("Obs log dump created, id: %s", logDump.LogTransferId)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Transfer")
		err = transfers.Delete(client, logDump.LogTransferId)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to List Log Transfers")
	listLogs, err := transfers.List(client, transfers.ListTransfersOpts{})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listLogs)
	if len(listLogs) < 1 {
		t.Error("Log dump wasn't found")
	}

	t.Logf("Attempting to Update Log Transfer")
	updateTransferOpts := transfers.UpdateTransferOpts{
		TransferId: logDump.LogTransferId,
		TransferInfo: &transfers.TransferInfoUpdate{
			StorageFormat:  "RAW",
			TransferStatus: "ENABLE",
			TransferDetail: &transfers.TransferDetail{
				ObsPeriod:          3,
				ObsPeriodUnit:      "hour",
				ObsBucketName:      bucketName,
				ObsDirPreFixName:   "dir-test",
				ObsPrefixName:      "test",
				ObsTimeZone:        "UTC+01:00",
				ObsTimeZoneId:      "Africa/Lagos",
				ObsEncryptedEnable: false,
			},
		},
	}
	logDumpUpdated, err := transfers.Update(client, updateTransferOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "RAW", logDumpUpdated.LogTransferInfo.LogStorageFormat)
}
