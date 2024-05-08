package script

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Offset     int    `q:"offset"`
	Limit      int    `q:"limit"`
	ScriptName string `q:"scriptName"`
}

// List is used to query the script list. A maximum of 1000 scripts can be returned for each query.
// Send request GET /v1/{project_id}/scripts?offset={offset}&limit={limit}&scriptName={scriptName}
func List(client *golangsdk.ServiceClient, opts ListOpts, workspace string) (*ListResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(scriptsEndpoint).WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	var reqOpts *golangsdk.RequestOpts
	if workspace != "" {
		reqOpts.MoreHeaders = map[string]string{HeaderWorkspace: workspace}
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, reqOpts)
	if err != nil {
		return nil, err
	}

	var res ListResp
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type ListResp struct {
	// Total is the total number of scripts.
	Total int `json:"total"`
	// Scripts is a list of scripts.
	Scripts []*Script `json:"scripts"`
}
