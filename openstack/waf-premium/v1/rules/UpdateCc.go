package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateCc is used to update a CC attack protection rule.
func UpdateCc(client *golangsdk.ServiceClient, policyId, ruleId string, opts CreateCcOpts) (*CcRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/waf/policy/{policy_id}/cc/{rule_id}
	raw, err := client.Put(client.ServiceURL("waf", "policy", policyId, "cc", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}
	var res CcRule
	return &res, extract.Into(raw.Body, &res)
}
