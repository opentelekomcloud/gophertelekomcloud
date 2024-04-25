package resource

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Offset       int    `q:"offset"`
	Limit        int    `q:"limit"`
	ResourceName string `q:"resourceName"`
}

// List is used to query a resource list. During the query, you can specify the page number and the maximum number of records on each page.
// Send request GET /v1/{project_id}/resources?offset={offset}&limit={limit}&resourceName={resourceName}
func List(client *golangsdk.ServiceClient, opts *ListOpts, workspace string) (*ListResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(resourcesEndpoint).WithQueryParams(&opts).Build()
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

	var res *ListResp
	err = extract.Into(raw.Body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type ListResp struct {
	// Total is the total number of resources.
	Total int `json:"total"`
	// Resources is a list of resources.
	Resources []*Resource `json:"resources"`
}
