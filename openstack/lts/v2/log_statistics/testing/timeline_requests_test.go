package testing

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/log_statistics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListTimelineTrafficStatistics(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleTimeline(t, http.StatusOK, expectedTimelineRequest, timelineResponse)

	actual, err := log_statistics.ListTimelineTrafficStatistics(fake.ServiceClient(), timelineOpts())
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []log_statistics.TimelineStatistic{
		{Timestamp: 1669046400000, Value: 82485944.2},
		{Timestamp: 1669071600000, Value: 0},
	}, actual)
}

func TestListTimelineTrafficStatisticsAllowsZeroTimestamps(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleTimeline(t, http.StatusOK, `
{
  "start_time": 0,
  "end_time": 0,
  "search_type": "storage",
  "period": 1,
  "resource_type": "tenant"
}`, `{"results":[]}`)

	zero := int64(0)
	actual, err := log_statistics.ListTimelineTrafficStatistics(
		fake.ServiceClient(),
		log_statistics.ListTimelineTrafficStatisticsOpts{
			Timezone:     "Asia/Shanghai",
			StartTime:    &zero,
			EndTime:      &zero,
			Period:       1,
			ResourceType: "tenant",
			SearchType:   "storage",
		},
	)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []log_statistics.TimelineStatistic{}, actual)
}

func TestListTimelineTrafficStatisticsRejectsMissingInput(t *testing.T) {
	_, err := log_statistics.ListTimelineTrafficStatistics(
		fake.ServiceClient(),
		log_statistics.ListTimelineTrafficStatisticsOpts{},
	)
	if err == nil {
		t.Fatal("expected missing required input to return an error")
	}
}

func TestListTimelineTrafficStatisticsReturnsBadRequestError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleTimeline(
		t,
		http.StatusBadRequest,
		expectedTimelineRequest,
		`{"errorCode":"LTS.0009","errorMessage":"resource_id must not be empty"}`,
	)

	_, err := log_statistics.ListTimelineTrafficStatistics(fake.ServiceClient(), timelineOpts())
	if err == nil {
		t.Fatal("expected bad request error")
	}
}
