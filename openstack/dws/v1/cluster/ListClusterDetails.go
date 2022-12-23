package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

func ListClusterDetails(client *golangsdk.ServiceClient, clusterId string) (*ClusterDetail, error) {
	// GET /v1.0/{project_id}/clusters/{cluster_id}
	raw, err := client.Get(client.ServiceURL("clusters", clusterId), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res ClusterDetail
	err = extract.IntoStructPtr(raw.Body, &res, "cluster")
	return &res, err
}

type ClusterDetail struct {
	// Cluster ID
	Id string `json:"id"`
	// Cluster name
	Name string `json:"name"`
	// Cluster status. The value can be one of the following:
	// CREATING
	// AVAILABLE
	// UNAVAILABLE
	// CREATION FAILED
	Status string `json:"status"`
	// Data warehouse version
	Version string `json:"version"`
	// Last modification time of a cluster. Format: ISO8601:YYYY-MM-DDThh:mm:ssZ
	Updated string `json:"updated"`
	// Cluster creation time. Format: ISO8601: YYYY-MM-DDThh:mm:ssZ
	Created string `json:"created"`
	// Service port of a cluster. The value ranges from 8000 to 30000. The default value is 8000.
	Port int `json:"port"`
	// Private network connection information about the cluster.
	Endpoints []Endpoints `json:"endpoints"`
	// Unused
	Nodes []Nodes `json:"nodes"`
	// Labels in a cluster
	Tags []Tags `json:"tags"`
	// Administrator name
	UserName string `json:"user_name"`
	// Number of nodes in a cluster. The value ranges from 2 to 256.
	NumberOfNode int `json:"number_of_node"`
	// Number of events
	RecentEvent int `json:"recent_event"`
	// AZ
	AvailabilityZone string `json:"availability_zone"`
	// Enterprise project ID. The value 0 indicates the ID of the default enterprise project.
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// Node type
	NodeType string `json:"node_type"`
	// VPC ID
	VpcId string `json:"vpc_id"`
	// Subnet ID
	SubnetId string `json:"subnet_id"`
	// Public IP address. If the parameter is not specified, public connection is not used by default.
	PublicIp PublicIp `json:"public_ip"`
	// Public network connection information about the cluster. If the parameter is not specified, the public network connection information is not used by default.
	PublicEndpoints []PublicEndpoints `json:"public_endpoints"`
	// The key indicates an ongoing task. The value can be one of the following:
	// GROWING
	// RESTORING
	// SNAPSHOTTING
	// REPAIRING
	// CREATING
	// The value indicates the task progress.
	ActionProgress map[string]interface{} `json:"action_progress"`
	// Sub-status of clusters in the AVAILABLE state. The value can be one of the following:
	// NORMAL
	// READONLY
	// REDISTRIBUTING
	// REDISTRIBUTION-FAILURE
	// UNBALANCED
	// UNBALANCED | READONLY
	// DEGRADED
	// DEGRADED | READONLY
	// DEGRADED | UNBALANCED
	// UNBALANCED | REDISTRIBUTING
	// UNBALANCED | REDISTRIBUTION-FAILURE
	// READONLY | REDISTRIBUTION-FAILURE
	// UNBALANCED | READONLY | REDISTRIBUTION-FAILURE
	// DEGRADED | REDISTRIBUTION-FAILURE
	// DEGRADED | UNBALANCED | REDISTRIBUTION-FAILURE
	// DEGRADED | UNBALANCED | READONLY | REDISTRIBUTION-FAILURE
	// DEGRADED | UNBALANCED | READONLY
	SubStatus string `json:"sub_status"`
	// Cluster management task. The value can be one of the following:
	// RESTORING
	// SNAPSHOTTING
	// GROWING
	// REBOOTING
	// SETTING_CONFIGURATION
	// CONFIGURING_EXT_DATASOURCE
	// DELETING_EXT_DATASOURCE
	// REBOOT_FAILURE
	// RESIZE_FAILURE
	TaskStatus string `json:"task_status"`
	// Parameter group details
	ParameterGroup ParameterGroup `json:"parameter_group,omitempty"`
	// Node type ID
	NodeTypeId string `json:"node_type_id"`
	// Security group ID
	SecurityGroupId string `json:"security_group_id"`
	// List of private network IP addresses
	PrivateIp []string `json:"private_ip"`
	// Cluster maintenance window
	MaintainWindow MaintainWindow `json:"maintain_window"`
	// Cluster scale-out details
	ResizeInfo ResizeInfo `json:"resize_info,omitempty"`
	// Cause of failure. If the parameter is left empty, the cluster is in the normal state.
	FailedReasons FailedReason `json:"failed_reasons,omitempty"`
}

type Endpoints struct {
	// Private network connection information
	ConnectInfo string `json:"connect_info,omitempty"`
	// JDBC URL on the private network. The following is the default format:
	//jdbc:postgresql://< connect_info>/<YOUR_DATABASE_name>
	JdbcUrl string `json:"jdbc_url,omitempty"`
}

type Nodes struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type Tags struct {
	// Key. A key can contain a maximum of 36 Unicode characters, which cannot be null.
	// The first and last characters cannot be spaces.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed. It cannot contain the following characters: =*<>\,|/
	Key string `json:"key"`
	// Value. A value can contain a maximum of 43 Unicode characters, which can be null. The first and last characters cannot be spaces.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed. It cannot contain the following characters: =*<>\,|/
	Value string `json:"value"`
}

type PublicEndpoints struct {
	// Public network connection information
	PublicConnectInfo string `json:"public_connect_info,omitempty"`
	// JDBC URL of the public network. The following is the default format:
	//jdbc:postgresql://< public_connect_info>/<YOUR_DATABASE_name>
	JdbcUrl string `json:"jdbc_url,omitempty"`
}

type ParameterGroup struct {
	// Parameter group ID
	Id string `json:"id"`
	// Parameter group name
	Name string `json:"name"`
	// Cluster parameter status. The value can be one of the following:
	// In-Sync: synchronized
	// Applying: in application
	// Pending-Reboot: restart for the modification to take effect
	// Sync-Failure: application failure
	Status string `json:"status"`
}

type MaintainWindow struct {
	// Maintenance time in each week in the unit of day. The value can be one of the following:
	// Mon
	// Tue
	// Wed
	// Thu
	// Fri
	// Sat
	// Sun
	Day string `json:"day,omitempty"`
	// Maintenance start time in HH:mm format. The time zone is GMT+0.
	StartTime string `json:"start_time,omitempty"`
	// Maintenance end time in HH:mm format. The time zone is GMT+0.
	EndTime string `json:"end_time,omitempty"`
}

type ResizeInfo struct {
	// Number of nodes after the scale-out
	TargetNodeNum int `json:"target_node_num,omitempty"`
	// Number of nodes before the scale-out
	OriginNodeNum int `json:"origin_node_num,omitempty"`
	// Scale-out status. The value can be one of the following:
	// GROWING
	// RESIZE_FAILURE
	ResizeStatus string `json:"resize_status,omitempty"`
	// Scale-out start time. Format: ISO8601:YYYY-MM-DDThh:mm:ss
	StartTime string `json:"start_time,omitempty"`
}

type FailedReason struct {
	// Error code
	ErrorCode string `json:"error_code,omitempty"`
	// Error message
	ErrorMsg string `json:"error_msg,omitempty"`
}
