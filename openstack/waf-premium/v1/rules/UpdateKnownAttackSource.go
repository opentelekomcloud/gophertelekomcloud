package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateKnownAttackSourceOpts struct {
	// Block duration, in seconds.
	// If prefix long is selected for the rule type, the value for block_time ranges from 301 to 1800.
	// If prefix short is selected for the rule type, the value for block_time ranges from 0 to 300.
	BlockTime string `json:"block_time" required:"true"`
	// Rule description
	Description string `json:"description"`
}

// UpdateKnownAttackSource is used update a known attack source rule.
func UpdateKnownAttackSource(client *golangsdk.ServiceClient, policyId, ruleId string, opts UpdateKnownAttackSourceOpts) (*KnownAttackSourceRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/waf/policy/{policy_id}/punishment/{rule_id}
	raw, err := client.Put(client.ServiceURL("waf", "policy", policyId, "punishment", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}
	var res KnownAttackSourceRule
	return &res, extract.Into(raw.Body, &res)
}
