package resources

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListResourcesOpts struct {
	// Specifies the cloud service name.
	// Maximum length: 20
	Provider string
	// Specifies the resource type.
	// Maximum length: 20
	Type string
	// Specifies the region ID.
	// Maximum length: 36
	RegionId string `q:"region_id,omitempty"`
	// Specifies the enterprise project ID.
	// Maximum length: 36
	EpId string `q:"ep_id,omitempty"`
	// Specifies the resource tag.
	Tag map[string]string `q:"tag,omitempty"`
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

func ListResources(client *golangsdk.ServiceClient, opts ListResourcesOpts) (*ListResourcesResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/resource-manager/domains/{domain_id}/provider/{provider}/type/{type}/resources
	raw, err := client.Get(client.ServiceURL("provider", opts.Provider, "type", opts.Type, "resources")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResourcesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResourcesResponse struct {
	// Specifies the resource list
	Resources []ResourceEntity `json:"resources,omitempty"`
	// Specifies the pagination object.
	PageInfo PageInfo `json:"page_info,omitempty"`
}

type ResourceEntity struct {
	// Specifies the resource ID.
	Id string `json:"id,omitempty"`
	// Specifies the resource name.
	Name string `json:"name,omitempty"`
	// Specifies the cloud service name. For details, see Supported Resource.
	Provider string `json:"provider,omitempty"`
	// Specifies the resource type.
	Type string `json:"type,omitempty"`
	// Specifies the region ID.
	RegionId string `json:"region_id,omitempty"`
	// Specifies the project ID in OpenStack.
	ProjectId string `json:"project_id,omitempty"`
	// Specifies the project name in OpenStack.
	ProjectName string `json:"project_name,omitempty"`
	// Specifies the enterprise project ID.
	EpId string `json:"ep_id,omitempty"`
	// Specifies the enterprise project name.
	EpName string `json:"ep_name,omitempty"`
	// Specifies the resource checksum.
	Checksum string `json:"checksum,omitempty"`
	// Specifies the time when the resource was created.
	Created string `json:"created,omitempty"`
	// Specifies the time when the resource was updated.
	Updated string `json:"updated,omitempty"`
	// Specifies the status of the operation that causes the resource change.
	ProvisioningState string `json:"provisioning_state,omitempty"`
	// Specifies the resource tags.
	Tags map[string]string `json:"tags,omitempty"`
	// Specifies the detailed properties of the resource.
	Properties map[string]interface{} `json:"properties,omitempty"`
}

type PageInfo struct {
	// Specifies the resource quantity on the current page.
	// Minimum value: 0
	// Maximum value: 200
	CurrentCount int32 `json:"current_count,omitempty"`
	// Specifies the marker value of the next page.
	// Minimum length: 4
	// Maximum length: 400
	NextMarker string `json:"next_marker,omitempty"`
}
