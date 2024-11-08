package query

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListServicesOpts struct {
	// Specifies the number of records returned on each page during pagination query.
	Limit *int `q:"limit"`
	// Specifies the pagination offset.
	Offset *int `q:"offset"`
	// Specifies whether resources are collected by default. tracked indicates that resources are collected by default,
	// and untracked indicates that resources are not collected by default.
	Track string `q:"track"`
}

func ListServices(client *golangsdk.ServiceClient, domainId string, opts ListServicesOpts) ([]Service, error) {
	// GET /v1/resource-manager/domains/{domain_id}/providers
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", domainId, "providers").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ServicesPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractServices(pages)
}

type ServicesPage struct {
	pagination.NewSinglePageBase
}

func ExtractServices(r pagination.NewPage) ([]Service, error) {
	var s struct {
		Services []Service `json:"resource_providers"`
	}
	err := extract.Into(bytes.NewReader((r.(ServicesPage)).Body), &s)
	return s.Services, err
}

type Service struct {
	// Specifies the cloud service name.
	ServiceName string `json:"provider"`
	// Specifies the display name of the cloud service.
	// You can set the language by configuring X-Language in the request header.
	DisplayName string `json:"display_name"`
	// Specifies the display name of the cloud service type.
	// You can set the language by configuring X-Language in the request header.
	CategoryDisplayName string `json:"category_display_name"`
	// Specifies the resource type list.
	ResourceTypes []ResourceTypes `json:"resource_types"`
}

type ResourceTypes struct {
	// Specifies the resource type.
	Name string `json:"name"`
	// Specifies the display name of the resource type.
	DisplayName string `json:"display_name"`
	// Specifies whether the resource is a global resource.
	Global bool `json:"global"`
	// Specifies the list of supported regions.
	Regions []string `json:"regions"`
	// Specifies the console endpoint ID.
	ConsoleEndpointId string `json:"console_endpoint_id"`
	// Specifies the URL of the resource list page on the console.
	ConsoleListUrl string `json:"console_list_url"`
	// Specifies the URL of the resource details page on the console.
	ConsoleDetailUrl string `json:"console_detail_url"`
	// Specifies whether resources are collected by default.
	// tracked indicates that resources are collected by default,
	// and untracked indicates that resources are not collected by default.
	Track string `json:"track"`
}
