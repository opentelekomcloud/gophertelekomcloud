package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetAntiLeakage is used to query an information leakage prevention rule by ID.
func GetAntiLeakage(client *golangsdk.ServiceClient, policyId, ruleId string) (*AntiLeakageRule, error) {
	// GET /v1/{project_id}/waf/policy/{policy_id}/antileakage/{rule_id}
	raw, err := client.Get(client.ServiceURL("waf", "policy", policyId, "antileakage", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res AntiLeakageRule
	return &res, extract.Into(raw.Body, &res)
}
