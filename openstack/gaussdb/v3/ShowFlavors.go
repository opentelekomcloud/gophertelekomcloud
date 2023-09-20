package v3

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ShowFlavorsOpts struct {
	// DB engine.
	DatabaseName string
	// DB engine version. Currently, only MySQL 8.0 is supported.
	VersionName string `q:"version_name"`
	// AZ mode. Its value can be single or multi and is case-insensitive.
	AvailabilityZoneMode string `q:"availability_zone_mode"`
	// Specification code.
	SpecCode string `q:"spec_code"`
}

func ShowGaussMySqlFlavors(client *golangsdk.ServiceClient, opts ShowFlavorsOpts) ([]MysqlFlavorsInfo, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/mysql/v3/{project_id}/flavors/{database_name}
	raw, err := client.Get(client.ServiceURL("flavors", opts.DatabaseName)+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []MysqlFlavorsInfo
	err = extract.IntoSlicePtr(raw.Body, &res, "flavors")
	return res, err
}

type MysqlFlavorsInfo struct {
	// Number of vCPUs. For example, the value 1 indicates 1 vCPU.
	Vcpus string `json:"vcpus"`
	// Memory size in GB
	Ram string `json:"ram"`
	// Specification type. The value can be arm
	Type string `json:"type"`
	// Specification ID. The value must be unique.
	Id string `json:"id"`
	// Resource specification code. Its value is same as the value of flavor_ref, for example, gaussdb.mysql.4xlarge.arm.8.
	SpecCode string `json:"spec_code"`
	// DB version number.
	VersionName string `json:"version_name"`
	// Instance type. Currently, its value can only be Cluster.
	InstanceMode string `json:"instance_mode"`
	// Specification status in an AZ. Its value can be any of the following:
	// normal: on sale.
	// unsupported: Not supported.
	// sellout: sold out.
	AzStatus map[string]string `json:"az_status"`
}
