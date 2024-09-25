package products

import (
	"fmt"
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get products
func Get(client *golangsdk.ServiceClient, engine string) (*GetResp, error) {
	if len(engine) == 0 {
		return nil, fmt.Errorf("the parameter \"engine\" cannot be empty, it is required")
	}

	paths := strings.SplitN(client.Endpoint, "v2", 2)
	url := fmt.Sprintf("%sv2/%s/%s", paths[0], engine, "products")

	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetResp struct {
	// Message engine of DMS.
	Engine string `json:"engine"`
	// Supported versions.
	Versions []string `json:"versions"`
	// Product specification details.
	Products []EngineProduct `json:"products"`
}

type EngineProduct struct {
	// Product type. Currently, single-node and cluster types are supported.
	Type string `json:"type"`
	// Product ID.
	ProductId string `json:"product_id"`
	// ECS flavor.
	ECSFlavorId string `json:"ecs_flavor_id"`
	// Billing Code.
	BillingCode string `json:"billing_code"`
	// CPU architecture.
	ArchTypes []string `json:"arch_types"`
	// Billing mode. hourly: pay-per-use
	ChargingMode []string `json:"charging_mode"`
	// List of supported disk I/O types.
	IOS []EngineIOS `json:"ios"`
	// List of features supported by instances of the current specifications.
	SupportFeatures []*EngineSupportFeature `json:"support_features"`
	// Attribute of instances of the current specifications.
	Properties           EngineProperties `json:"properties"`
	QingtianIncompatible bool             `json:"qingtian_incompatible"`
}

type EngineIOS struct {
	// Disk I/O code.
	IOSpec string `json:"io_spec"`
	// Disk type.
	Type string `json:"type"`
	// Available AZs.
	AvailableZones []string `json:"available_zones"`
	// Unavailable AZs.
	UnavailableZones []string `json:"unavailable_zones"`
}

type EngineSupportFeature struct {
	// Feature name.
	Name string `json:"name"`
	// Description of the features supported by the instance.
	Properties EngineSupportFeaturesProperty `json:"properties"`
}

type EngineSupportFeaturesProperty struct {
	// Maximum number of dumping tasks.
	MaxTask string `json:"max_task"`
	// Minimum number of dumping tasks.
	MinTask string `json:"min_task"`
	// Maximum number of dumping nodes.
	MaxNode string `json:"max_node"`
	// Minimum number of dumping nodes.
	MinNode string `json:"min_node"`
}

type EngineProperties struct {
	NetworkBandwidth           string `json:"network_bandwidth"`
	MaxReplicaPerBroker        string `json:"max_replica_per_broker"`
	EngineVersions             string `json:"engine_versions"`
	MaxSSLConnectionsPerBroker string `json:"max_ssl_connections_per_broker"`

	// Maximum number of partitions of each broker.
	MaxPartitionPerBroker string `json:"max_partition_per_broker"`
	// Maximum number of brokers.
	MaxBroker string `json:"max_broker"`
	// Maximum storage space of each broker. The unit is GB.
	MaxStoragePerNode string `json:"max_storage_per_node"`
	// Maximum number of consumers of each broker.
	MaxConsumerPerBroker string `json:"max_consumer_per_broker"`
	// Minimum number of brokers.
	MinBroker string `json:"min_broker"`
	// Maximum bandwidth of each broker.
	MaxBandwidthPerBroker string `json:"max_bandwidth_per_broker"`
	// Minimum storage space of each broker. The unit is GB.
	MinStoragePerNode string `json:"min_storage_per_node"`
	// Maximum TPS of each broker.
	MaxTPSPerBroker string `json:"max_tps_per_broker"`
	// Alias of product_id.
	ProductAlias string `json:"product_alias"`
}
