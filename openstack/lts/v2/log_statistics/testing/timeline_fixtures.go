package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/log_statistics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	expectedTimelineRequest = `
{
  "start_time": 1668614400000,
  "end_time": 1668787200000,
  "search_type": "write",
  "period": 1,
  "resource_type": "tenant"
}`
	timelineResponse = `
{
  "results": [
    {
      "timestamp": 1669046400000,
      "value": 82485944.2
    },
    {
      "timestamp": 1669071600000,
      "value": 0
    }
  ]
}`
)

func timelineOpts() log_statistics.ListTimelineTrafficStatisticsOpts {
	startTime := int64(1668614400000)
	endTime := int64(1668787200000)
	return log_statistics.ListTimelineTrafficStatisticsOpts{
		Timezone:     "Asia/Shanghai",
		StartTime:    &startTime,
		EndTime:      &endTime,
		Period:       1,
		ResourceType: "tenant",
		SearchType:   "write",
	}
}

func handleTimeline(t *testing.T, status int, request, response string) {
	th.Mux.HandleFunc("/lts/timeline-traffic-statistics", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json;charset=UTF-8")
		th.AssertEquals(t, "Asia/Shanghai", r.URL.Query().Get("timezone"))
		th.TestJSONRequest(t, r, request)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = fmt.Fprint(w, response)
	})
}
