package alarms

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func DeleteAlarm(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /V1.0/{project_id}/alarms/{alarm_id}
	_, err = client.Delete(client.ServiceURL("alarms", id), nil)
	return
}
