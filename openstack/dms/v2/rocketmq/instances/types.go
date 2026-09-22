package instances

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"

// Instance contains the details of a RocketMQ instance.
type Instance struct {
	// Instance name.
	Name string `json:"name"`
	// Message engine.
	//    rocketmq: RocketMQ message engine.
	//    reliability: RocketMQ message engine alias.
	Engine string `json:"engine"`
	// Instance status.
	Status string `json:"status"`
	// Instance description.
	Description string `json:"description"`
	// Instance type.
	//    single.basic: 5.x single-node basic edition.
	//    cluster.basic: 5.x cluster basic edition.
	Type string `json:"type"`
	// Instance specification.
	Specification string `json:"specification"`
	// Instance version.
	EngineVersion string `json:"engine_version"`
	// Instance ID.
	InstanceID string `json:"instance_id"`
	// Billing mode. 1 indicates pay-per-use.
	ChargingMode int `json:"charging_mode"`
	// VPC ID.
	VpcID string `json:"vpc_id"`
	// VPC name.
	VpcName string `json:"vpc_name"`
	// Time when creation is complete, as the offset in milliseconds
	// from 1970-01-01 00:00:00 UTC to the specified time.
	CreatedAt string `json:"created_at"`
	// RocketMQ instance flavor.
	ProductID string `json:"product_id"`
	// Security group ID.
	SecurityGroupID string `json:"security_group_id"`
	// Security group name.
	SecurityGroupName string `json:"security_group_name"`
	// Subnet ID.
	SubnetID string `json:"subnet_id"`
	// Subnet name. Returned when querying a single instance.
	SubnetName string `json:"subnet_name"`
	// Subnet route.
	SubnetCIDR string `json:"subnet_cidr"`
	// List of AZ IDs.
	AvailableZones []string `json:"available_zones"`
	// List of AZ names.
	AvailableZoneNames []string `json:"available_zone_names"`
	// User ID.
	UserID string `json:"user_id"`
	// Username.
	UserName string `json:"user_name"`
	// Time at which the maintenance window starts, in format HH:mm:ss.
	MaintainBegin string `json:"maintain_begin"`
	// Time at which the maintenance window ends, in format HH:mm:ss.
	MaintainEnd string `json:"maintain_end"`
	// Storage space, in GB.
	StorageSpace int `json:"storage_space"`
	// Used message storage space, in GB.
	UsedStorageSpace int `json:"used_storage_space"`
	// Total storage space, in GB.
	TotalStorageSpace int `json:"total_storage_space"`
	// Whether public access is enabled.
	EnablePublicIP bool `json:"enable_publicip"`
	// IDs of the EIPs bound to the instance, separated by commas (,).
	PublicIpID string `json:"publicip_id"`
	// Public IP address.
	PublicIpAddress string `json:"publicip_address"`
	// Whether SSL is enabled.
	SslEnable bool `json:"ssl_enable"`
	// Cross-VPC access information.
	CrossVpcInfo string `json:"cross_vpc_info"`
	// Storage resource ID.
	StorageResourceID string `json:"storage_resource_id"`
	// Storage specification code.
	//    dms.physical.storage.high.v2: high I/O disk
	//    dms.physical.storage.ultra.v2: ultra-high I/O disk
	//    dms.physical.storage.general: general-purpose SSD
	//    dms.physical.storage.extreme: extreme SSD
	StorageSpecCode string `json:"storage_spec_code"`
	// Service type.
	ServiceType string `json:"service_type"`
	// Storage type.
	StorageType string `json:"storage_type"`
	// Whether IPv6 is enabled.
	IPv6Enable bool `json:"ipv6_enable"`
	// Whether disk encryption is enabled.
	DiskEncrypted bool `json:"disk_encrypted"`
	// Whether access control is enabled.
	EnableACL bool `json:"enable_acl"`
	// Number of brokers.
	BrokerNum int `json:"broker_num"`
	// Whether domain name access to an instance is enabled.
	DNSEnable bool `json:"dns_enable"`
	// Metadata address.
	NameSrvAddress string `json:"namesrv_address"`
	// Metadata domain name.
	NameSrvDomainName string `json:"namesrv_domain_name"`
	// Service data address.
	BrokerAddress string `json:"broker_address"`
	// Public network metadata address.
	PublicNameSrvAddress string `json:"public_namesrv_address"`
	// Public metadata domain name.
	PublicNameSrvDomainName string `json:"public_namesrv_domain_name"`
	// Public network service data address.
	PublicBrokerAddress string `json:"public_broker_address"`
	// gRPC connection address, displayed only for RocketMQ 5.x.
	GrpcAddress string `json:"grpc_address"`
	// gRPC connection domain name, displayed only for RocketMQ 5.x.
	GrpcDomainName string `json:"grpc_domain_name"`
	// Public gRPC connection address, displayed only for RocketMQ 5.x.
	PublicGrpcAddress string `json:"public_grpc_address"`
	// Public gRPC domain name, displayed only for RocketMQ 5.x.
	PublicGrpcDomainName string `json:"public_grpc_domain_name"`
	// Enterprise project ID.
	EnterpriseProjectID string `json:"enterprise_project_id"`
	// Tag list.
	Tags []tags.ResourceTag `json:"tags"`
	// Resource specification.
	ResourceSpecCode string `json:"resource_spec_code"`
	// Production TPS proportion.
	ProducePortion int `json:"produce_portion"`
	// Consumption TPS proportion.
	ConsumePortion int `json:"consume_portion"`
	// Whether the instance has disaster recovery (DR).
	DREnable bool `json:"dr_enable"`
	// Whether a restart is required to configure SSL.
	ConfigSslNeedRestartProcess bool `json:"config_ssl_need_restart_process"`
	// Security protocol used by an instance.
	TlsMode string `json:"tls_mode"`
	// Architecture type.
	//    X86: complex instruction set compute.
	//    ARM: simplified instruction set compute.
	ArchType string `json:"arch_type"`
	// Whether automatic disk capacity expansion is enabled.
	AutoVolumeExpandEnable bool `json:"auto_volume_expand_enable"`
}
