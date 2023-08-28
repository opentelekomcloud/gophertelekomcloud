package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetAntiTamper s used to query a web tamper protection rule by ID.
func GetAntiTamper(client *golangsdk.ServiceClient, policyId, ruleId string) (*AntiTamperRule, error) {
	// GET /v1/{project_id}/waf/policy/{policy_id}/antitamper/{rule_id}
	raw, err := client.Get(client.ServiceURL("waf", "policy", policyId, "antitamper", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res AntiTamperRule
	return &res, extract.Into(raw.Body, &res)
}
