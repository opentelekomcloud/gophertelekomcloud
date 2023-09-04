package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListCustomOpts struct {
	// Number of records on each page.
	// The maximum value is 100. If this parameter is not specified, the default value -1 is used.
	// All policies are queried regardless of the value of Page
	PageSize int64 `q:"pagesize,omitempty"`
	// Page. Default value: 1
	Page int `q:"page,omitempty"`
}

// ListCustoms is used to query the list of CC attack protection rules.
func ListCustoms(client *golangsdk.ServiceClient, policyId string, opts ListCustomOpts) ([]CustomRule, error) {
	query, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/{project_id}/waf/policy/{policy_id}/custom
	url := client.ServiceURL("waf", "policy", policyId, "custom") + query.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []CustomRule
	err = extract.IntoSlicePtr(raw.Body, &res, "items")
	return res, err
}
