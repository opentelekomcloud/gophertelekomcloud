package monitor

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

func DeleteAlarmRule(client *golangsdk.ServiceClient, alarmId int64) (err error) {
	// DELETE /v2/{project_id}/ams/alarms/{alarm_id}
	_, err = client.Delete(client.ServiceURL("ams", "alarms", strconv.FormatInt(alarmId, 10)), nil)
	return
}
