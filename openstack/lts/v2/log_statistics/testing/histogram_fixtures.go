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
	expectedHistogramRequest = `
{
  "group_id": "00330565-5baf-4e0d-bd16-ba0c6b951d9a",
  "stream_id": "715cda3b-e17f-492a-a6ca-98a1ba16ad8c",
  "end_time": "1637820813605",
  "start_time": "1637817213605",
  "key_word": "test",
  "step_interval": 6000,
  "is_iterative": false
}`
	histogramResponse = `
{
  "count": 1,
  "histogram": [
    {
      "num": 1,
      "startTime": 1637821594579,
      "endTime": 1637821595000
    },
    {
      "num": 0,
      "startTime": 1637821654000,
      "endTime": 1637821654579
    }
  ],
  "isQueryComplete": true
}`
)

func histogramOpts() log_statistics.ListLogHistogramOpts {
	iterative := false
	return log_statistics.ListLogHistogramOpts{
		StartTime:    "1637817213605",
		EndTime:      "1637820813605",
		StepInterval: 6000,
		GroupID:      "00330565-5baf-4e0d-bd16-ba0c6b951d9a",
		StreamID:     "715cda3b-e17f-492a-a6ca-98a1ba16ad8c",
		Keyword:      "test",
		IsIterative:  &iterative,
	}
}

func handleHistogram(t *testing.T, status int, response string) {
	th.Mux.HandleFunc("/lts/keyword-count", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json;charset=UTF-8")
		th.TestJSONRequest(t, r, expectedHistogramRequest)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = fmt.Fprint(w, response)
	})
}
