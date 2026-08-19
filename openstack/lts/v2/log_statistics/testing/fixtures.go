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
	expectedTopNRequest = `
{
  "sort_by": "storage",
  "is_desc": true,
  "resource_type": "log_stream",
  "filter": {},
  "start_time": 1668668183969,
  "end_time": 1669272983969,
  "search_list": ["index", "write", "storage"],
  "topn": 100
}`
	topNResponse = `
{
  "results": [
    {
      "index_traffic": 6825703991,
      "log_stream_id": "a14dacb0-5a13-43a8-89a3-ea5424d95133",
      "log_stream_name": "ELB",
      "storage": 15659303771,
      "write_traffic": 1365140798.2
    }
  ]
}`
)

func topNOpts() log_statistics.ListTopNTrafficStatisticsOpts {
	startTime := int64(1668668183969)
	endTime := int64(1669272983969)
	isDesc := true
	filter := map[string]string{}
	return log_statistics.ListTopNTrafficStatisticsOpts{
		EndTime:      &endTime,
		IsDesc:       &isDesc,
		ResourceType: "log_stream",
		SortBy:       "storage",
		StartTime:    &startTime,
		TopN:         100,
		Filter:       &filter,
		SearchList:   []string{"index", "write", "storage"},
	}
}

func handleTopN(t *testing.T, status int, request, response string) {
	th.Mux.HandleFunc("/lts/topn-traffic-statistics", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json;charset=UTF-8")
		th.TestJSONRequest(t, r, request)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = fmt.Fprint(w, response)
	})
}
