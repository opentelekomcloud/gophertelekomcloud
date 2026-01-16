package resourcegroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts contains the options for querying resource groups.
type ListOpts struct {
	// Specifies the enterprise project ID.
	// The value can be a UUID or "0".
	EnterpriseProjectId string `q:"enterprise_project_id"`
	// Specifies the resource group name. Fuzzy matching is supported.
	// It can contain 1 to 128 characters.
	GroupName string `q:"group_name"`
	// Specifies the resource group ID.
	// It contains 24 characters and starts with "rg".
	GroupId string `q:"group_id"`
	// Specifies the pagination offset.
	// Default: 0, Range: 0-10000
	Offset int `q:"offset,omitempty"`
	// Specifies the number of records on each page.
	// Default: 100, Range: 1-100
	Limit int `q:"limit,omitempty"`
	// Specifies how resources are added to the resource group.
	// Possible values: EPS, TAG, Manual
	Type string `q:"type"`
}

// ListResponse contains the response from the List request.
type ListResponse struct {
	// Specifies the total number of resource groups.
	Count int `json:"count"`
	// Specifies the list of resource groups.
	ResourceGroups []ResourceGroup `json:"resource_groups"`
}

// ResourceGroup represents a resource group in the list response.
type ResourceGroup struct {
	// Specifies the resource group name.
	GroupName string `json:"group_name"`
	// Specifies the resource group ID.
	GroupId string `json:"group_id"`
	// Specifies the time when the resource group was created.
	// The value is in UTC format: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	CreateTime string `json:"create_time"`
	// Specifies the enterprise project ID.
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// Specifies how resources are added to the resource group.
	// Possible values: EPS, TAG, Manual
	Type string `json:"type"`
}

// List returns a list of resource groups.
func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("resource-groups").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/resource-groups
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
