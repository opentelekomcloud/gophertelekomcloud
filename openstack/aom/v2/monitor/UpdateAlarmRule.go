package monitor

import "github.com/opentelekomcloud/gophertelekomcloud"

func UpdateAlarmRule(client *golangsdk.ServiceClient, opts Threshold) (*UpdateAlarmRuleResponse, error) {
	// PUT /v2/{project_id}/ams/alarms
}

type UpdateAlarmRuleResponse struct {
	// Response code.
	ErrorCode string `json:"errorCode"`
	// Response message.
	ErrorMessage string `json:"errorMessage"`
	// Threshold rule code.
	AlarmId *int64 `json:"alarmId"`
}
