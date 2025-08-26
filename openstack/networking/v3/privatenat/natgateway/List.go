package natgateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListGatewaysQueryParams struct {
	// Specifies the number of records displayed on each page.
	// The value ranges from 1 to 2000. Default value: 2000
	Limit int `q:"limit"`
	// Specifies the start resource ID of pagination query. If the parameter is left blank, only resources on the first page are queried.
	// The value is obtained from next_marker or previous_marker in PageInfo queried last time.
	Marker string `q:"marker"`
	// Specifies whether to query resources on the previous page.
	PageReverse bool `q:"page_reverse"`
	// Specifies the private NAT gateway ID.
	Id []string `q:"id"`
	// Specifies the private NAT gateway name.
	// Only digits, letters, underscores (_), and hyphens (-) are allowed.
	Name []string `q:"name"`
	// Provides supplementary information about the private NAT gateway.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description []string `q:"description"`
	// Specifies the private NAT gateway specifications.
	// The value can be: Small, Medium, Large, Extra-large.
	// Default value: Small.
	Spec []string `q:"spec"`
	// Project ID
	ProjectId []string `q:"project_id"`
	// Specifies the private NAT gateway status.
	// The value can be:
	// ACTIVE: The private NAT gateway is running properly.
	// FROZEN: The private NAT gateway is frozen.
	Status []string `q:"status"`
	// VPC ID
	VpcId []string `q:"vpc_id"`
	// Specifies the ID of the subnet where the private NAT gateway works.
	VirSubnetID []string `q:"virsubnet_id"`
	// Specifies the ID of the enterprise project that is associated with the private NAT gateway when the private NAT gateway is created.
	EnterpriseProjectId []string `q:"enterprise_project_id"`
}

// This function is used to query private NAT gateways.
func List(client *golangsdk.ServiceClient, opts ListGatewaysQueryParams) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("private-nat", "gateways").
		WithQueryParams(opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v3/{project_id}/private-nat/gateways
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	// Specifies the response body for the private NAT gateways.
	Gateways []PrivateNATGateway `json:"gateways"`
	// Request ID
	RequestID string `json:"request_id"`
	// Specifies the pagination information.
	PageInfo PageInfo `json:"page_info"`
}

type PageInfo struct {
	// Specifies the ID of the last record in this query, which can be used in the next query.
	NextMarker string `json:"next_marker"`
	// Specifies the ID of the first record in the pagination query result.
	// When page_reverse is set to true, this parameter is used together to query resources on the previous page.
	PreviousMarker string `json:"previous_marker"`
	// Specifies the ID of the last record in the pagination query result. It is usually used to query resources on the next page. Value range: 1-200
	CurrentCount int `json:"current_count"`
}
