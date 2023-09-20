package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListInstancesOpts struct {
	// Instance ID The asterisk (*) is reserved for the system. If the instance ID starts with *,
	// it indicates that fuzzy match is performed based on the value following *.
	// Otherwise, the exact match is performed based on the instance ID. The value cannot contain only asterisks (*).
	Id string `json:"id,omitempty"`
	// Instance name The asterisk (*) is reserved for the system. If the instance name starts with *,
	// it indicates that fuzzy match is performed based on the value following *.
	// Otherwise, the exact match is performed based on the instance name. The value cannot contain only asterisks (*).
	Name string `json:"name,omitempty"`
	// Instance type to be queried. Currently, its value can only be Cluster.
	Type string `json:"type,omitempty"`
	// DB type. Currently, only gaussdb-mysql is supported.
	DatastoreType string `json:"datastore_type,omitempty"`
	// VPC ID
	// Method 1: Log in to the VPC console and view the VPC ID on the VPC details page.
	// Method 2: See section "Querying VPCs" in the Virtual Private Cloud API Reference.
	VpcId string `json:"vpc_id,omitempty"`
	// Network ID of the subnet
	// Method 1: Log in to the VPC console and click the target subnet on the Subnets page to view the network ID on the displayed page.
	// Method 2: See section "Querying Subnets" under "APIs" or section "Querying Networks" under "OpenStack Neutron APIs" in the Virtual Private Cloud API Reference.
	SubnetId string `json:"subnet_id,omitempty"`
	// Private IP address
	PrivateIp string `json:"private_ip,omitempty"`
	// Index offset. If offset is set to N, the resource query starts from the N+1 piece of data.
	// The default value is 0, indicating that the query starts from the first piece of data. The value must be a positive integer.
	Offset int `json:"offset,omitempty"`
	// Number of records to be queried. The default value is 100. The value must be a positive integer.
	// The minimum value is 1 and the maximum value is 100.
	Limit int `json:"limit,omitempty"`
	// Query based on the instance tag key and value. {key} indicates the tag key, and {value} indicates the tag value.
	// To query instances with multiple tag keys and values, separate key-value pairs with commas (,).
	// The key must be unique. Multiple keys are in AND relationship.
	Tags string `json:"tags,omitempty"`
}

func ListInstances(client *golangsdk.ServiceClient, opts ListInstancesOpts) (*ListInstancesResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/mysql/v3/{project_id}/instances
	raw, err := client.Get(client.ServiceURL("instances")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListInstancesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListInstancesResponse struct {
	// Instance list information.
	Instances []InstancesListInfo `json:"instances"`
	// Total number of records.
	TotalCount int `json:"total_count"`
}

type InstancesListInfo struct {
	// Instance ID
	Id string `json:"id"`
	// Project ID of a tenant in a region
	ProjectId string `json:"project_id"`
	// The number of nodes.
	NodeCount int `json:"node_count"`
	// Instance name
	Name string `json:"name"`
	// Instance status
	Status string `json:"status"`
	// Private IP address for write It is a blank string until an ECS is created.
	PrivateIps []string `json:"private_ips"`
	// Public IP address list
	PublicIps []string `json:"public_ips"`
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
}

type FlavorInfo struct {
	// Number of vCPUs
	Vcpus string `json:"vcpus"`
	// Memory size in GB
	Ram string `json:"ram"`
}

type VolumeInfo struct {
	// Disk type
	Type string `json:"type"`
	// Used disk size in GB
	Size string `json:"size"`
}

type TagItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
