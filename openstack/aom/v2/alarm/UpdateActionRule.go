package alarm

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateActionRuleOpts struct {
	// Rule name.
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
	// Action type. Default: Notification.
	Type string `json:"type" required:"true"`
	// Message template.
	NotificationTemplate string `json:"notification_template" required:"true"`
	// Creation time (UTC timestamp, in milliseconds).
	CreateTime int64 `json:"create_time,omitempty"`
	// Modification time (UTC timestamp, in milliseconds).
	UpdateTime int64 `json:"update_time,omitempty"`
	// Time zone.
	TimeZone string `json:"time_zone,omitempty"`
	// SMN topic. The total number of topics cannot exceed 5.
	SmnTopics []SmnTopic `json:"smn_topics" required:"true"`
}

// UpdateActionRule is used to modify an alarm action rule.
// The rule is identified by rule_name, project_id and user_name in the request body.
func UpdateActionRule(client *golangsdk.ServiceClient, opts UpdateActionRuleOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// PUT /v2/{project_id}/alert/action-rules
	_, err = client.Put(client.ServiceURL("alert", "action-rules"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
