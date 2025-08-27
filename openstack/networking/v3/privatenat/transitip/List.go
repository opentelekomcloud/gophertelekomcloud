package transitip

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListTransitIpsQueryParams struct {
	// Specifies the number of records displayed on each page.
	// The value ranges from 1 to 2000. Default value: 2000
	Limit int `q:"limit"`
	// Specifies the start resource ID of pagination query. If the parameter is left blank, only resources on the first page are queried.
	// The value is obtained from next_marker or previous_marker in PageInfo queried last time.
	Marker string `q:"marker"`
	// Specifies whether to query resources on the previous page.
	PageReverse bool `q:"page_reverse"`
	// Specifies the ID of the transit IP address.
	Id []string `q:"id"`
	// Project ID
	ProjectId []string `q:"project_id"`
	// Specifies the network interface ID of the transit IP address.
	NetworkInterfaceId []string `q:"network_interface_id"`
	// Specifies the transit IP address.
	IpAddress []string `q:"ip_address"`
	// Specifies the ID of the private NAT gateway associated with the transit IP address.
	GatewayId []string `q:"gateway_id"`
	// Specifies the ID of the enterprise project that is associated with the transit IP address when the transit IP address is assigned.
	EnterpriseProjectId []string `q:"enterprise_project_id"`
	// Specifies the subnet ID of the current tenant.
	VirSubnetID []string `q:"virsubnet_id"`
	// Specifies the transit subnet ID.
	TransitSubnetId []string `q:"transit_subnet_id"`
	// Provides supplementary information about the transit IP address.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description []string `q:"description"`
}

// This function is used to query transit IP addresses.
func List(client *golangsdk.ServiceClient, opts ListResponse) (*ListResponse, error) {
	// GET /v3/{project_id}/private-nat/transit-ips
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("private-nat", "transit-ips").
		WithQueryParams(opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	// Specifies the response body for querying transit IP addresses.
	TransitIps []TransitIP `json:"transit_ips"`
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
