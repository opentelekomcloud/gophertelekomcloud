package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateAntiCrawlerOpts struct {
	// URL to which the rule applies.
	Url string `json:"url" required:"true"`
	// Rule matching logic
	// 1: Include
	// 2: Not include
	// 3: Equal
	// 4: Not equal
	// 5: Prefix is
	// 6: Prefix is not
	// 7: Suffix is
	// 8: Suffix is not
	Logic int `json:"logic" required:"true"`
	// Rule name.
	Name string `json:"name" required:"true"`
}

// UpdateAntiCrawler is used to update a JavaScript anti-crawler rule.
func UpdateAntiCrawler(client *golangsdk.ServiceClient, policyId, ruleId string, opts UpdateAntiCrawlerOpts) (*AntiCrawlerRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/waf/policy/{policy_id}/anticrawler/{rule_id}
	raw, err := client.Put(client.ServiceURL("waf", "policy", policyId, "anticrawler", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}
	var res AntiCrawlerRule
	return &res, extract.Into(raw.Body, &res)
}
