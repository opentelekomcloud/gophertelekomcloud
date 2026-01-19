package oneclickalarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateNotificationsOpts contains the options for batch modifying alarm notifications.
type UpdateNotificationsOpts struct {
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

// UpdateNotifications batch modifies alarm notifications in alarm rules
// for one service with one-click monitoring enabled.
func UpdateNotifications(client *golangsdk.ServiceClient, oneClickAlarmId string, opts UpdateNotificationsOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// PUT /v2/{project_id}/one-click-alarms/{one_click_alarm_id}/notifications
	_, err = client.Put(client.ServiceURL("one-click-alarms", oneClickAlarmId, "notifications"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
