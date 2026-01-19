package oneclickalarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains the options for enabling one-click monitoring.
type CreateOpts struct {
	// Specifies the one-click monitoring ID.
	// It can contain 1 to 64 alphanumeric characters.
	OneClickAlarmId string `json:"one_click_alarm_id" required:"true"`
	// Specifies the dimension names for metric and event alarm rules.
	DimensionNames DimensionNames `json:"dimension_names" required:"true"`
	// Specifies whether to enable the alarm notification.
	NotificationEnabled bool `json:"notification_enabled"`
	// Specifies the action to be triggered when an alarm is generated.
	// A maximum of 10 actions are supported.
	AlarmNotifications []Notification `json:"alarm_notifications,omitempty"`
	// Specifies the action to be triggered after an alarm is cleared.
	// A maximum of 10 actions are supported.
	OkNotifications []Notification `json:"ok_notifications,omitempty"`
	// Specifies the time when the alarm notification was enabled.
	// The value is in the format of HH:MM.
	NotificationBeginTime string `json:"notification_begin_time,omitempty"`
	// Specifies the time when the alarm notification was disabled.
	// The value is in the format of HH:MM.
	NotificationEndTime string `json:"notification_end_time,omitempty"`
}

// DimensionNames specifies dimension names for metric and event alarm rules.
type DimensionNames struct {
	// Specifies the dimension strings for metric alarm rules.
	// A maximum of 100 items are supported.
	Metric []string `json:"metric,omitempty"`
	// Specifies the dimension strings for event alarm rules.
	// A maximum of 100 items are supported.
	Event []string `json:"event,omitempty"`
}

// Notification represents an alarm notification.
type Notification struct {
	// Specifies the notification type.
	// Possible values: notification, contact
	Type string `json:"type" required:"true"`
	// Specifies the list of SMN topic URNs.
	// A maximum of 20 items are supported.
	NotificationList []string `json:"notification_list" required:"true"`
}

// Create enables one-click monitoring.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /v2/{project_id}/one-click-alarms
	raw, err := client.Post(client.ServiceURL("one-click-alarms"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		OneClickAlarmId string `json:"one_click_alarm_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.OneClickAlarmId, err
}
