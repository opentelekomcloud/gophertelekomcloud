package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetCustom is used to query a precise protection rule by ID.
func GetCustom(client *golangsdk.ServiceClient, policyId, ruleId string) (*CustomRule, error) {
	// GET /v1/{project_id}/waf/policy/{policy_id}/custom/{rule_id}
	raw, err := client.Get(client.ServiceURL("waf", "policy", policyId, "custom", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res CustomRule
	return &res, extract.Into(raw.Body, &res)
}
