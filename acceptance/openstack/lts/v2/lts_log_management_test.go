package v2

import (
	"strconv"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/alarm"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/favorites"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/log_statistics"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/streams"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsFavoriteLifecycle(t *testing.T) {
	clientV2, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)
	clientV10, err := clients.NewLtsV10Client()
	th.AssertNoErr(t, err)
	groupID, streamID := createLtsLogResource(t, clientV2)

	t.Log("Attempting to add a log stream to favorites")
	favorite, err := favorites.Create(clientV10, favorites.CreateOpts{
		ResourceID:   streamID,
		ResourceType: "log_stream",
		LogGroupID:   groupID,
		LogStreamID:  streamID,
		IsGlobal:     true,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, streamID, favorite.ResourceID)

	t.Cleanup(func() {
		t.Log("Attempting to remove the log stream from favorites")
		err := favorites.Delete(clientV10, favorites.DeleteOpts{ResourceID: favorite.ResourceID})
		th.AssertNoErr(t, err)
	})
}

func TestLtsLogStatistics(t *testing.T) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)
	groupID, streamID := createLtsLogResource(t, client)
	endTime := time.Now().UnixMilli()
	startTime := endTime - int64(time.Hour/time.Millisecond)

	t.Log("Attempting to collect top-N traffic statistics")
	filter := map[string]string{"log_stream_id": streamID}
	topN, err := log_statistics.ListTopNTrafficStatistics(
		client,
		log_statistics.ListTopNTrafficStatisticsOpts{
			StartTime:    &startTime,
			EndTime:      &endTime,
			IsDesc:       pointerto.Bool(true),
			ResourceType: "log_stream",
			SortBy:       "storage",
			TopN:         5,
			Filter:       &filter,
			SearchList:   []string{"index", "write", "storage"},
		},
	)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, topN)

	t.Log("Attempting to query timeline traffic statistics")
	timeline, err := log_statistics.ListTimelineTrafficStatistics(
		client,
		log_statistics.ListTimelineTrafficStatisticsOpts{
			Timezone:     "UTC",
			StartTime:    &startTime,
			EndTime:      &endTime,
			Period:       1,
			ResourceType: "log_stream",
			ResourceID:   streamID,
			SearchType:   "write",
		},
	)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, timeline)

	t.Log("Attempting to query the log histogram")
	histogram, err := log_statistics.ListLogHistogram(
		client,
		log_statistics.ListLogHistogramOpts{
			StartTime:    strconv.FormatInt(startTime, 10),
			EndTime:      strconv.FormatInt(endTime, 10),
			StepInterval: 60000,
			GroupID:      groupID,
			StreamID:     streamID,
			Keyword:      "acceptance",
		},
	)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, histogram)
}

func TestLtsDeleteActiveAlarm(t *testing.T) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)
	startsAt := time.Now().UnixMilli()

	t.Log("Attempting to clear a synthetic active alarm")
	err = alarm.DeleteActiveAlarm(
		client,
		client.DomainID,
		alarm.DeleteActiveAlarmOpts{
			Events: []alarm.DeleteActiveAlarmEvent{
				{
					Metadata: &alarm.DeleteActiveAlarmMetadata{
						EventType:        "alarm",
						EventID:          tools.RandomString("acceptance-", 8),
						EventSeverity:    "Info",
						EventName:        "acceptance-test",
						ResourceType:     "log_stream",
						ResourceID:       "acceptance-test",
						ResourceProvider: "LTS",
						LTSAlarmType:     "keywords",
					},
					StartsAt: &startsAt,
				},
			},
		},
	)
	th.AssertNoErr(t, err)
}

func createLtsLogResource(t *testing.T, client *golangsdk.ServiceClient) (string, string) {
	t.Helper()

	groupID, err := groups.Create(client, groups.CreateOpts{
		LogGroupName: tools.RandomString("acceptance-group-", 5),
		TTLInDays:    7,
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Log("Attempting to delete the acceptance log group")
		err := groups.Delete(client, groupID)
		th.AssertNoErr(t, err)
	})

	streamID, err := streams.Create(client, streams.CreateOpts{
		GroupId:       groupID,
		LogStreamName: tools.RandomString("acceptance-stream-", 5),
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Log("Attempting to delete the acceptance log stream")
		err := streams.Delete(client, streams.DeleteOpts{
			GroupId:  groupID,
			StreamId: streamID,
		})
		th.AssertNoErr(t, err)
	})

	return groupID, streamID
}
