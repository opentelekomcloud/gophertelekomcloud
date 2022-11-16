package resources

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListAllResourcesOpts struct {
	// Specifies the region ID.
	// Maximum length: 36
	RegionId string `q:"region_id,omitempty"`
	// Specifies the resource type.
	// Maximum length: 40
	Type string `q:"type,omitempty"`
	// Specifies the maximum number of records to return.
	// Minimum value: 1
	// Maximum value: 200
	Limit int32 `q:"limit,omitempty"`
	// Specifies the pagination parameter.
	// You can use the marker value returned in the previous request as the number of the first page of records to return in this request.
	// Minimum length: 4
	// Maximum length: 400
	Marker int64 `q:"marker,omitempty"`
}

func ListAllResources(client *golangsdk.ServiceClient, opts ListAllResourcesOpts) (*ListAllResourcesResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/resource-manager/domains/{domain_id}/all-resources
	raw, err := client.Get(client.ServiceURL("all-resources")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListAllResourcesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListAllResourcesResponse struct {
	// Specifies the resource list
	Resources []ResourceEntity `json:"resources,omitempty"`
	// Specifies the pagination object.
	PageInfo PageInfo `json:"page_info,omitempty"`
}
