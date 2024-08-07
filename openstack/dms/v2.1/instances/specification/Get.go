package specification

import (
	"fmt"
	"net/http"
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
)

type GetSpecOpts struct {
	// Message engine.
	Engine string
	// Product edition.
	//    advanced: premium edition
	Type string `q:"type,omitempty"`
}

// GetSpec is used to query the product information for instance specification modification.
// Send GET /v2/{engine}/{project_id}/instances/{instance_id}/extend
func GetSpec(client *golangsdk.ServiceClient, instanceId string, opts GetSpecOpts) (*SpecResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(instances.ResourcePath, instanceId, extendPath).WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	urlS := client.ServiceURL(url.String())

	// Here we should patch a client, because for an increasing url path is different.
	// For all requests we use schema	/v2/{project_id}/instances
	// But for these					/v2/{engine}/{project_id}/instances
	paths := strings.SplitN(urlS, "v2", 2)
	urlStr := fmt.Sprintf("%sv2/%s%s", paths[0], opts.Engine, paths[1])

	var raw *http.Response
	raw, err = client.Get(urlStr, nil, nil)
	if err != nil {
		return nil, err
	}

	var res SpecResp
	err = extract.Into(raw.Body, &res)
	return &res, err

}

type SpecResp struct {
	// Message engine: Kafka.
	Engine string `json:"engine"`
	// Versions supported by the message engine.
	Versions []string `json:"versions"`
	// Product information for specification modification.
	Products []ExtendProductInfo `json:"products"`
}

type ExtendProductInfo struct {
	// Instance type.
	Type string `json:"type"`
	// Product ID.
	ProductId string `json:"product_id"`
	// ECS flavor used by the product.
	ECSFlavorId string `json:"ecs_flavor_id"`
	// Supported CPU architectures.
	ArchTypes []string `json:"arch_types"`
	// Supported billing modes.
	ChargingMode []string `json:"charging_mode"`
	// Disk I/O information.
	IOS []*ExtendProductIOS `json:"ios"`
	// Supported features.
	SupportFeatures []*SupportFeature `json:"support_features"`
	// Product specification description.
	Properties Properties `json:"properties"`
	// AZs where there are available resources.
	AvailableZones []string `json:"available_zones"`
	// AZs where resources are sold out.
	UnavailableZones []string `json:"unavailable_zones"`
}

type ExtendProductIOS struct {
	// Storage I/O specification.
	IOSpec string `json:"io_spec"`
	// AZs where there are available resources.
	AvailableZones []string `json:"available_zones"`
	// I/O type.
	Type string `json:"type"`
	// AZs where resources are sold out.
	UnavailableZones []string `json:"unavailable_zones"`
}

type SupportFeature struct {
	// Feature name.
	Name string `json:"name"`
	// Key-value pair of a feature.
	Properties map[string]string `json:"properties"`
}

type Properties struct {
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
