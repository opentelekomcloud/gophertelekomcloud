package resourcegroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListResourcesOpts contains the options for querying resources in a resource group.
type ListResourcesOpts struct {
	// Specifies the resource dimension.
	// Multiple dimensions are separated by commas and sorted alphabetically.
	DimName string `q:"dim_name"`
	// Specifies the resource dimension value.
	// Fuzzy matching is not supported.
	// It can contain 1 to 256 characters.
	DimValue string `q:"dim_value"`
	// Specifies the resource health status.
	// Possible values: health, unhealthy, no_alarm_rule
	Status string `q:"status"`
	// Specifies the number of records on each page.
	// Default: 100, Range: 1-1000
	Limit int `q:"limit,omitempty"`
	// Specifies the pagination offset.
	// Default: 0, Range: 0-10000
	Offset int `q:"offset,omitempty"`
	// Specifies the enterprise project ID.
	ExtendRelationId string `q:"extend_relation_id"`
}

// ListResourcesResponse contains the response from the ListResources request.
type ListResourcesResponse struct {
	// Specifies the total number of resources.
	Count int `json:"count"`
	// Specifies the list of resources.
	Resources []ResourceInfo `json:"resources"`
}

// ResourceInfo represents a resource in the list response.
type ResourceInfo struct {
	// Specifies the resource health status.
	// Possible values: health, unhealthy, no_alarm_rule
	Status string `json:"status"`
	// Specifies the resource dimension information.
	Dimensions []ResourceDimension `json:"dimensions"`
}

// ListResources returns a list of resources for a specified service type in a resource group.
func ListResources(client *golangsdk.ServiceClient, groupId string, service string, opts ListResourcesOpts) (*ListResourcesResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-groups", groupId, "services", service, "resources").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/resource-groups/{group_id}/services/{service}/resources
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ListResourcesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
