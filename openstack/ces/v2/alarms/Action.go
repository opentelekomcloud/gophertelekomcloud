package alarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ActionOpts contains the options for enabling or disabling alarm rules in batches.
type ActionOpts struct {
	// Specifies the IDs of the alarm rules.
	// A maximum of 100 alarm rules can be enabled or disabled at a time.
	AlarmIds []string `json:"alarm_ids" required:"true"`
	// Specifies whether to enable the alarm rules.
	// true: enables alarm rules, false: disables alarm rules
	AlarmEnabled bool `json:"alarm_enabled"`
}

// ActionResponse contains the response from the Action request.
type ActionResponse struct {
	// Specifies the IDs of the alarm rules that are enabled or disabled.
	AlarmIds []string `json:"alarm_ids"`
}

// Action enables or disables alarm rules in batches.
func Action(client *golangsdk.ServiceClient, opts ActionOpts) (*ActionResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/alarms/action
	raw, err := client.Post(client.ServiceURL("alarms", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ActionResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
