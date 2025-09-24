package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListQueryParams struct {
	// Specifies the number of records displayed on each page.
	// The value ranges from 1 to 2000.
	Limit int `q:"limit"`
	// Specifies the start resource ID of pagination query. If the parameter is left blank, only resources on the first page are queried.
	// The value is obtained from next_marker or previous_marker in PageInfo queried last time.
	Marker string `q:"marker"`
	// Specifies the Security group rule ID.
	Id []string `q:"id"`
	// Specifies the ID of the security group to which the security group rule belongs. Multiple IDs can be specified for filtering.
	SecurityGroupId []string `q:"security_group_id"`
	// Protocol specified in the security group rule. Multiple protocols can be specified for filtering.
	Protocol []string `q:"protocol"`
	// Specifies the supplementary information about the security group rule.
	// This field can be used to filter security groups. Multiple descriptions can be specified for filtering.
	Description []string `q:"description"`
	// Specifies the ID of the remote security group. Multiple IDs can be specified for filtering.
	RemoteGroupId []string `q:"remote_group_id"`
	// Inbound or outbound direction of a security group rule.
	// The value can be: ingress or egress
	Direction string `q:"direction"`
	// Action of the security group rule.
	Action string `q:"action"`
	// Remote IP address or CIDR block.
	RemoteIpPrefix string `q:"remote_ip_prefix"`
	// Security group rule priority. Multiple priorities can be specified for filtering.
	Priority []int `q:"priority"`
	// Type of the security group rule. Value range: IPv4, ipv4, IPv6, or ipv6.
	EtherType []string `q:"ethertype"`
	// Project ID
	ProjectId []string `q:"project_id"`
	// ID of the remote IP address group. Multiple IDs can be specified for filtering.
	RemoteAddressGroupId []string `q:"remote_address_group_id"`
}

// This function is used to query all security group rules.
func List(client *golangsdk.ServiceClient, opts ListQueryParams) (*ListResponse, error) {
	// GET /v3/{project_id}/vpc/security-group-rules
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("security-group-rules").
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
	// Specifies the response body for querying security group rules.
	SecurityGroupRules []SecurityGroupRule `json:"security_group_rules"`
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
