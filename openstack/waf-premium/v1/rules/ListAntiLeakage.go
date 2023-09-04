package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListAntiLeakageOpts struct {
	// Number of records on each page.
	// The maximum value is 100. If this parameter is not specified, the default value -1 is used.
	// All policies are queried regardless of the value of Page
	PageSize int64 `q:"pagesize,omitempty"`
	// Page. Default value: 1
	Page int `q:"page,omitempty"`
}

// ListAntiLeakage is used to query the list of information leakage prevention rules.
func ListAntiLeakage(client *golangsdk.ServiceClient, policyId string, opts ListAntiLeakageOpts) ([]AntiLeakageRule, error) {
	query, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/{project_id}/waf/policy/{policy_id}/antileakage
	url := client.ServiceURL("waf", "policy", policyId, "antileakage") + query.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []AntiLeakageRule
	err = extract.IntoSlicePtr(raw.Body, &res, "items")
	return res, err
}
