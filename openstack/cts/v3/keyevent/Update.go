package keyevent

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateNotificationOpts struct {
	// Notification name.
	NotificationName string `json:"notification_name"`
	// Operation type. Possible options include complete and customized.
	// If you choose complete, you do not need to specify operations and notify_user_list,
	// and notifications will be sent when any supported operations occur on any of the connected cloud services.
	// If you choose customized, notifications will be sent when operations defined in operations occur.
	// Enumerated values:
	// complete
	// customized
	OperationType string `json:"operation_type"`
	// Operation list.
	Operations []Operations `json:"operations,omitempty"`
	// List of users whose operations will trigger notifications.
	// Currently, up to 50 users in 10 user groups can be configured.
	NotifyUserList []NotificationUsers `json:"notify_user_list,omitempty"`
	// Notification status. Possible options include enabled and disabled.
	// Enumerated values:
	// enabled
	// disabled
	Status string `json:"status"`
	// Topic URN.
	// To obtain the topic_urn, call the SMN API for querying topics.
	// Example URN: urn:smn:regionId:f96188c7ccaf4ffba0c9aa149ab2bd57:test_topic_v2
	TopicId string `json:"topic_id,omitempty"`
	// Notification ID.
	NotificationId string `json:"notification_id"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateNotificationOpts) (*NotificationResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/notifications
	raw, err := client.Put(client.ServiceURL("notifications"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*NotificationResponse, error) {
	if err != nil {
		return nil, err
	}

	var res NotificationResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
