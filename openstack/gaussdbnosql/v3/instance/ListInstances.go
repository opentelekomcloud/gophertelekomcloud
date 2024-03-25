package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListInstancesOpts struct {
	// Instance ID
	// If you use asterisk (*) at the beginning of the id, fuzzy search results are returned.
	// Otherwise, the exact results are returned.
	Id string `q:"id,omitempty"`
	// Instance name
	// If you use asterisk (*) at the beginning of the name, fuzzy search results are returned.
	// Otherwise, the exact results are returned.
	Name string `q:"name,omitempty"`
	// Instance type
	// Cluster indicates the GaussDB(for Cassandra) cluster instance.
	// If the datastore_type parameter is not transferred, this parameter is automatically ignored.
	Mode string `q:"mode,omitempty"`
	// Database type
	// If the value is cassandra, GaussDB(for Cassandra) DB instances are queried.
	// If this parameter is not transferred, all DB instances are queried.
	DatastoreType string `q:"datastore_type,omitempty"`
	// VPC ID. To obtain this parameter value, use either of the following methods:
	// Method 1: Log in to VPC console and view the VPC ID on the VPC details page.
	// Method 2: See the "Querying VPCs" section in the Virtual Private Cloud API Reference.
	VpcId string `q:"vpc_id,omitempty"`
	// Network ID of the subnet. To obtain this parameter value, use either of the following methods:
	// Method 1: Log in to VPC console and click the target subnet on the Subnets page. You can view the subnet ID on the displayed page.
	// Method 2: See the "Querying Subnets" section in the Virtual Private Cloud API Reference.
	SubnetId string `q:"subnet_id,omitempty"`
	// Index position. The query starts from the next instance creation time indexed by this parameter under a specified project.
	// If offset is set to N, the resource query starts from the N+1 piece of data
	// The value must be greater than or equal to 0. If this parameter is not transferred,
	// offset is set to 0 by default, indicating that the query starts from the latest created instance.
	Offset int32 `q:"offset,omitempty"`
	// The maximum allowed number of instances.
	// The value ranges from 1 to 100. If this parameter is not transferred, the first 100 instances are queried by default.
	Limit int32 `q:"limit,omitempty"`
}

func ListInstances(client *golangsdk.ServiceClient, opts ListInstancesOpts) (*ListInstancesResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/instances
	raw, err := client.Get(client.ServiceURL("instances")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListInstancesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListInstancesResponse struct {
	Instances  []ListInstancesResult `json:"instances"`
	TotalCount int32                 `json:"total_count"`
}

type ListInstancesResult struct {
	// Instance ID
	Id string `json:"id"`
	// Instance name
	Name string `json:"name"`
	// Instance status
	// Valid value:
	// normal: indicates that the instance is running properly.
	// abnormal: indicates that the instance is abnormal.
	// creating: indicates that the instance is being created.
	// data_disk_full: indicates that the instance disk is full.
	// createfail: indicates that the instance failed to be created.
	// enlargefail: indicates that nodes failed to be added to the instance.
	Status string `json:"status"`
	// Database port
	Port string `json:"port"`
	// Instance type, which is the same as the request parameter.
	Mode string `json:"mode"`
	// Region where the instance is deployed.
	Region string `json:"region"`
	// Database information.
	Datastore InstancesDatastore `json:"datastore"`
	// Storage engine
	// The value is "rocksDB".
	Engine string `json:"engine"`
	// Instance creation time
	Created string `json:"created"`
	// The time when an instance is updated.
	Updated string `json:"updated"`
	// Default username. The value is rwuser.
	DbUserName string `json:"db_user_name"`
	// VPC ID
	VpcId string `json:"vpc_id"`
	// Subnet ID
	SubnetId string `json:"subnet_id"`
	// Security group ID
	SecurityGroupId string `json:"security_group_id"`
	// Backup policy.
	BackupStrategy InstancesBackupStrategy `json:"backup_strategy"`
	// The value is set to "0".
	PayMode string `json:"pay_mode"`
	// Maintenance time window
	MaintenanceWindow string `json:"maintenance_window"`
	// Group information.
	Groups []InstancesGroup `json:"groups"`
	// Time zone
	TimeZone string `json:"time_zone"`
	// The operation that is executed on the instance.
	Actions []string `json:"actions"`
}

type InstancesDatastore struct {
	// Database type
	Type string `json:"type"`
	// Database version
	Version string `json:"version"`
}

type InstancesBackupStrategy struct {
	// Backup time window. Automated backups will be triggered during the backup time window. The current time is the UTC time.
	StartTime string `json:"start_time"`
	// Backup retention days. Value range: 0-35.
	KeepDays int32 `json:"keep_days"`
}

type InstancesGroup struct {
	// Group ID
	Id string `json:"id"`
	// Group status
	Status string `json:"status"`
	// Volume information.
	Volume Volume `json:"volume"`
	// Node information
	Nodes []InstancesNode `json:"nodes"`
}

type Volume struct {
	// Storage space. Unit: GB
	Size string `json:"size"`
	// Storage space usage. Unit: GB
	Used string `json:"used"`
}

type InstancesNode struct {
	// Node ID
	Id string `json:"id"`
	// Node name
	Name string `json:"name"`
	// Node status
	Status string `json:"status"`
	// Private IP address of a node. This parameter value exists after the ECS is created. Otherwise, the value is "".
	PrivateIp string `json:"private_ip"`
	// Bound EIP. This parameter is valid only for nodes bound with EIPs.
	PublicIp string `json:"public_ip"`
	// Resource specification code
	// For details, see the value of the flavors.spec_code parameter in 5.3 Querying All Instance Specifications.
	SpecCode string `json:"spec_code"`
	// AvailabilityZone
	AvailabilityZone string `json:"availability_zone"`
	// Whether scaling in instances is supported.
	// true: indicates that the node scale-in is supported.
	// false: indicates that the node scale-in is not supported.
	SupportReduce bool `json:"support_reduce"`
}
