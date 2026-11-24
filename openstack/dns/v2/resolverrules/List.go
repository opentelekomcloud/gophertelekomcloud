package resolverrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts contains the query parameters of the List function.
type ListOpts struct {
	// Domain name of the endpoint rule to be queried.
	DomainName string `q:"domain_name,omitempty"`
	// Name of the endpoint rule to be queried.
	Name string `q:"name,omitempty"`
	// Endpoint ID.
	EndpointID string `q:"endpoint_id,omitempty"`
	// ID of an endpoint rule.
	ID string `q:"id,omitempty"`
	// Number of resources on each page. The value ranges from 0 to 500. The default value is 500.
	Limit int `q:"limit,omitempty"`
	// Start offset of the pagination query. The value ranges from 0 to 2147483647. The default value is 0.
	// If marker is not left blank, the query starts from the resource specified by marker.
	Offset int `q:"offset,omitempty"`
	// Start resource ID of pagination query.
	// If the parameter is left blank, only resources on the first page are queried.
	Marker string `q:"marker,omitempty"`
}

// List is used to query endpoint rules.
// Send GET /v2.1/resolverrules
func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("resolverrules").WithQueryParams(&opts).Build()
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
	// List of endpoint rules.
	ResolverRules []ResolverRule `json:"resolver_rules"`
	// Link of the current resource or other related resources.
	// When a response is broken into pages, a next link is provided to retrieve all results.
	Links PageLink `json:"links"`
	// Number of resources that meet the query condition.
	Metadata Metadata `json:"metadata"`
	// Pagination information.
	PageInfo PageInfo `json:"page_info"`
}

// PageLink contains pagination links.
type PageLink struct {
	// Link to the current resource.
	Self string `json:"self"`
	// Link to the next page.
	Next string `json:"next"`
}

// Metadata contains the number of resources that meet the query condition.
type Metadata struct {
	// Number of resources that meet the filter criteria. The number is irrelevant to limit or offset.
	TotalCount int `json:"total_count"`
}

// PageInfo contains pagination information.
type PageInfo struct {
	// Marker of the next page.
	NextMarker string `json:"next_marker"`
	// Marker of the previous page.
	PreviousMarker string `json:"previous_marker"`
	// Page size.
	CurrentCount int `json:"current_count"`
}
