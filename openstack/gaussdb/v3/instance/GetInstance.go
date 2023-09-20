package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetInstance(client *golangsdk.ServiceClient, id string) (*GetInstanceInfo, error) {
	// GET https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}
	raw, err := client.Get(client.ServiceURL("instances", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		Instance GetInstanceInfo `json:"instance"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Instance, err
}

type GetInstanceInfo struct {
	// Instance ID
	Id string `json:"id"`
	// Project ID of a tenant in a region
	ProjectId string `json:"project_id"`
	// DB instance remarks
	Alias string `json:"alias"`
	// The number of nodes.
	NodeCount int `json:"node_count"`
	// Instance name
	Name string `json:"name"`
	// Instance status
	Status string `json:"status"`
	// Private IP address for write It is a blank string until an ECS is created.
	PrivateIps []string `json:"private_ips"`
	// Public IP address string
	PublicIps string `json:"public_ips"`
	// Database port
	Port string `json:"port"`
	// Instance type. The value is Cluster.
	Type string `json:"type"`
	// Region where the instance is deployed
	Region string `json:"region"`
	// Database information
	Datastore Datastore `json:"datastore"`
	// Used backup space in GB
	BackupUsedSpace float64 `json:"backup_used_space"`
	// Creation time in the "yyyy-mm-ddThh:mm:ssZ" format.
	// T is the separator between the calendar and the hourly notation of time.
	// Z indicates the time zone offset. For example, for French Winter Time (FWT), the time offset is shown as +0200.
	// The value is empty unless the instance creation is complete.
	Created string `json:"created"`
	// Update time. The format is the same as that of the created field.
	// The value is empty unless the instance creation is complete.
	Updated string `json:"updated"`
	// Private IP address for write
	PrivateWriteIps []string `json:"private_write_ips"`
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
	// Specification code
	FlavorRef string `json:"flavor_ref"`
	// Specification description
	FlavorInfo FlavorInfo `json:"flavor_info"`
	// Node information
	Nodes *[]NodeInfo `json:"nodes"`
	// AZ type. It can be single or multi.
	AzMode string `json:"az_mode"`
	// Primary AZ
	MasterAzCode string `json:"master_az_code"`
	// Maintenance window in the UTC format
	MaintenanceWindow string `json:"maintenance_window"`
	// Storage disk information
	Volume VolumeInfo `json:"volume"`
	// Backup policy
	BackupStrategy BackupStrategy `json:"backup_strategy"`
	// Time zone
	TimeZone string `json:"time_zone"`
	// Billing mode, which is yearly/monthly or pay-per-use (default setting).
	ChargeInfo ChargeInfo `json:"charge_info"`
	// Dedicated resource pool ID. This parameter is returned only when the instance belongs to a dedicated resource pool.
	DedicatedResourceId string `json:"dedicated_resource_id"`
	// Tag list
	Tags []TagItem `json:"tags"`
	// Proxy information
	Proxies *[]Proxies `json:"proxies"`
}

type Proxies struct {
	PoolId  string `json:"pool_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type NodeInfo struct {
	// Instance ID
	Id string `json:"id"`
	// Node name
	Name string `json:"name"`
	// Node type, which can be master or slave.
	Type string `json:"type"`
	// Node status
	Status string `json:"status"`
	// Database port
	Port int `json:"port"`
	// Private IP address for read of the node
	PrivateReadIps []string `json:"private_read_ips"`
	// Storage disk information
	Volume *NodeVolumeInfo `json:"volume"`
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
	Priority int `json:"priority"`
}

type NodeVolumeInfo struct {
	// Disk type
	Type string `json:"type"`
	// Used disk size in GB
	Used string `json:"used"`
}
