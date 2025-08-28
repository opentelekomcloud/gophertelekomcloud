package snatrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListSnatRulesQueryParams struct {
	// Specifies the number of records displayed on each page.
	// The value ranges from 1 to 2000. Default value: 2000
	Limit int `q:"limit"`
	// Specifies the start resource ID of pagination query. If the parameter is left blank, only resources on the first page are queried.
	// The value is obtained from next_marker or previous_marker in PageInfo queried last time.
	Marker string `q:"marker"`
	// Specifies whether to query resources on the previous page.
	PageReverse bool `q:"page_reverse"`
	// Specifies the SNAT rule ID.
	Id []string `q:"id"`
	// Project ID
	ProjectId []string `q:"project_id"`
	// Provides supplementary information about the private NAT gateway.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description []string `q:"description"`
	// Specifies the private NAT gateway ID.
	GatewayId []string `q:"gateway_id"`
	// Specifies the CIDR block that matches the SNAT rule.
	Cidr []string `q:"cidr"`
	// Specifies the ID of the subnet that matches the SNAT rule.
	VirSubnetId []string `q:"virsubnet_id"`
	// Specifies the ID of the transit IP address.
	TransitIpId []string `q:"transit_ip_id"`
	// Specifies the transit IP address.
	TransitIpAddress []string `q:"transit_ip_address"`
	// Specifies the ID of the enterprise project that is associated with the SNAT rule when the SNAT rule is created.
	EnterpriseProjectId []string `q:"enterprise_project_id"`
}

// This function is used to query SNAT rules.
func List(client *golangsdk.ServiceClient, opts ListSnatRulesQueryParams) (*ListResponse, error) {
	// GET /v3/{project_id}/private-nat/snat-rules
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("private-nat", "snat-rules").
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
	// Specifies the response body for querying SNAT rules.
	SnatRules []PrivateSnat `json:"snat_rules"`
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
