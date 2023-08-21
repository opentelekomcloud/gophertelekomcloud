package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetKnownAttackSource is used to query a known attack source rule by ID.
func GetKnownAttackSource(client *golangsdk.ServiceClient, policyId, ruleId string) (*KnownAttackSourceRule, error) {
	// GET /v1/{project_id}/waf/policy/{policy_id}/punishment/{rule_id}
	raw, err := client.Get(client.ServiceURL("waf", "policy", policyId, "punishment", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res KnownAttackSourceRule
	return &res, extract.Into(raw.Body, &res)
}
