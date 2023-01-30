package lifecycle

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// Get a instance with detailed information by id
func Get(client *golangsdk.ServiceClient, id string) (*Instance, error) {
	raw, err := client.Get(client.ServiceURL("instances", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Instance
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Instance struct {
	// DCS instance name.
	Name string `json:"name"`
	// DCS instance engine.
	Engine string `json:"engine"`
	// DCS instance cache capacity. Unit: GB.
	Capacity int `json:"capacity"`
	// Cache capacity of a small-scale, single-node DCS instance.
	CapacityMinor string `json:"capacity_minor"`
	// IP address for connecting to the DCS instance For a cluster instance,
	// multiple IP addresses are returned and separated by commas (,).
	// For example, 192.168.0.1,192.168.0.2.
	IP string `json:"ip"`
	// Port number of the cache node.
	Port int `json:"port"`
	// Cache instance status.
	Status string `json:"status"`
	// true: This instance is a libos instance.
	// false: This instance is not a libos instance.
	Libos bool `json:"libos"`
	// Brief description of the DCS instance.
	Description string `json:"description"`
	// Total memory size.
	// Unit: MB.
	MaxMemory int `json:"max_memory"`
	// Size of the used memory.
	// Unit: MB.
	UsedMemory int `json:"used_memory"`
	// DCS instance ID.
	InstanceID string `json:"instance_id"`
	// Resource specifications.
	// For example:
	// dcs.single_node: indicates a DCS instance in single-node mode.
	// dcs.master_standby: indicates a DCS instance in master/standby mode.
	// dcs.cluster: indicates a DCS instance in cluster mode.
	ResourceSpecCode string `json:"resource_spec_code"`
	// Cache engine version.
	EngineVersion string `json:"engine_version"`
	// Internal DCS version.
	InternalVersion string `json:"internal_version"`
	// Billing mode. 0: pay-per-use.
	ChargingMode int `json:"charging_mode"`
	// VPC ID.
	VPCID string `json:"vpc_id"`
	// VPC name.
	VPCName string `json:"vpc_name"`
	// Time at which the DCS instance is created.
	// For example, 2017-03-31T12:24:46.297Z.
	CreatedAt string `json:"created_at"`
	// Error code returned when the DCS instance fails to be created or is abnormal.
	ErrorCode string `json:"error_code"`
	// User ID.
	UserID string `json:"user_id"`
	// Username.
	UserName string `json:"user_name"`
	// Time at which the maintenance time window starts.
	// Format: hh:mm:ss.
	MaintainBegin string `json:"maintain_begin"`
	// Time at which the maintenance time window ends.
	// Format: hh:mm:ss.
	MaintainEnd string `json:"maintain_end"`
	// An indicator of whether a DCS instance can be accessed in password-free mode.
	// true: indicates that a DCS instance can be accessed without a password.
	// false: indicates that a DCS instance can be accessed only after password authentication.
	NoPasswordAccess string `json:"no_password_access"`
	// Username used for accessing a DCS instance with password authentication.
	AccessUser string `json:"access_user"`
	// An indicator of whether public access is enabled for a DCS Redis instance. Options:
	EnablePublicIp bool `json:"enable_publicip"`
	// ID of the elastic IP address bound to a DCS Redis instance.
	// The parameter value is null if public access is disabled.
	PublicIpId string `json:"publicip_id"`
	// Elastic IP address bound to a DCS Redis instance.
	// The parameter value is null if public access is disabled.
	PublicIpAddress string `json:"publicip_address"`
	// An indicator of whether to enable SSL for public access to a DCS Redis instance.
	EnableSsl bool `json:"enable_ssl"`
	// An indicator of whether an upgrade task has been created for a DCS instance.
	ServiceUpgrade bool `json:"service_upgrade"`
	// Upgrade task ID.
	// If the value of service_upgrade is set to true, the value of this parameter is the ID of the upgrade task.
	// If the value of service_upgrade is set to false, the value of this parameter is empty.
	ServiceTaskId string `json:"service_task_id"`
	// Edition of DCS for Redis. Options:
	// generic: standard edition
	// libos: high-performance edition
	ProductType string `json:"product_type"`
	// CPU architecture. Options: x86_64 and aarch_64.
	CpuType string `json:"cpu_type"`
	// Memory type. Options: DRAM and SCM.
	StorageType string `json:"storage_type"`
	// DCS instance type. Options:
	// single: single-node
	// ha: master/standby
	// cluster: Redis Cluster
	// proxy: Proxy Cluster
	CacheMode string `json:"cache_mode"`
	// Time when the instance started running. 2022-07-06T09:32:16.502Z
	LaunchedAt string `json:"launched_at"`
	// AZ where a cache node resides. The value of this parameter in the response contains an AZ ID.
	AvailableZones []string `json:"available_zones"`
	// Subnet ID.
	SubnetID string `json:"subnet_id"`
	// Security group ID.
	SecurityGroupID string `json:"security_group_id"`
	// Backend address of a cluster instance.
	BackendAddrs string `json:"backend_addrs"`
	// Cloud service type code.
	CloudServiceTypeCode string `json:"cloud_service_type_code"`
	// Cloud resource type code.
	CloudResourceTypeCode string `json:"cloud_resource_type_code"`
	// Security group name.
	SecurityGroupName string `json:"security_group_name"`
	// Subnet name.
	SubnetName string `json:"subnet_name"`
	// Subnet segment.
	SubnetCIDR string `json:"subnet_cidr"`
	// Order ID.
	OrderID string `json:"order_id"`
	// Backup policy.
	// This parameter is available for master/standby and cluster DCS instances.
	InstanceBackupPolicy BackupPolicy `json:"instance_backup_policy"`
	// Instance tag key and value.
	Tags []tags.ResourceTag `json:"tags"`
	// Product specification code.
	SpecCode string `json:"spec_code"`
	// Domain name of the instance.
	DomainName string `json:"domain_name"`
	// Read-only domain name.
	ReadonlyDomainName string `json:"readonly_domain_name"`
	// Scenario where the instance is frozen.
	FreezeScene string `json:"freeze_scene"`
	// Update time. 2022-07-06T09:32:16.502Z
	UpdateAt string `json:"update_at"`
	// Task status.
	TaskStatus string `json:"task_status"`
	// Whether the instance is free of charge.
	IsFree bool `json:"is_free"`
	// AZs with available resources.
	AzCodes []string `json:"az_codes"`
	// Role in cross-region DR.
	CrrRole string `json:"crr_role"`
	// Product specification code.
	InQuerySpecCode string `json:"inquery_spec_code"`
	// Whether slow query logs are supported.
	SupportSlowLogFlag string `json:"support_slow_log_flag"`
	// IPv6 address.
	Ipv6 string `json:"ipv6"`
	// Whether IPv6 is enabled.
	EnableIpv6 bool `json:"enable_ipv6"`
	// Number of databases in the instance.
	DbNumber int `json:"db_number"`
	// Whether ACL is supported.
	SupportAcl bool `json:"support_acl"`
	// Task response.
	Task string `json:"task"`
	// Number of shards.
	ShardingCount int `json:"sharding_count"`
}

type BackupPolicy struct {
	BackupPolicyId string               `json:"backup_policy_id"`
	CreatedAt      string               `json:"created_at"`
	UpdatedAt      string               `json:"updated_at"`
	Policy         InstanceBackupPolicy `json:"policy"`
	TenantId       string               `json:"tenant_id"`
}
