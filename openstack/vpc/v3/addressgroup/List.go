package addressgroup

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
	// Specifies the IP address group ID.
	Id []string `q:"id"`
	// Specifies the IP address group name.
	Name []string `q:"name"`
	// IP address version of an IP address group.
	// 4: IPv4 address group
	// 6: IPv6 address group
	IpVersion int `q:"ip_version"`
	// Specifies the Description about an IP address group.
	// This field can be used to precisely filter IP address groups. Multiple descriptions can be specified for filtering.
	Description []string `q:"description"`
}

// This function is used to query all IP address groups of a tenant.
func List(client *golangsdk.ServiceClient, opts ListQueryParams) (*ListResponse, error) {
	// GET /v3/{project_id}/vpc/address-groups
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("address-groups").
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
	// Specifies the response body for querying IP address groups.
	AddressGroups []AddressGroup `json:"address_groups"`
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
