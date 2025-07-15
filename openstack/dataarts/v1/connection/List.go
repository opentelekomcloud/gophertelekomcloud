package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Offset         int    `q:"offset"`
	Limit          int    `q:"limit"`
	ConnectionName string `q:"connectionName"`
}

type ListResp struct {
	// Total the number of connections.
	Total int `q:"total"`
	// Connection the Connection list.
	Connections []Connection `q:"connections"`
}

// List is used to query a connection list.
// Send request GET /v1/{project_id}/connections?offset={offset}&limit={limit}&connectionName={connectionName}
func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(connectionsEndpoint).WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResp
	err = extract.IntoSlicePtr(raw.Body, &res, "")
	return &res, err
}
