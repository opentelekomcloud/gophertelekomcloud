package services

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BasicOpts struct {
	// Specifies the name of the VPC endpoint service.
	// Either this parameter or the id parameter must be selected.
	Name string `q:"endpoint_service_name"`
	// Specifies the unique ID of the VPC endpoint service.
	// Either this parameter or the name parameter must be selected.
	ID string `q:"id"`
}

func GetBasicInfo(client *golangsdk.ServiceClient, opts BasicOpts) (*BasicService, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("vpc-endpoint-services", "describe").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	// GET /v1/{project_id}/vpc-endpoint-services/describe?endpoint_service_name={endpoint_service_name}&id={endpoint_service_id}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res BasicService
	return &res, extract.Into(raw.Body, &res)
}

type BasicService struct {
	// Specifies the unique ID of the VPC endpoint service.
	ID string `json:"id"`
	// Specifies the name of the VPC endpoint service.
	ServiceName string `json:"service_name"`
	// Specifies the type of the VPC endpoint service.
	ServiceType string `json:"service_type"`
	// Specifies the resource type.
	ServerType ServerType `json:"server_type"`
	// Specifies the creation time of the VPC endpoint service.
	// The UTC time format is used: YYYY-MM-DDTHH:MM:SSZ.
	CreatedAt string `json:"created_at"`
	// Specifies whether the associated VPC endpoint carries a charge.
	IsCharge bool `json:"is_charge"`
}
