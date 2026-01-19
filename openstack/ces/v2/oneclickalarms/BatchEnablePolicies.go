package oneclickalarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// BatchEnablePoliciesOpts contains the options for batch enabling or disabling alarm policies.
type BatchEnablePoliciesOpts struct {
	// Specifies the IDs of alarm policies.
	// A maximum of 100 alarm policy IDs are supported.
	AlarmPolicyIds []string `json:"alarm_policy_ids" required:"true"`
	// Specifies whether to enable the alarm policies.
	// true: enable, false: disable
	Enabled bool `json:"enabled"`
}

// BatchEnablePolicies batch enables or disables alarm policies in alarm rules
// for one service with one-click monitoring enabled.
func BatchEnablePolicies(client *golangsdk.ServiceClient, oneClickAlarmId, alarmId string, opts BatchEnablePoliciesOpts) ([]string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/one-click-alarms/{one_click_alarm_id}/alarms/{alarm_id}/policies/action
	raw, err := client.Put(client.ServiceURL("one-click-alarms", oneClickAlarmId, "alarms", alarmId, "policies", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		AlarmPolicyIds []string `json:"alarm_policy_ids"`
	}
	err = extract.Into(raw.Body, &res)
	return res.AlarmPolicyIds, err
}
