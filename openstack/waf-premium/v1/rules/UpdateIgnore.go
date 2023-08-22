package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateIgnore is used to update a global protection whitelist (false alarm masking) rule.
func UpdateIgnore(client *golangsdk.ServiceClient, policyId, ruleId string, opts CreateIgnoreOpts) (*IgnoreRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/waf/policy/{policy_id}/ignore/{rule_id}
	raw, err := client.Put(client.ServiceURL("waf", "policy", policyId, "ignore", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}
	var res IgnoreRule
	return &res, extract.Into(raw.Body, &res)
}
