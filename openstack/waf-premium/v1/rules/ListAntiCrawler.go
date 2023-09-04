package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListAntiCrawlerOpts struct {
	// Number of records on each page.
	// The maximum value is 100. If this parameter is not specified, the default value -1 is used.
	// All policies are queried regardless of the value of Page
	PageSize int64 `q:"pagesize,omitempty"`
	// Page. Default value: 1
	Page int `q:"page,omitempty"`
	// JavaScript anti-crawler rule protection mode.
	// anticrawler_except_url: In this mode, all paths are protected except the one specified in the queried anti-crawler rule.
	// anticrawler_specific_url: In this mode, the path specified in the queried rule is protected.
	Type string `q:"type,omitempty"`
}

// ListAntiCrawlers is used to query the list of JavaScript anti-crawler rules.
func ListAntiCrawlers(client *golangsdk.ServiceClient, policyId string, opts ListAntiCrawlerOpts) ([]AntiCrawlerRule, error) {
	query, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/{project_id}/waf/policy/{policy_id}/anticrawler
	url := client.ServiceURL("waf", "policy", policyId, "anticrawler") + query.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []AntiCrawlerRule
	err = extract.IntoSlicePtr(raw.Body, &res, "items")
	return res, err
}
