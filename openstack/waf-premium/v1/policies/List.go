package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Policy, error) {
	query, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/{project_id}/waf/policy
	url := client.ServiceURL("waf", "policy") + query.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Policy
	err = extract.IntoSlicePtr(raw.Body, &res, "items")
	return res, err
}

type ListOpts struct {
	// Number of records on each page.
	// The maximum value is 100. If this parameter is not specified, the default value -1 is used.
	// All policies are queried regardless of the value of Page
	PageSize string `q:"pageSize,omitempty"`
	// Page. Default value: 1
	Page string `q:"page,omitempty"`
	// Policy name. Fuzzy search is supported.
	Name string `q:"name,omitempty"`
}

type Certificates struct {
	// Certificate ID.
	ID string `json:"id"`
	// Certificate name.
	Name string `json:"name"`
	// Timestamp when the certificate is uploaded.
	CreatedAt int64 `json:"timestamp"`
}
