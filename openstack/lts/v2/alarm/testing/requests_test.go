package testing

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/alarm"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestDeleteActiveAlarm(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleDeleteActiveAlarm(t, http.StatusOK, expectedDeleteActiveAlarmRequest)

	err := alarm.DeleteActiveAlarm(fake.ServiceClient(), "domain-id", deleteActiveAlarmOpts(1629947408497))
	th.AssertNoErr(t, err)
}

func TestDeleteActiveAlarmAllowsZeroStartsAt(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleDeleteActiveAlarm(t, http.StatusOK, `
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
      "starts_at": 0
    }
  ]
}`)

	err := alarm.DeleteActiveAlarm(fake.ServiceClient(), "domain-id", deleteActiveAlarmOpts(0))
	th.AssertNoErr(t, err)
}

func TestDeleteActiveAlarmRejectsMissingEvents(t *testing.T) {
	err := alarm.DeleteActiveAlarm(fake.ServiceClient(), "domain-id", alarm.DeleteActiveAlarmOpts{})
	if err == nil {
		t.Fatal("expected missing events to return an error")
	}
}

func TestDeleteActiveAlarmReturnsServerError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleDeleteActiveAlarm(t, http.StatusInternalServerError, expectedDeleteActiveAlarmRequest)

	err := alarm.DeleteActiveAlarm(fake.ServiceClient(), "domain-id", deleteActiveAlarmOpts(1629947408497))
	if err == nil {
		t.Fatal("expected server error")
	}
}
