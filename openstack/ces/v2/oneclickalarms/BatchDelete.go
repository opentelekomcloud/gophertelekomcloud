package oneclickalarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// BatchDeleteOpts contains the options for batch disabling one-click monitoring.
type BatchDeleteOpts struct {
	// Specifies the IDs of services that need to disable one-click monitoring.
	// A maximum of 100 IDs are supported.
	OneClickAlarmIds []string `json:"one_click_alarm_ids" required:"true"`
}

// BatchDelete batch disables one-click monitoring.
func BatchDelete(client *golangsdk.ServiceClient, opts BatchDeleteOpts) ([]string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/one-click-alarms/batch-delete
	raw, err := client.Post(client.ServiceURL("one-click-alarms", "batch-delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		OneClickAlarmIds []string `json:"one_click_alarm_ids"`
	}
	err = extract.Into(raw.Body, &res)
	return res.OneClickAlarmIds, err
}
