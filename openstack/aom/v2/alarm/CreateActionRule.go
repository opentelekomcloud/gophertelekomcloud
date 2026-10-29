package alarm

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateActionRuleOpts struct {
	// Alarm notification rule name.
	// Enter a maximum of 100 characters and do not start or end with a special character.
	// Only letters, digits, and underscores (_) are allowed.
	RuleName string `json:"rule_name" required:"true"`
	// Project ID.
	ProjectId string `json:"project_id" required:"true"`
	// Member account name.
	UserName string `json:"user_name" required:"true"`
	// Rule description.
	// Enter a maximum of 1,024 characters and do not start or end with an underscore (_).
	// Only digits, letters, underscores (_), asterisks (*), and spaces are allowed.
	Desc string `json:"desc,omitempty"`
	// Rule type.
	// 1: Prometheus monitoring.
	// 3: Log monitoring.
	Type string `json:"type" required:"true"`
	// Message template.
	NotificationTemplate string `json:"notification_template" required:"true"`
	// Creation time (UTC timestamp, in milliseconds).
	CreateTime int64 `json:"create_time,omitempty"`
	// Modification time (UTC timestamp, in milliseconds).
	UpdateTime int64 `json:"update_time,omitempty"`
	// Time zone.
	TimeZone string `json:"time_zone,omitempty"`
	// SMN topic. Maximum: 5.
	SmnTopics []SmnTopic `json:"smn_topics" required:"true"`
}

type SmnTopic struct {
	// Topic display name, which will be the name of an email sender.
	// Maximum: 192 bytes. This parameter is left blank by default.
	DisplayName string `json:"display_name,omitempty"`
	// Name of the topic. Enter 1 to 255 characters starting with a letter or digit.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	Name string `json:"name" required:"true"`
	// SMN message push policy. Options: 0 and 1.
	PushPolicy int `json:"push_policy" required:"true"`
	// Status of the topic subscriber.
	// 0: The topic has been deleted or the subscription list of this topic is empty.
	// 1: The subscription object is in the subscribed state.
	// 2: The subscription object is in the unsubscribed or canceled state.
	Status int `json:"status,omitempty"`
	// Unique resource identifier of the topic.
	TopicUrn string `json:"topic_urn" required:"true"`
}

// CreateActionRule is used to add an alarm action rule.
func CreateActionRule(client *golangsdk.ServiceClient, opts CreateActionRuleOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/alert/action-rules
	_, err = client.Post(client.ServiceURL("alert", "action-rules"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
