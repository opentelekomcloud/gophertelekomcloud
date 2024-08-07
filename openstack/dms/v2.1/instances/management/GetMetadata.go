package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
)

const (
	managementPath = "management"
	clusterPath    = "cluster"
)

// GetMetadata is used to query Kafka cluster metadata.
// Send GET /v2/{project_id}/instances/{instance_id}/management/cluster
func GetMetadata(client *golangsdk.ServiceClient, instanceId string) (*GetMetadataResp, error) {
	raw, err := client.Get(client.ServiceURL(instances.ResourcePath, instanceId, managementPath, clusterPath), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetMetadataResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetMetadataResp struct {
	// Details of the result of the cross-VPC access modification.
	Cluster *Cluster `json:"cluster"`
}

type Cluster struct {
	// Controller ID.
	Controller string `json:"controller"`
	// Broker list.
	Brokers []*Broker `json:"brokers"`
	// Number of topics.
	TopicCount int `json:"topics_count"`
	// Number of partitions.
	PartitionsCount int `json:"partitions_count"`
	// Number of online partitions.
	OnlinePartitionCount int `json:"online_partitions_count"`
	// Number of replicas.
	ReplicasCount int `json:"replicas_count"`
	// Total number of in-sync replicas (ISRs).
	ISRReplicasCount int `json:"isr_replicas_count"`
	// Number of consumer groups.
	ConsumersCount int `json:"consumers_count"`
}

type Broker struct {
	// Broker IP address.
	Host string `json:"host"`
	// Port number.
	Port int `json:"port"`
	// Broker ID.
	BrokerID string `json:"broker_id"`
	// Whether the broker is a controller.
	IsController bool `json:"is_controller"`
	// Server version.
	Version string `json:"version"`
	// Broker registration time, which is a Unix timestamp.
	RegisterTime int64 `json:"register_time"`
	// Whether Kafka brokers can be connected.
	IsHealth bool `json:"is_health"`
}
