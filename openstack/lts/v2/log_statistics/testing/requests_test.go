package testing

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/log_statistics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListTopNTrafficStatistics(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleTopN(t, http.StatusOK, expectedTopNRequest, topNResponse)

	actual, err := log_statistics.ListTopNTrafficStatistics(fake.ServiceClient(), topNOpts())
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []log_statistics.TrafficStatistic{
		{
			IndexTraffic:  6825703991,
			Storage:       15659303771,
			WriteTraffic:  1365140798.2,
			LogStreamID:   "a14dacb0-5a13-43a8-89a3-ea5424d95133",
			LogStreamName: "ELB",
		},
	}, actual)
}

func TestListTopNTrafficStatisticsAllowsMeaningfulZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleTopN(t, http.StatusOK, `
{
  "sort_by": "storage",
  "is_desc": false,
  "resource_type": "tenant",
  "filter": {},
  "start_time": 0,
  "end_time": 0,
  "search_list": ["storage"],
  "topn": 1
}`, `{"results":[]}`)

	zero := int64(0)
	desc := false
	filter := map[string]string{}
	actual, err := log_statistics.ListTopNTrafficStatistics(
		fake.ServiceClient(),
		log_statistics.ListTopNTrafficStatisticsOpts{
			EndTime:      &zero,
			IsDesc:       &desc,
			ResourceType: "tenant",
			SortBy:       "storage",
			StartTime:    &zero,
			TopN:         1,
			Filter:       &filter,
			SearchList:   []string{"storage"},
		},
	)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []log_statistics.TrafficStatistic{}, actual)
}

func TestListTopNTrafficStatisticsRejectsMissingInput(t *testing.T) {
	_, err := log_statistics.ListTopNTrafficStatistics(
		fake.ServiceClient(),
		log_statistics.ListTopNTrafficStatisticsOpts{},
	)
	if err == nil {
		t.Fatal("expected missing required input to return an error")
	}
}

func TestListTopNTrafficStatisticsReturnsBadRequestError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleTopN(
		t,
		http.StatusBadRequest,
		expectedTopNRequest,
		`{"errorCode":"LTS.0208","errorMessage":"The log stream does not existed"}`,
	)

	_, err := log_statistics.ListTopNTrafficStatistics(fake.ServiceClient(), topNOpts())
	if err == nil {
		t.Fatal("expected bad request error")
	}
}
