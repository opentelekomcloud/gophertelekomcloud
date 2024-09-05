package instances

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type QueryInstancesOpts struct {
	// Specifies the Index offset.
	// The query starts from the next piece of data indexed by this parameter. The value is 0 by default.
	// The value must be a positive integer.
	Offset int `q:"offset"`
	// A maximum of instances to be queried.
	// Value range: 1 to 128.
	// If the parameter value is not specified, 10 DDM instances are queried by default.
	Limit int `q:"limit"`
}

// This function is used to query DDM instances.
func QueryInstances(client *golangsdk.ServiceClient, opts QueryInstancesOpts) ([]ShowInstanceBeanResponse, error) {
	// GET /v1/{project_id}/instances?offset={offset}&limit={limit}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return InstancesPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractInstances(pages)
}

func ExtractInstances(r pagination.NewPage) ([]ShowInstanceBeanResponse, error) {
	var queryInstancesResponse QueryInstancesResponse
	err := extract.Into(bytes.NewReader((r.(InstancesPage)).Body), &queryInstancesResponse)
	return queryInstancesResponse.Instances, err
}

type InstancesPage struct {
	pagination.NewSinglePageBase
}

type QueryInstancesResponse struct {
	// DDM instance information.
	Instances []ShowInstanceBeanResponse `json:"instances"`
	// Number of DDM instances of a tenant.
	InstanceNum int `json:"instance_num"`
	// Current page.
	PageNo int `json:"page_no"`
	// Data records on the current page.
	PageSize int `json:"page_size"`
	// Total records.
	TotalRecord int `json:"total_record"`
	// Total pages.
	TotalPage int `json:"total_page"`
}

type ShowInstanceBeanResponse struct {
	// DDM instance ID.
	Id string `json:"id"`
	// DDM instance status.
	Status string `json:"status"`
	// Name of the created DDM instance.
	Name string `json:"name"`
	// Time when the DDM instance is created.
	Created string `json:"created"`
	// Time when the DDM instance is last updated.
	Updated string `json:"updated"`
	// AZ name.
	AvailableZone string `json:"available_zone"`
	// VPC ID.
	VpcId string `json:"vpc_id"`
	// Subnet ID.
	SubnetId string `json:"subnet_id"`
	// Security group ID.
	SecurityGroupId string `json:"security_group_id"`
	// Number of nodes.
	NodeCount int `json:"node_count"`
	// Address for accessing the DDM instance.
	AccessIp string `json:"access_ip"`
	// Port for accessing the DDM instance.
	AccessPort string `json:"access_port"`
	// Number of CPUs.
	CoreCount string `json:"core_count"`
	// Memory size in GB.
	RamCapacity string `json:"ram_capacity"`
	// Response message. Not returned if no abnormality occurs.
	ErrorMsg string `json:"error_msg,omitempty"`
	// Node status.
	NodeStatus string `json:"node_status"`
	// Project ID of a tenant in a region.
	ProjectId string `json:"project_id"`
	// Engine version (core instance version).
	EngineVersion string `json:"engine_version"`
}
