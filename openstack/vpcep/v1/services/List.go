package services

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Specifies the name of the VPC endpoint service. The value is not case-sensitive and supports fuzzy match.
	Name string `q:"endpoint_service_name"`
	// Specifies the unique ID of the VPC endpoint service.
	ID string `q:"id"`
	// Specifies the status of the VPC endpoint service.
	//    creating: indicates the VPC endpoint service is being created.
	//    available: indicates the VPC endpoint service is connectable.
	//    failed: indicates the creation of the VPC endpoint service failed.
	//    deleting: indicates the VPC endpoint service is being deleted.
	Status Status `q:"status"`
	// Specifies the sorting field of the VPC endpoint service list.
	SortKey string `q:"sort_key"`
	// Specifies the sorting method of the VPC endpoint service list.
	SortDir string `q:"sort_dir"`
	// Specifies the maximum number of VPC endpoint services displayed on each page.
	Limit *int `q:"limit"`
	// Specifies the offset.
	Offset *int `q:"offset"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Service, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("vpc-endpoint-services").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ServicePage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractServices(pages)
}

type ServicePage struct {
	pagination.NewSinglePageBase
}

func ExtractServices(r pagination.NewPage) ([]Service, error) {
	var s struct {
		Services []Service `json:"endpoint_services"`
	}
	err := extract.Into(bytes.NewReader((r.(ServicePage)).Body), &s)
	return s.Services, err
}
