package products

import (
	"fmt"
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Engine string `json:"-"`
	// The product ID.
	ProductId string `q:"product_id"`
	// Number of records to query.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
}

// Get products
func List(client *golangsdk.ServiceClient, opts ListOpts) (*GetResp, error) {
	if len(opts.Engine) == 0 {
		return nil, fmt.Errorf("the parameter \"engine\" cannot be empty, it is required")
	}

	paths := strings.SplitN(client.Endpoint, "v2", 2)
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%sv2/%s/%s", paths[0], opts.Engine, "products")

	raw, err := client.Get(url+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetResp struct {
	// Total number of products.
	Total int `json:"total"`
	// Offset of the next page.
	NextOffset int `json:"next_offset"`
	// Offset of the previous page.
	PreviousOffset int `json:"previous_offset"`
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

	// RocketMQ only properties.

	// Maximum number of topics.
	MaxTopic string `json:"max_topic"`
	// Broker quantity.
	BrokerNum string `json:"broker_num"`
	// Number of billing cores of an entire instance.
	Core string `json:"core"`
	// Maximum number of consumers in an instance.
	MaxConsumer string `json:"max_consumer"`
	// Traffic unit, rcu x max_tps_per_rcu = maximum flavor TPS.
	RCU string `json:"rcu"`
	// Maximum storage space, in GB.
	MaxStorage string `json:"max_storage"`
	// Minimum storage space, in GB.
	MinStorage string `json:"min_storage"`
	// Maximum TPS of each RCU.
	MaxTPSPerRCU string `json:"max_tps_per_rcu"`
	// Maximum number of topics that can be created on each broker.
	MaxTopicPerBroker string `json:"max_topic_per_broker"`
}
