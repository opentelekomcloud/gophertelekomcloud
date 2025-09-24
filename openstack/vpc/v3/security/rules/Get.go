package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to obtain details about a security group rule.
func Get(client *golangsdk.ServiceClient, ruleId string) (*SecurityGroupRule, error) {
	// GET /v3/{project_id}/vpc/security-group-rules/{security_group_rule_id}
	raw, err := client.Get(client.ServiceURL("security-group-rules", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res SecurityGroupRuleResponse
	err = extract.Into(raw.Body, &res)
	return &res.SecurityGroupRule, err
}
