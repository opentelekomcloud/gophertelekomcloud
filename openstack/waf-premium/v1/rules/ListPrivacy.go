package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListPrivacyOpts struct {
	// Number of records on each page.
	// The maximum value is 100. If this parameter is not specified, the default value -1 is used.
	// All policies are queried regardless of the value of Page
	PageSize int64 `q:"pagesize,omitempty"`
	// Page. Default value: 1
	Page int `q:"page,omitempty"`
}

// ListPrivacy is used to query the data masking rule list.
func ListPrivacy(client *golangsdk.ServiceClient, policyId string, opts ListPrivacyOpts) ([]PrivacyRule, error) {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/{project_id}/waf/policy/{policy_id}/privacy
	url := client.ServiceURL("waf", "policy", policyId, "privacy") + query.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []PrivacyRule
	err = extract.IntoSlicePtr(raw.Body, &res, "items")
	return res, err
}
