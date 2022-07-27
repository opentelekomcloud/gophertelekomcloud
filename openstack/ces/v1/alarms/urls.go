package alarms

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

const alarms = "alarms"

// GET /V1.0/{project_id}/alarms
// POST /V1.0/{project_id}/alarms
func alarmsURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(alarms)
}

// GET /V1.0/{project_id}/alarms/{alarm_id}
// DELETE /V1.0/{project_id}/alarms/{alarm_id}
func alarmIdURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(alarms, id)
}

// PUT /V1.0/{project_id}/alarms/{alarm_id}/action
func alarmActionURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(alarms, id, "action")
}
