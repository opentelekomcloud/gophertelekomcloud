package services

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListPublicOpts struct {
	// Specifies the name of the public VPC endpoint service.
	// The value is not case-sensitive and supports fuzzy match.
	Name string `q:"endpoint_service_name"`
	// Specifies the unique ID of the public VPC endpoint service.
	ID string `q:"id"`
	// Specifies the sorting field of the VPC endpoint service list.
	SortKey string `q:"sort_key"`
	// Specifies the sorting method of the VPC endpoint service list.
	SortDir string `q:"sort_dir"`
	// Specifies the maximum number of VPC endpoint services displayed on each page.
	Limit *int `q:"limit"`
	// Specifies the offset.
	Offset *int `q:"offset"`
}

func ListPublic(client *golangsdk.ServiceClient, opts ListPublicOpts) ([]PublicService, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("vpc-endpoint-services", "public").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return PublicServicePage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractPublicServices(pages)
}

type PublicServicePage struct {
	pagination.NewSinglePageBase
}

func ExtractPublicServices(r pagination.NewPage) ([]PublicService, error) {
	var s struct {
		Services []PublicService `json:"endpoint_services"`
	}
	err := extract.Into(bytes.NewReader((r.(PublicServicePage)).Body), &s)
	return s.Services, err
}

func (p ServicePage) IsEmpty() (bool, error) {
	srv, err := ExtractServices(p)
	if err != nil {
		return false, err
	}
	return len(srv) == 0, err
}

type PublicService struct {
	// Specifies the unique ID of the public VPC endpoint service.
	ID string `json:"id"`
	// Specifies the owner of the VPC endpoint service.
	Owner string `json:"owner"`
	// Specifies the name of the public VPC endpoint service.
	ServiceName string `json:"service_name"`
	// Specifies the type of the VPC endpoint service.
	ServiceType ServiceType `json:"service_type"`
	// Specifies the creation time of the VPC endpoint service.
	// The UTC time format is used: YYYY-MM-DDTHH:MM:SSZ.
	CreatedAt string `json:"created_at"`
	// Specifies whether the associated VPC endpoint carries a charge.
	IsCharge bool `json:"is_charge"`
}
