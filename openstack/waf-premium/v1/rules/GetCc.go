package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetCc is used to query a CC attack protection rule by ID.
func GetCc(client *golangsdk.ServiceClient, policyId, ruleId string) (*CcRule, error) {
	// GET /v1/{project_id}/waf/policy/{policy_id}/cc/{rule_id}
	raw, err := client.Get(client.ServiceURL("waf", "policy", policyId, "cc", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res CcRule
	return &res, extract.Into(raw.Body, &res)
}
