package alarm

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListActionRules is used to query the alarm action rule list.
func ListActionRules(client *golangsdk.ServiceClient) ([]ActionRule, error) {
	// GET /v2/{project_id}/alert/action-rules
	raw, err := client.Get(client.ServiceURL("alert", "action-rules"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		ActionRules []ActionRule `json:"action_rules"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ActionRules, err
}

// ActionRule represents an alarm action rule returned by the list API.
type ActionRule struct {
	// Rule name.
	// Enter a maximum of 100 characters and do not start or end with a special character.
	// Only letters, digits, and underscores (_) are allowed.
	RuleName string `json:"rule_name"`
	// Project ID.
	ProjectId string `json:"project_id"`
	// Member account name.
	UserName string `json:"user_name"`
	// Rule description.
	// Enter a maximum of 1,024 characters and do not start or end with an underscore (_).
	// Only digits, letters, underscores (_), asterisks (*), and spaces are allowed.
	Desc string `json:"desc"`
	// Action type. Default: Notification.
	Type string `json:"type"`
	// Message template.
	NotificationTemplate string `json:"notification_template"`
	// Creation time (UTC timestamp, in milliseconds).
	CreateTime int64 `json:"create_time"`
	// Modification time (UTC timestamp, in milliseconds).
	UpdateTime int64 `json:"update_time"`
	// Time zone.
	TimeZone string `json:"time_zone"`
	// SMN topic. The total number of topics cannot exceed 5.
	SmnTopics []SmnTopic `json:"smn_topics"`
}
