package schemas

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type QueryAvailableDbOpts struct {
	// Specifies the Index offset.
	// The query starts from the next piece of data indexed by this parameter. The value is 0 by default.
	// The value must be a positive integer.
	Offset int `q:"offset"`
	// Maximum records to be queried.
	// Value range: 1 to 128.
	// If the parameter value is not specified, 10 are queried by default.
	Limit int `q:"limit"`
}

// This function is used to query DB instances that can be used for creating a schema.
func QueryAvailableDb(client *golangsdk.ServiceClient, instanceId string, opts QueryAvailableDbOpts) ([]QueryAvailableRdsList, error) {
	// GET /v1/{project_id}/instances/{instance_id}/rds?offset={offset}&limit={limit}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceId, "rds").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return DbInstancesPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractDbInstances(pages)
}

func ExtractDbInstances(r pagination.NewPage) ([]QueryAvailableRdsList, error) {
	var queryAvailableDbResponse QueryAvailableDbResponse
	err := extract.Into(bytes.NewReader((r.(DbInstancesPage)).Body), &queryAvailableDbResponse)
	return queryAvailableDbResponse.Instances, err
}

type DbInstancesPage struct {
	pagination.NewSinglePageBase
}

type QueryAvailableDbResponse struct {
	// DB instances that can be used for creating a schema
	Instances []QueryAvailableRdsList `json:"instances"`
	// Which page the server starts returning items.
	Offset int `json:"offset"`
	// Number of records displayed on each page
	Limit int `json:"limit"`
	// Total Collections
	Total int `json:"total"`
}

// QueryAvailableRdsList represents the details of a DB instance.
type QueryAvailableRdsList struct {
	// DB instance ID
	ID string `json:"id"`
	// Project ID of the tenant that the DB instance belongs to
	ProjectID string `json:"projectId"`
	// DB instance status
	Status string `json:"status"`
	// DB instance name
	Name string `json:"name"`
	// Engine name of the DB instance
	EngineName string `json:"engineName"`
	// Engine version of the DB instance
	EngineSoftwareVersion string `json:"engineSoftwareVersion"`
	// Private IP address for connecting to the DB instance
	PrivateIP string `json:"privateIp"`
	// DB instance type (primary/standby or single-node)
	Mode string `json:"mode"`
	// Port for connecting to the DB instance
	Port int `json:"port"`
	// AZ
	AZCode string `json:"azCode"`
	// Time zone
	TimeZone string `json:"timeZone"`
}
