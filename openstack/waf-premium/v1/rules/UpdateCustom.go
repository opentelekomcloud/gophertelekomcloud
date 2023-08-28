package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateCustom is used to update a precise protection rule.
func UpdateCustom(client *golangsdk.ServiceClient, policyId, ruleId string, opts CreateCustomOpts) (*CustomRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/waf/policy/{policy_id}/custom/{rule_id}
	raw, err := client.Put(client.ServiceURL("waf", "policy", policyId, "custom", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}
	var res CustomRule
	return &res, extract.Into(raw.Body, &res)
}
