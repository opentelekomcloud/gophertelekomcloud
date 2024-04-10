package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListStorageTypesOpts struct {
	// Specifies the DB engine name. Its value can be any of the following and is case-insensitive:
	// MySQL
	// PostgreSQL
	// SQLServer
	DatabaseName string
	// Specifies the database version.
	VersionName string `q:"version_name" required:"true"`
}

func ListStorageTypes(client *golangsdk.ServiceClient, opts ListStorageTypesOpts) (*ListStorageTypesResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("storage-type", opts.DatabaseName).WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/storage-type/{database_name}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListStorageTypesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListStorageTypesResponse struct {
	// Indicates the DB instance specifications information list.
	StorageType []Storage `json:"storage_type,omitempty"`
	// Indicates the dsspool specifications information list.
	DsspoolInfo []DssPoolInfo `json:"dsspool_info,omitempty"`
}

type Storage struct {
	// Indicates the storage type. Its value can be any of the following:
	// COMMON: SATA storage.
	// ULTRAHIGH: SSD storage.
	Name string `json:"name"`
	// Indicates the specification status in an AZ. Its value can be any of the following:
	// normal: indicates that the specifications in the AZ are available.
	// unsupported: indicates that the specifications are not supported by the AZ.
	// sellout: indicates that the specifications in the AZ are sold out.
	AzStatus map[string]string `json:"az_status"`
	// Indicates the performance specifications. Its value can be any of the following:
	// normal: general-enhanced
	// normal2: general-enhanced II
	// armFlavors: Kunpeng general-enhanced
	// highPerformancePrivilegeEdition: Ultra-high I/O (advanced)
	SupportComputeGroupType []string `json:"support_compute_group_type,omitempty"`
}

type DssPoolInfo struct {
	// Indicates the name of the AZ where dsspool is located.
	AzName string `json:"az_name"`
	// Indicates the available capacity of dsspool.
	FreeCapacityGb string `json:"free_capacity_gb"`
	// Indicates the dsspool volume type.
	DsspoolVolumeType string `json:"dsspool_volume_type"`
	// Indicates the dsspool ID.
	DsspoolId string `json:"dsspool_id"`
	// Indicates the dsspool status. Its value can be any of the following:
	// available
	// deploying
	// enlarging
	// frozen
	// sellout
	DsspoolStatus string `json:"dsspool_status"`
}
