package acl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to set the priority of an ACL protection rule.
func SetRulePriority(client *golangsdk.ServiceClient, ruleId string, opts OrderRuleAclDto) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/acl-rule/order/{acl_rule_id}
	raw, err := client.Put(client.ServiceURL("acl-rule", "order", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res SetRulePriorityResponse
	return &res.Data.ID, extract.Into(raw.Body, &res)
}

type SetRulePriorityResponse struct {
	// Data of the return value for updating priority
	Data OrderRuleId `json:"data"`
}

type OrderRuleId struct {
	// Rule ID
	ID string `json:"id"`
}
