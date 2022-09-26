package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowGaussMySqlInstanceInfo(client *golangsdk.ServiceClient, instanceId string) (*MysqlInstanceInfoDetail, error) {
	// GET https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}
	raw, err := client.Get(client.ServiceURL("instances", instanceId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res MysqlInstanceInfoDetail
	err = extract.IntoSlicePtr(raw.Body, &res, "instance")
	return &res, err
}

type MysqlInstanceInfoDetail struct {
	// Instance ID
	Id string `json:"id"`
	// Instance name
	Name string `json:"name"`
	// Project ID of a tenant in a region
	ProjectId string `json:"project_id"`
	// Instance status If the value is BUILD, the instance is being created.
	// If the value is ACTIVE, the instance is normal. If the value is FAILED, the instance is abnormal.
	// If the value is FROZEN, the instance is frozen. If the value is MODIFYING, the instance is being scaled up.
	// If the value is REBOOTING, the instance is being rebooted. If the value is RESTORING, the instance is being restored.
	// If the value is MODIFYING INSTANCE TYPE, the instance is changing from primary node to a read replica.
	// If the value is SWITCHOVER, the primary/standby switchover or failover is being performed.
	// If the value is MIGRATING, the instance is being migrated. If the value is BACKING UP, the instance is being backed up.
	// If the value is MODIFYING DATABASE PORT, the database port is being changed.
	// If the value is STORAGE FULL, the instance storage space is full.
	Status string `json:"status"`
	// Database port
	Port string `json:"port"`
	// Instance type. The value is Cluster.
	Type string `json:"type"`
	// The number of nodes.
	NodeCount int32 `json:"node_count"`
	// Database information
	Datastore MysqlDatastore `json:"datastore"`
	// Used backup space in GB
	BackupUsedSpace int64 `json:"backup_used_space"`
	// Creation time in the "yyyy-mm-ddThh:mm:ssZ" format.
	// T is the separator between the calendar and the hourly notation of time.
	// Z indicates the time zone offset.
	// For example, for French Winter Time (FWT), the time offset is shown as +0200.
	// The value is empty unless the instance creation is complete.
	Created string `json:"created"`
	// Update time. The format is the same as that of the created field.
	// The value is empty unless the instance creation is complete.
	Updated string `json:"updated"`
	// Private IP address for write
	PrivateWriteIps []string `json:"private_write_ips"`
	// Public IP address of the instance
	PublicIps string `json:"public_ips"`
	// Default username
	DbUserName string `json:"db_user_name"`
	// VPC ID
	VpcId string `json:"vpc_id"`
	// Network ID of the subnet
	SubnetId string `json:"subnet_id"`
	// Security group ID
	SecurityGroupId string `json:"security_group_id"`
	// ID of the parameter template used for creating an instance or ID of the latest parameter template that is applied to an instance.
	ConfigurationId string `json:"configuration_id"`
	// Backup policy
	BackupStrategy MysqlBackupStrategy `json:"backup_strategy"`
	// Node information
	Nodes []MysqlInstanceNodeInfo `json:"nodes"`
	// Time zone
	TimeZone string `json:"time_zone"`
	// AZ type. It can be single or multi.
	AzMode string `json:"az_mode"`
	// Primary AZ
	MasterAzCode string `json:"master_az_code"`
	// Maintenance window in the UTC format
	MaintenanceWindow string `json:"maintenance_window"`
	// Tags for managing instances
	Tags []MysqlTags `json:"tags"`
	// Dedicated resource pool ID. This parameter is returned only when the instance belongs to a dedicated resource pool.
	DedicatedResourceId string `json:"dedicated_resource_id"`
}

type MysqlInstanceNodeInfo struct {
	// Instance ID
	Id string `json:"id"`
	// Node name
	Name string `json:"name"`
	// Node type, which can be master or slave.
	Type string `json:"type"`
	// Node status
	Status string `json:"status"`
	// Database port
	Port int32 `json:"port"`
	// Private IP address for read of the node
	PrivateReadIps []string `json:"private_read_ips"`
	// Storage disk information
	Volume MysqlInstanceNodeVolumeInfo `json:"volume"`
	// AZ
	AzCode string `json:"az_code"`
	// Region where the instance is located
	RegionCode string `json:"region_code"`
	// Creation time yyyy-mm-ddThh:mm:ssZ
	Created string `json:"created"`
	// Update time
	Updated string `json:"updated"`
	// Specification code
	FlavorRef string `json:"flavor_ref"`
	// Maximum number of connections
	MaxConnections string `json:"max_connections"`
	// Number of vCPUs
	Vcpus string `json:"vcpus"`
	// Memory size in GB
	Ram string `json:"ram"`
	// Whether to reboot the instance for the parameter modifications to take effect.
	NeedRestart bool `json:"need_restart"`
	// Failover priority
	Priority int32 `json:"priority"`
}

type MysqlInstanceNodeVolumeInfo struct {
	// Disk type
	Type string `json:"type"`
	// Used disk size in GB
	Used string `json:"used"`
}
