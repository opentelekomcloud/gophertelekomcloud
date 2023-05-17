package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Instance, error) {
	// GET /v1/{project_id}/premium-waf/instance
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	url := client.ServiceURL("instance") + query.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Instance
	err = extract.IntoSlicePtr(raw.Body, &res, "items")
	return res, err
}

type ListOpts struct {
	// Fuzzy query for dedicated WAF engine names.
	// Only Prefix and Suffix match query are supported
	Name string `json:"instancename"`
}
