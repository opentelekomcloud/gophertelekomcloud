package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Instance name. Starts with a letter, consists of 4 to 64 characters,
	// and can contain only letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name" required:"true"`
	// Description of an instance. 0-1024 characters.
	Description string `json:"description,omitempty"`
	// Message engine type.
	//    rocketmq: RocketMQ message engine.
	//    reliability: RocketMQ message engine alias.
	Engine string `json:"engine" required:"true"`
	// Message engine version. Value: 5.x.
	EngineVersion string `json:"engine_version" required:"true"`
	// Storage space, in GB. The documented ranges are
	//    RocketMQ 5.x single-node: 100-30000
	//    RocketMQ 5.x cluster: 200-60000
	// but they do not always match what the service accepts: eu-de reports a
	// 300 GB minimum for rocketmq.b1.large.1. Prefer min_storage_per_node /
	// max_storage_per_node from the products API over these values.
	StorageSpace int `json:"storage_space" required:"true"`
	// VPC ID.
	VpcID string `json:"vpc_id" required:"true"`
	// Subnet ID.
	SubnetID string `json:"subnet_id" required:"true"`
	// Security group to which the instance belongs.
	SecurityGroupID string `json:"security_group_id" required:"true"`
	// IDs of the AZs where instance brokers reside and which have available resources.
	// A RocketMQ instance can be deployed in one or three AZs.
	AvailableZones []string `json:"available_zones" required:"true"`
	// RocketMQ instance flavor.
	// If type is single.basic, select single-node flavors, if type is cluster.basic, select cluster flavors.
	//    rocketmq.b1.large.1: single-node flavor, instance TPS 500
	//    rocketmq.b2.large.4: cluster flavor, instance TPS 2,000
	//    rocketmq.b2.large.8: cluster flavor, instance TPS 4,000
	//    rocketmq.b2.large.12: cluster flavor, instance TPS 6,000
	ProductID string `json:"product_id" required:"true"`
	// Whether to enable SSL-encrypted access. Default: false.
	SslEnable *bool `json:"ssl_enable,omitempty"`
	// Storage I/O flavor.
	//    dms.physical.storage.high.v2: high I/O disk
	//    dms.physical.storage.ultra.v2: ultra-high I/O disk
	//    dms.physical.storage.general: general-purpose SSD
	//    dms.physical.storage.extreme: extreme SSD
	StorageSpecCode string `json:"storage_spec_code" required:"true"`
	// Enterprise project ID. Mandatory for an enterprise project account.
	EnterpriseProjectID string `json:"enterprise_project_id,omitempty"`
	// Whether to enable access control.
	EnableACL *bool `json:"enable_acl,omitempty"`
	// Whether to enable the proxy function. Default: false.
	ProxyEnable *bool `json:"proxy_enable,omitempty"`
	// Whether to enable public access. Default: false.
	EnablePublicIP *bool `json:"enable_publicip,omitempty"`
	// IDs of the EIPs bound to the instance, separated by commas (,).
	// Mandatory if public access is enabled.
	PublicIpID string `json:"publicip_id,omitempty"`
	// Number of brokers.
	BrokerNum int `json:"broker_num" required:"true"`
	// Architecture type.
	//    X86: complex instruction set compute.
	//    ARM: simplified instruction set compute.
	ArchType string `json:"arch_type,omitempty"`
	// Security protocol used by an instance.
	TlsMode string `json:"tls_mode,omitempty"`
}

// CreateResponse is returned by Create and holds the ID of the new instance.
type CreateResponse struct {
	// Instance ID.
	InstanceID string `json:"instance_id"`
}

// Create a RocketMQ instance.
// Send POST to /v2/{project_id}/rocketmq/instances
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("rocketmq", "instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
