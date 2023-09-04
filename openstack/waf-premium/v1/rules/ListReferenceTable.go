package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListReferenceTableOpts struct {
	// Number of records on each page.
	// The maximum value is 100. If this parameter is not specified, the default value -1 is used.
	// All policies are queried regardless of the value of Page
	PageSize int64 `q:"pagesize,omitempty"`
	// Page. Default value: 1
	Page int `q:"page,omitempty"`
	// Reference table name, Fuzzy search is supported.
	Name string `q:"name,omitempty"`
}

// ListReferenceTable is used to query the reference table list.
func ListReferenceTable(client *golangsdk.ServiceClient, opts ListReferenceTableOpts) ([]ReferenceTable, error) {
	query, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/{project_id}/waf/valuelist
	url := client.ServiceURL("waf", "valuelist") + query.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []ReferenceTable
	err = extract.IntoSlicePtr(raw.Body, &res, "items")
	return res, err
}
