package endpoints

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts contains the query parameters of the List function.
type ListOpts struct {
	// Endpoint direction. The value can be:
	//   inbound: indicates an inbound endpoint.
	//   outbound: indicates an outbound endpoint.
	Direction string `q:"direction,required"`
	// ID of the VPC to which the endpoint to be queried belongs.
	VpcID string `q:"vpc_id,omitempty"`
	// Endpoint name.
	Name string `q:"name,omitempty"`
	// Number of resources on each page. The value ranges from 0 to 500. The default value is 500.
	Limit int `q:"limit,omitempty"`
	// Start offset of the pagination query. The value ranges from 0 to 2147483647. The default value is 0.
	Offset int `q:"offset,omitempty"`
}

// List is used to query endpoints.
// Send GET /v2.1/endpoints
func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("endpoints").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(v21URL(client, url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// ListResponse is the response of the List function.
type ListResponse struct {
	// Returned endpoints.
	Endpoints []Endpoint `json:"endpoints"`
	// Number of resources that meet the query condition.
	Metadata Metadata `json:"metadata"`
}

// Metadata contains the number of resources that meet the query condition.
type Metadata struct {
	// Number of resources that meet the filter criteria. The number is irrelevant to limit or offset.
	TotalCount int `json:"total_count"`
}
