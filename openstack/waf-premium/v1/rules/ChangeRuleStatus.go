package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ChangeStatusOpts struct {
	// Rule status. The value can be:
	// 0: The rule is disabled.
	// 1: The rule is enabled.
	Status int `json:"status" required:"true"`
}

// ChangeRuleStatus is used to change the status of a policy rule.
func ChangeRuleStatus(client *golangsdk.ServiceClient, PolicyId, Ruletype, RuleId string, opts ChangeStatusOpts) (*RuleStatus, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/waf/policy/{policy_id}/{ruletype}/{rule_id}/status}
	url := client.ServiceURL("waf", "policy", PolicyId, Ruletype, RuleId, "status")
	raw, err := client.Put(url, b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}

	var res RuleStatus
	return &res, extract.Into(raw.Body, &res)
}

type RuleStatus struct {
	// Rule ID.
	Id string `json:"id"`
	// Policy ID.
	PolicyId string `json:"policyid"`
	// Time when the rule was created.
	CreatedAt int64 `json:"timestamp"`
	// Rule Description.
	Description string `json:"description"`
	// Status. The options are 0 and 1. 0: Disabled. 1: Enabled.
	Status int `json:"status"`
}
