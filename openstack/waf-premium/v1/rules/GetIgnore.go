package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetIgnore is used to querying a global protection whitelist (formerly false alarm masking) rule by ID.
func GetIgnore(client *golangsdk.ServiceClient, policyId, ruleId string) (*IgnoreRule, error) {
	// GET /v1/{project_id}/waf/policy/{policy_id}/ignore/{rule_id}
	raw, err := client.Get(client.ServiceURL("waf", "policy", policyId, "ignore", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res IgnoreRule
	return &res, extract.Into(raw.Body, &res)
}
