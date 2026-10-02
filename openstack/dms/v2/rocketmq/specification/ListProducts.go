package specification

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListProductsOpts struct {
	// Product type. advanced: premium edition.
	Type string `q:"type,omitempty"`
	// Number of records to query.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
}

// ListProductsResponse is returned by ListProducts.
type ListProductsResponse struct {
	// Total number of products.
	Total int `json:"total"`
	// Offset of the next page.
	NextOffset int `json:"next_offset"`
	// Offset of the previous page.
	PreviousOffset int `json:"previous_offset"`
	// Message engine type: rocketmq or its alias reliability.
	Engine string `json:"engine"`
	// Versions supported by the message engine.
	Versions []string `json:"versions"`
	// Product information for specification modification.
	Products []Product `json:"products"`
}

type Product struct {
	// Instance type.
	//    single.basic: 5.x single-node basic edition.
	//    cluster.basic: 5.x cluster basic edition.
	Type string `json:"type"`
	// Product ID, for example rocketmq.b2.large.4.
	ProductID string `json:"product_id"`
	// ID of the ECS flavor.
	EcsFlavorID string `json:"ecs_flavor_id"`
	// Billing code.
	BillingCode string `json:"billing_code"`
	// Supported CPU architectures.
	ArchTypes []string `json:"arch_types"`
	// Supported billing modes.
	ChargingMode []string `json:"charging_mode"`
	// Disk I/O information.
	IOs []ProductIO `json:"ios"`
	// Product properties.
	Properties ProductProperties `json:"properties"`
	// AZ list.
	AvailableZones []string `json:"available_zones"`
	// Unavailable AZ list.
	UnavailableZones []string `json:"unavailable_zones"`
	// Supported feature list.
	SupportFeatures []SupportFeature `json:"support_features"`
	// Whether the instance is a QingTian one.
	QingtianIncompatible bool `json:"qingtian_incompatible"`
}

type ProductIO struct {
	// Storage I/O flavor.
	IOSpec string `json:"io_spec"`
	// AZ list.
	AvailableZones []string `json:"available_zones"`
	// I/O type: evs.
	Type string `json:"type"`
	// Unavailable AZ list.
	UnavailableZones []string `json:"unavailable_zones"`
}

type ProductProperties struct {
	// Message engine version: 5.x.
	EngineVersions string `json:"engine_versions"`
	// Maximum storage space of each broker, in GB.
	MaxStoragePerNode string `json:"max_storage_per_node"`
	// Minimum storage space of each broker, in GB.
	MinStoragePerNode string `json:"min_storage_per_node"`
	// Alias of product_id.
	ProductAlias string `json:"product_alias"`
	// Feature switch of the specification.
	Feature string `json:"feature"`
	// Maximum number of topics in an instance.
	MaxTopic string `json:"max_topic"`
	// Number of brokers.
	BrokerNum string `json:"broker_num"`
	// Number of billing cores in an instance.
	Core string `json:"core"`
	// Maximum number of consumers in an instance.
	MaxConsumer string `json:"max_consumer"`
	// Traffic unit, rcu x max_tps_per_rcu = maximum flavor TPS.
	RCU string `json:"rcu"`
	// Maximum storage space, in GB.
	MaxStorage string `json:"max_storage"`
	// Minimum storage space, in GB.
	MinStorage string `json:"min_storage"`
	// Maximum TPS per RCU.
	MaxTPSPerRCU string `json:"max_tps_per_rcu"`
}

type SupportFeature struct {
	// Feature name.
	Name string `json:"name"`
	// Key-value pairs of the feature.
	Properties map[string]string `json:"properties"`
}

// ListProducts queries the product information for instance specification modification.
// Send GET to /v2/{project_id}/rocketmq/instances/{instance_id}/extend
func ListProducts(client *golangsdk.ServiceClient, instanceID string, opts ListProductsOpts) (*ListProductsResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("rocketmq", "instances", instanceID, "extend").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListProductsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
