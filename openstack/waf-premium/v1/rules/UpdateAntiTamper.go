package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateAntiTamperCache is used to updating the cache for a web tamper protection Rule.
func UpdateAntiTamperCache(client *golangsdk.ServiceClient, policyId, ruleId string) (*AntiTamperRule, error) {
	// POST /v1/{project_id}/waf/policy/{policy_id}/antitamper/{rule_id}/refresh
	raw, err := client.Post(client.ServiceURL("waf", "policy", policyId, "antitamper", ruleId, "refresh"), nil, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}
	var res AntiTamperRule
	return &res, extract.Into(raw.Body, &res)
}
