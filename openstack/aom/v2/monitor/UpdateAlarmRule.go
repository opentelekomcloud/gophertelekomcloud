package monitor

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func UpdateAlarmRule(client *golangsdk.ServiceClient, opts Threshold) (*UpdateAlarmRuleResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/ams/alarms
	raw, err := client.Put(client.ServiceURL("ams", "alarms"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateAlarmRuleResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type UpdateAlarmRuleResponse struct {
	// Response code.
	ErrorCode string `json:"errorCode"`
	// Response message.
	ErrorMessage string `json:"errorMessage"`
	// Threshold rule code.
	AlarmId *int64 `json:"alarmId"`
}
