package endpoints

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListVPCsOpts contains the query parameters of the ListVPCs function.
type ListVPCsOpts struct {
	// Number of resources on each page. The value ranges from 0 to 500. The default value is 500.
	Limit int `q:"limit,omitempty"`
	// Start offset of the pagination query. The value ranges from 0 to 2147483647. The default value is 0.
	Offset int `q:"offset,omitempty"`
	// ID of the VPC to be queried.
	VpcID string `q:"vpc_id,omitempty"`
}

// ListVPCs is used to query VPCs of an endpoint.
// Send GET /v2.1/vpcs
func ListVPCs(client *golangsdk.ServiceClient, opts ListVPCsOpts) (*ListVPCsResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("vpcs").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(v21URL(client, url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListVPCsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// ListVPCsResponse is the response of the ListVPCs function.
type ListVPCsResponse struct {
	// VPCs of an endpoint.
	VPCs []VPC `json:"vpcs"`
	// Number of resources that meet the query condition.
	Metadata Metadata `json:"metadata"`
}

// VPC is the VpcsData object.
type VPC struct {
	// VPC ID, which is a UUID used to identify the VPC.
	ID string `json:"id"`
	// Number of inbound endpoints in a VPC.
	InboundEndpointCount int `json:"inbound_endpoint_count"`
	// Number of outbound endpoints in a VPC.
	OutboundEndpointCount int `json:"outbound_endpoint_count"`
}
