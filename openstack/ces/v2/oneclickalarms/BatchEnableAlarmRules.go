package oneclickalarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// BatchEnableAlarmRulesOpts contains the options for batch enabling or disabling alarm rules.
type BatchEnableAlarmRulesOpts struct {
	// Specifies the IDs of alarm rules.
	// A maximum of 100 alarm rule IDs are supported.
	AlarmIds []string `json:"alarm_ids" required:"true"`
	// Specifies whether to enable the alarm rules.
	// true: enable, false: disable
	AlarmEnabled bool `json:"alarm_enabled"`
}

// BatchEnableAlarmRules batch enables or disables alarm rules of one service in one-click monitoring.
func BatchEnableAlarmRules(client *golangsdk.ServiceClient, oneClickAlarmId string, opts BatchEnableAlarmRulesOpts) ([]string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/one-click-alarms/{one_click_alarm_id}/alarm-rules/action
	raw, err := client.Put(client.ServiceURL("one-click-alarms", oneClickAlarmId, "alarm-rules", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		AlarmIds []string `json:"alarm_ids"`
	}
	err = extract.Into(raw.Body, &res)
	return res.AlarmIds, err
}
