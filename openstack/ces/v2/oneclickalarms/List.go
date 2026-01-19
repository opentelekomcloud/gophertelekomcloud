package oneclickalarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// OneClickAlarm represents a one-click monitoring configuration.
type OneClickAlarm struct {
	// Specifies the one-click monitoring ID.
	OneClickAlarmId string `json:"one_click_alarm_id"`
	// Specifies the namespace of a service.
	Namespace string `json:"namespace"`
	// Provides supplementary information about the one-click monitoring.
	Description string `json:"description"`
	// Specifies whether one-click monitoring is enabled.
	Enabled bool `json:"enabled"`
}

// List queries services and resources in one-click monitoring.
func List(client *golangsdk.ServiceClient) ([]OneClickAlarm, error) {
	// GET /v2/{project_id}/one-click-alarms
	raw, err := client.Get(client.ServiceURL("one-click-alarms"), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		OneClickAlarms []OneClickAlarm `json:"one_click_alarms"`
	}
	err = extract.Into(raw.Body, &res)
	return res.OneClickAlarms, err
}
