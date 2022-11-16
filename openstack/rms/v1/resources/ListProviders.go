package resources

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListProvidersOpts struct {
	// Specifies the pagination offset.
	// Minimum value: 1
	// Maximum value: 1000
	Offset *int32 `q:"offset,omitempty"`
	// Specifies the maximum
	// number of records to return.
	// Minimum value: 1
	// Maximum value: 200
	Limit *int32 `q:"limit,omitempty"`
}

func ListProviders(client *golangsdk.ServiceClient, opts ListProvidersOpts) (*ListProvidersResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/resource-manager/domains/{domain_id}/providers
	raw, err := client.Get(client.ServiceURL("providers")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListProvidersResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListProvidersResponse struct {
	// Specifies the list of cloud service details.
	ResourceProviders *[]ResourceProviderResponse `json:"resource_providers,omitempty"`
	// Specifies the total number of cloud services supported by RMS.
	TotalCount *int32 `json:"total_count,omitempty"`
}

type ResourceProviderResponse struct {
	// Specifies the cloud service name. For details, see Supported Resource.
	Provider string `json:"provider,omitempty"`
	// Specifies the display name of the cloud service.
	// You can set the language by configuring XLanguage in the request header
	DisplayName string `json:"display_name,omitempty"`
	// Specifies the display name of the cloud service category.
	// You can set the language by configuring X-Language in the request header.
	// Currently supported categories: Computing, Network, Storage, Database, Security, EI Enterprise.
	CategoryDisplayName string `json:"category_display_name,omitempty"`
	// Specifies the resource type list.
	ResourceTypes []ResourceTypeResponse `json:"resource_types,omitempty"`
}

type ResourceTypeResponse struct {
	// Specifies the resource type.
	Name string `json:"name,omitempty"`
	// Specifies the display name of the resource type.
	// You can set the language by configuring X-Language in the request header.
	DisplayName string `json:"display_name,omitempty"`
	// Specifies whether the resource is a global resource.
	Global *bool `json:"global,omitempty"`
	// Specifies the list of supported regions.
	// Array of Strings.
	Regions []string `json:"regions,omitempty"`
	// Specifies the endpoint ID of the console.
	ConsoleEndpointId string `json:"console_endpoint_id,omitempty"`
	// Specifies the URL of the resource list page.
	ConsoleListUrl string `json:"console_list_url,omitempty"`
	// Specifies the URL of the resource details page.
	ConsoleDetailUrl string `json:"console_detail_url,omitempty"`
}
