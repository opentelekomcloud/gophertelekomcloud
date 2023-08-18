package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateAntiCrawlerOpts struct {
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
	// JavaScript anti-crawler rule type.
	// anticrawler_specific_url: used to protect a specific path specified by the rule.
	// anticrawler_except_url: used to protect all paths except the one specified by the rule.
	Type string `json:"type" required:"true"`
}

// CreateAntiCrawler will create a JavaScript anti-crawler rule  on the values in CreateOpts.
func CreateAntiCrawler(client *golangsdk.ServiceClient, policyId string, opts CreateAntiCrawlerOpts) (*AntiCrawlerRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/waf/policy/{policy_id}/anticrawler
	raw, err := client.Post(client.ServiceURL("waf", "policy", policyId, "anticrawler"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes:     []int{200},
			MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		})
	if err != nil {
		return nil, err
	}

	var res AntiCrawlerRule
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type AntiCrawlerRule struct {
	// Rule ID.
	ID string `json:"id"`
	// Policy ID.
	PolicyId string `json:"policyid"`
	// Timestamp the rule is created.
	CreatedAt int64 `json:"timestamp"`
	// URL to which the rule applies.
	Url string `json:"url"`
	// Rule matching logic
	// 1: Include
	// 2: Not include
	// 3: Equal
	// 4: Not equal
	// 5: Prefix is
	// 6: Prefix is not
	// 7: Suffix is
	// 8: Suffix is not
	Logic int `json:"logic"`
	// Rule name.
	Name string `json:"name"`
	// JavaScript anti-crawler rule type.
	// anticrawler_specific_url: used to protect a specific path specified by the rule.
	// anticrawler_except_url: used to protect all paths except the one specified by the rule.
	Type string `json:"type"`
	// Rule status. The value can be 0 or 1.
	Status int `json:"status"`
}
