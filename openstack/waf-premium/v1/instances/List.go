package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Instance, error) {
	// GET /v1/{project_id}/premium-waf/instance
	var opts2 interface{} = opts
	query, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	url := client.ServiceURL("premium-waf", "instance") + query.String()
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
