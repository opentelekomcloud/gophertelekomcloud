package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// Instance response
type Instance struct {
	// Instance name.
	Name string `json:"name"`
	// Message engine.
	Engine string `json:"engine"`
	// Version.
	EngineVersion string `json:"engine_version"`
	// Instance description.
	Description string `json:"description"`
	// Instance specifications.
	Specification string `json:"specification"`
	// Message storage space in GB.
	StorageSpace int `json:"storage_space"`
	// Number of partitions in a Kafka instance.
	PartitionNum string `json:"partition_num"`
	// Used message storage space in GB.
	UsedStorageSpace int `json:"used_storage_space"`
	// IP address of an instance.
	ConnectAddress string `json:"connect_address"`
	// Port of an instance.
	Port int `json:"port"`
	// Instance status.
	Status string `json:"status"`
	// Instance ID.
	InstanceID string `json:"instance_id"`
	// Resource specification code.
	//    dms.instance.kafka.cluster.c3.mini: Kafka instance with 100 MB/s bandwidth
	//    dms.instance.kafka.cluster.c3.small.2: Kafka instance with 300 MB/s bandwidth
	//    dms.instance.kafka.cluster.c3.middle.2: Kafka instance with 600 MB/s bandwidth
	//    dms.instance.kafka.cluster.c3.high.2: Kafka instance with 1200 MB/s bandwidth
	ResourceSpecCode string `json:"resource_spec_code"`
	// Billing mode. 1: pay-per-use.
	ChargingMode int `json:"charging_mode"`
	// VPC id.
	VPCID string `json:"vpc_id"`
	// VPC name.
	VPCName string `json:"vpc_name"`
	// Time when the instance was created.
	// The time is in the format of timestamp, that is, the offset milliseconds from 1970-01-01 00:00:00 UTC to the specified time.
	CreatedAt string `json:"created_at"`
	// Subnet name.
	SubnetName string `json:"subnet_name"`
	// Subnet CIDR block.
	SubnetCIDR string `json:"subnet_cidr"`
	// User ID.
	UserID string `json:"user_id"`
	// User name.
	UserName string `json:"user_name"`
	// Username for accessing the instance.
	AccessUser string `json:"access_user"`
	// Time at which the maintenance time window starts. The format is HH:mm:ss.
	MaintainBegin string `json:"maintain_begin"`
	// Time at which the maintenance time window ends. The format is HH:mm:ss.
	MaintainEnd string `json:"maintain_end"`
	// Whether public access is enabled for the instance.
	//    true: enabled
	//    false: disabled
	EnablePublicIP bool `json:"enable_publicip"`
	// Whether security authentication is enabled.
	//    true: enable
	//    false: disabled
	SslEnable bool `json:"ssl_enable"`
	// Indicates whether to enable encrypted replica transmission among brokers.
	//    true: enable
	//    false: disable
	BrokerSslEnable bool `json:"broker_ssl_enable"`
	// Security protocol to use after SASL is enabled.
	//    SASL_SSL: Data is encrypted with SSL certificates for high-security transmission.
	//    SASL_PLAINTEXT: Data is transmitted in plaintext with username and password authentication. This protocol uses the SCRAM-SHA-512 mechanism and delivers high performance.
	KafkaSecurityProtocol string `json:"kafka_security_protocol"`
	// Authentication mechanism used after SASL is enabled.
	//    PLAIN: simple username and password verification.
	//    SCRAM-SHA-512: user credential verification, which is more secure than PLAIN.
	SASLEnabledMechanisms []string `json:"sasl_enabled_mechanisms"`
	// Indicates whether to enable two-way authentication.
	SSLTwoWayEnable bool `json:"ssl_two_way_enable"`
	// Whether the certificate can be replaced.
	CertReplaced bool `json:"cert_replaced"`
	// Enterprise project ID.
	EnterpriseProjectID string `json:"enterprise_project_id"`
	// Instance type. The value can be a cluster.
	Type string `json:"type"`
	// Product ID.
	ProductID string `json:"product_id"`
	// Security group ID.
	SecurityGroupID string `json:"security_group_id"`
	// Security group name.
	SecurityGroupName string `json:"security_group_name"`
	// Subnet ID.
	SubnetID string `json:"subnet_id"`
	// AZ to which the instance brokers belong. The AZ ID is returned.
	AvailableZones []string `json:"available_zones"`
	// Name of the AZ to which the instance node belongs. The AZ name is returned.
	AvailableZoneNames []string `json:"available_zone_names"`
	// Message storage space in GB.
	TotalStorageSpace int `json:"total_storage_space"`
	// Instance public access address. This parameter is available only when public access is enabled for the instance.
	PublicConnectAddress string `json:"public_connect_address"`
	// Storage resource ID.
	StorageResourceID string `json:"storage_resource_id"`
	// I/O specifications.
	StorageSpecCode string `json:"storage_spec_code"`
	// Service type.
	ServiceType string `json:"service_type"`
	// Storage class.
	StorageType string `json:"storage_type"`
	// Message retention policy.
	RetentionPolicy string `json:"retention_policy"`
	// Whether public access is enabled for Kafka.
	KafkaPublicStatus string `json:"kafka_public_status"`
	// Public network access bandwidth.
	PublicBandWidth int `json:"public_bandwidth"`
	// Whether Kafka Manager is enabled.
	KafkaManagerEnabled bool `json:"kafka_manager_enabled"`
	// Indicates whether to enable a new certificate.
	NewAuthCert bool `json:"new_auth_cert"`
	// Cross-VPC access information.
	CrossVpcInfo string `json:"cross_vpc_info"`
	// Number of connectors.
	ConnectorNodeNum int `json:"connector_node_num"`
	// Kafka REST connection address.
	RestConnectAddress string `json:"rest_connect_address"`
	// Connection address on the tenant side.
	PodConnectAddress string `json:"pod_connect_address"`
	// Whether disk encryption is enabled.
	DiskEncrypted bool `json:"disk_encrypted"`
	// Disk encryption key. If disk encryption is not enabled, this parameter is left blank.
	DiskEncryptedKey string `json:"disk_encrypted_key"`
	// Private connection address of a Kafka instance.
	KafkaPrivateConnectAddress string `json:"kafka_private_connect_address"`
	// Cloud Eye version.
	CesVersion string `json:"ces_version"`
	// Time when public access was enabled for an instance. The value can be true, actived, closed, or false.
	PublicAccessEnabled string `json:"public_access_enabled"`
	// Node quantity.
	NodeNum int `json:"node_num"`
	// Indicates whether access control is enabled.
	EnableACL bool `json:"enable_acl"`
	// Whether billing based on new specifications is enabled.
	NewSpecBillingEnabled bool `json:"new_spec_billing_enabled"`
	// Broker quantity.
	BrokerNum int `json:"broker_num"`
	// Tag list.
	Tags []tags.ResourceTag `json:"tags"`
	// Indicates whether DR is enabled.
	DREnable bool `json:"dr_enable"`

	// Possible these fields are deprecated.
	ManagementConnectAddress string `json:"management_connect_address"`
	MessageQueryInstEnable   bool   `json:"message_query_inst_enable"`
	SupportFeatures          string `json:"support_features"`
}

// Get an instance with detailed information by id
// Send GET /v2/{project_id}/instances/{instance_id}
func Get(client *golangsdk.ServiceClient, id string) (*Instance, error) {
	raw, err := client.Get(client.ServiceURL(ResourcePath, id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Instance
	err = extract.Into(raw.Body, &res)
	return &res, err
}
