package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetBlacklist is used to query a blacklist or whitelist rule by ID.
func GetBlacklist(client *golangsdk.ServiceClient, policyId, ruleId string) (*BlacklistRule, error) {
	// GET /v1/{project_id}/waf/policy/{policy_id}/whiteblackip/{rule_id}
	raw, err := client.Get(client.ServiceURL("waf", "policy", policyId, "whiteblackip", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res BlacklistRule
	return &res, extract.Into(raw.Body, &res)
}
