package testing

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/alarm"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const expectedDeleteActiveAlarmRequest = `
{
  "events": [
    {
      "metadata": {
        "event_type": "alarm",
        "event_id": "1",
        "event_severity": "Critical",
        "event_name": "demo",
        "resource_type": "Log group/stream.",
        "resource_id": "lts-group-demo/lts-topic-demo",
        "resource_provider": "LTS",
        "lts_alarm_type": "keywords/sql"
      },
      "starts_at": 1629947408497
    }
  ]
}`

func deleteActiveAlarmOpts(startsAt int64) alarm.DeleteActiveAlarmOpts {
	return alarm.DeleteActiveAlarmOpts{
		Events: []alarm.DeleteActiveAlarmEvent{
			{
				Metadata: &alarm.DeleteActiveAlarmMetadata{
					EventType:        "alarm",
					EventID:          "1",
					EventSeverity:    "Critical",
					EventName:        "demo",
					ResourceType:     "Log group/stream.",
					ResourceID:       "lts-group-demo/lts-topic-demo",
					ResourceProvider: "LTS",
					LTSAlarmType:     "keywords/sql",
				},
				StartsAt: &startsAt,
			},
		},
	}
}

func handleDeleteActiveAlarm(t *testing.T, status int, expectedBody string) {
	th.Mux.HandleFunc("/domain-id/lts/alarms/sql-alarm/clear", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json;charset=UTF-8")
		th.TestJSONRequest(t, r, expectedBody)
		w.WriteHeader(status)
	})
}
