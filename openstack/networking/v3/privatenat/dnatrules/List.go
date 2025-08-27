package dnatrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListDnatRulesQueryParams struct {
	// Specifies the number of records displayed on each page.
	// The value ranges from 1 to 2000. Default value: 2000
	Limit int `q:"limit"`
	// Specifies the start resource ID of pagination query. If the parameter is left blank, only resources on the first page are queried.
	// The value is obtained from next_marker or previous_marker in PageInfo queried last time.
	Marker string `q:"marker"`
	// Specifies whether to query resources on the previous page.
	PageReverse bool `q:"page_reverse"`
	// Specifies the DNAT rule ID.
	Id []string `q:"id"`
	// Project ID
	ProjectId []string `q:"project_id"`
	// Specifies the ID of the enterprise project that is associated with the DNAT rule when the DNAT rule is created.
	EnterpriseProjectId []string `q:"enterprise_project_id"`
	// Provides supplementary information about the private NAT gateway.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description []string `q:"description"`
	// Specifies the private NAT gateway ID.
	GatewayId []string `q:"gateway_id"`
	// Specifies the ID of the transit IP address.
	TransitIpId []string `q:"transit_ip_id"`
	// Specifies the transit IP address.
	ExternalIpAddress []string `q:"external_ip_address"`
	// Specifies the port ID of the resource that the NAT gateway is bound to. The resource can be a compute instance, load balancer (v2 or v3), or virtual IP address.
	NetworkInterfaceId []string `q:"network_interface_id"`
	// Specifies the backend resource type of the DNAT rule.
	// The type can be:
	// COMPUTE: The backend resource is a compute instance.
	// VIP: The backend resource is a virtual IP address.
	// ELB: The backend resource is a v2 load balancer.
	// ELBv3: The backend resource is a v3 load balancer.
	// CUSTOMIZE: The backend resource is a user-defined IP address.
	Type []string `q:"type"`
	// Specifies the port IP address that the NAT gateway uses. The resource can be a compute instance, load balancer (v2 or v3), or virtual IP address.
	PrivateIpAddress []string `q:"private_ip_address"`
	// Specifies the DNAT rule protocol type.
	// TCP, UDP, and ANY are supported.
	// The protocol number of TCP, UDP, and ANY are 6, 17, and 0, respectively.
	Protocol []string `q:"protocol"`
	// Specifies the port number of the resource, which can be a compute instance, load balancer (v2 or v3), or virtual IP address.
	InternalServicePort []string `q:"internal_service_port"`
	// Specifies the port number of the transit IP address.
	TransitServicePort []string `q:"transit_service_port"`
	// Specifies the time when the DNAT rule was created. It is a UTC time in yyyy-mm-ddThh:mm:ssZ format.
	CreatedAt string `q:"created_at"`
	// Specifies the time when the DNAT rule was updated. It is a UTC time in yyyy-mm-ddThh:mm:ssZ format.
	UpdatedAt string `q:"updated_at"`
}

// This function is used to query DNAT rules.
func List(client *golangsdk.ServiceClient, opts ListDnatRulesQueryParams) (*ListResponse, error) {
	// GET /v3/{project_id}/private-nat/dnat-rules
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("private-nat", "dnat-rules").
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
	// Specifies the response body for querying DNAT rules.
	DnatRules []PrivateDnat `json:"dnat_rules"`
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
