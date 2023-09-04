package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetAntiCrawler is used to query a JavaScript anti-crawler rule by ID.
func GetAntiCrawler(client *golangsdk.ServiceClient, policyId, ruleId string) (*AntiCrawlerRule, error) {
	// GET /v1/{project_id}/waf/policy/{policy_id}/anticrawler/{rule_id}
	raw, err := client.Get(client.ServiceURL("waf", "policy", policyId, "anticrawler", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res AntiCrawlerRule
	return &res, extract.Into(raw.Body, &res)
}
