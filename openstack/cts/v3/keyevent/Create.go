package keyevent

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateNotificationOpts struct {
	// Notification name.
	NotificationName string `json:"notification_name"`
	// Operation type. Possible options include complete and customized.
	// If you choose complete, you do not need to specify operations and notify_user_list,
	// and notifications will be sent when any supported operations occur on any of the connected cloud services.
	// If you choose customized, notifications will be sent when operations defined in operations occur.
	// Enumerated values:
	// 	complete
	// 	customized
	OperationType OperationType `json:"operation_type"`
	// Operation list.
	Operations []Operations `json:"operations,omitempty"`
	// List of users whose operations will trigger notifications. Currently, up to 50 users in 10 user groups can be configured.
	NotifyUserList []NotificationUsers `json:"notify_user_list,omitempty"`
	// Topic URN.
	// 	To obtain the topic_urn, call the SMN API for querying topics.
	// Example URN: urn:smn:regionId:f96188c7ccaf4ffba0c9aa149ab2bd57:test_topic_v2
	TopicId string `json:"topic_id,omitempty"`
}

type OperationType string

const (
	Complete   OperationType = "complete"
	Customized OperationType = "customized"
)

type Operations struct {
	// Cloud service. The value must be the acronym of a cloud service that has been connected with CTS. It is a word composed of uppercase letters.
	// For cloud services that can be connected with CTS, see section "Supported Services and Operations" in the Cloud Trace Service User Guide.
	ServiceType string `json:"service_type"`
	// Resource type.
	ResourceType string `json:"resource_type"`
	// Trace name.
	TraceNames []string `json:"trace_names"`
}

type NotificationUsers struct {
	// IAM user group.
	UserGroup string `json:"user_group"`
	// IAM user.
	UserList []string `json:"user_list"`
}

func Create(client *golangsdk.ServiceClient, opts CreateNotificationOpts) (*NotificationResponse, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/notifications
	raw, err := client.Post(client.ServiceURL("notifications"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res NotificationResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateNotificationStatus string

const (
	Enabled  CreateNotificationStatus = "enabled"
	Disabled CreateNotificationStatus = "disabled"
)

type NotificationType string

const (
	Smn NotificationType = "smn"
)
