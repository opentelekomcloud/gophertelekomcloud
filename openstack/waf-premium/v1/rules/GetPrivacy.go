package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetPrivacy is used to query a data masking rule by ID.
func GetPrivacy(client *golangsdk.ServiceClient, policyId, ruleId string) (*PrivacyRule, error) {
	// GET /v1/{project_id}/waf/policy/{policy_id}/privacy/{rule_id}
	raw, err := client.Get(client.ServiceURL("waf", "policy", policyId, "privacy", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res PrivacyRule
	return &res, extract.Into(raw.Body, &res)
}
