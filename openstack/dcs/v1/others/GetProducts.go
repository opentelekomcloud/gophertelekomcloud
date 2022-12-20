package others

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetProducts(client *golangsdk.ServiceClient) ([]Product, error) {
	// remove projectId from endpoint
	// GET /v1.0/products
	raw, err := client.Get(strings.Replace(client.ServiceURL("products"), "/"+client.ProjectID, "", -1), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Product
	err = extract.IntoSlicePtr(raw.Body, &res, "products")
	return res, err
}

type Product struct {
	// Product ID used to differentiate DCS specifications.
	ProductID string `json:"product_id"`
	// DCS instance specification code.
	SpecCode string `json:"spec_code"`
	// DCS instance type. Options:
	// single: single-node
	// ha: master/standby
	// cluster: Redis Cluster
	// proxy: Proxy Cluster
	// ha_rw_split: read/write splitting
	CacheMode string `json:"cache_mode"`
	// Edition of DCS for Redis.
	ProductType string `json:"product_type"`
	// CPU architecture.
	CpuType string `json:"cpu_type"`
	// Storage type.
	StorageType string `json:"storage_type"`
	// Details of the specifications.
	Details Detail `json:"details"`
	// Cache engine.
	Engine string `json:"engine"`
	// Cache engine version.
	EngineVersions string `json:"engine_versions"`
	// DCS specifications. The value subjects to the returned specifications.
	SpecDetails string `json:"spec_details"`
	// Detailed DCS specifications, including the maximum number of connections and maximum memory size.
	SpecDetails2 string `json:"spec_details2"`
	// Billing mode. Value: Hourly.
	ChargingType string `json:"charging_type"`
	// Price of the DCS service to which you can subscribe. (This parameter has been abandoned.)
	Price float64 `json:"price"`
	// Currency.
	Currency string `json:"currency"`
	// Product type.
	// Options: instance and obs_space.
	ProdType string `json:"prod_type"`
	// Cloud service type code.
	CloudServiceTypeCode string `json:"cloud_service_type_code"`
	// Cloud resource type code.
	CloudResourceTypeCode string `json:"cloud_resource_type_code"`
	// AZs with available resources.
	Flavors []Flavor `json:"flavors"`
	// Billing item.
	BillingFactor string `json:"billing_factor"`
}

type Detail struct {
	// Specification (total memory) of the DCS instance.
	Capacity float32 `json:"capacity"`
	// Maximum bandwidth supported by the specification.
	MaxBandwidth int `json:"max_bandwidth"`
	// Maximum number of clients supported by the specification,
	// which is usually equal to the maximum number of connections.
	MaxClients int `json:"max_clients"`
	// Maximum number of connections supported by the specification.
	MaxConnections int `json:"max_connections"`
	// Maximum inbound bandwidth supported by the specification,
	// which is usually equal to the maximum bandwidth.
	MaxInBandwidth int `json:"max_in_bandwidth"`
	// Maximum available memory.
	MaxMemory float32 `json:"max_memory"`
	// Number of tenant IP addresses corresponding to the specifications.
	TenantIpCount int `json:"tenant_ip_count"`
	// Number of shards supported by the specifications.
	ShardingNum int `json:"sharding_num"`
	// Number of proxies supported by Proxy Cluster instances of the specified specifications.
	// If the instance is not a Proxy Cluster instance, the value of this parameter is 0.
	ProxyNum int `json:"proxy_num"`
	// Number of DBs of the specifications.
	DbNumber int `json:"db_number"`
}

type Flavor struct {
	// Specification (total memory) of the DCS instance.
	Capacity string `json:"capacity"`
	// Memory unit.
	Unit string `json:"unit"`
	// AZ ID.
	AvailableZones []string `json:"available_zones"`
}
