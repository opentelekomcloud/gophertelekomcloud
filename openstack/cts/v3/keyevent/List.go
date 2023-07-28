package keyevent

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListNotificationsOpts struct {
	// Notification type.
	// Enumerated value:
	// smn
	NotificationType string
	// Notification name. If this parameter is not specified, all key event notifications configured in the current tenant account are returned.
	NotificationName string `q:"notification_name,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListNotificationsOpts) ([]NotificationResponse, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET /v3/{project_id}/notifications/{notification_type}
	url := client.ServiceURL("notifications", opts.NotificationType) + q.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []NotificationResponse
	err = extract.IntoSlicePtr(raw.Body, &res, "notifications")
	return res, err
}

type NotificationResponse struct {
	// Notification name.
	NotificationName string `json:"notification_name,omitempty"`
	// Operation type. Possible options include complete and customized.
	// Enumerated values:
	// customized
	// complete
	OperationType string `json:"operation_type,omitempty"`
	// Operation list.
	Operations []Operations `json:"operations,omitempty"`
	// List of users whose operations will trigger notifications.
	// Currently, up to 50 users in 10 user groups can be configured.
	NotifyUserList []NotificationUsers `json:"notify_user_list,omitempty"`
	// Notification status. Possible options include enabled and disabled.
	// Enumerated values:
	// enabled
	// disabled
	Status string `json:"status,omitempty"`
	// Unique resource ID of an SMN topic. You can obtain the ID by querying the topic list.
	TopicId string `json:"topic_id,omitempty"`
	// Unique notification ID.
	NotificationId string `json:"notification_id,omitempty"`
	// Notification type.
	// Enumerated value:
	// smn
	NotificationType string `json:"notification_type,omitempty"`
	// Project ID.
	ProjectId string `json:"project_id,omitempty"`
	// Time when a notification rule was created.
	CreateTime int64 `json:"create_time,omitempty"`
}
