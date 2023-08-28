package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListBlacklistOpts struct {
	// Number of records on each page.
	// The maximum value is 100. If this parameter is not specified, the default value -1 is used.
	// All policies are queried regardless of the value of Page
	PageSize int64 `q:"pagesize,omitempty"`
	// Page. Default value: 1
	Page int `q:"page,omitempty"`
	// Rule name, Fuzzy search is supported.
	Name string `q:"name,omitempty"`
}

// ListBlacklists is used to query the list of blacklist and whitelist rules.
func ListBlacklists(client *golangsdk.ServiceClient, policyId string, opts ListBlacklistOpts) ([]BlacklistRule, error) {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/{project_id}/waf/policy/{policy_id}/whiteblackip
	url := client.ServiceURL("waf", "policy", policyId, "whiteblackip") + query.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []BlacklistRule
	err = extract.IntoSlicePtr(raw.Body, &res, "items")
	return res, err
}
