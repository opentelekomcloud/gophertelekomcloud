package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

func ListClusters(client *golangsdk.ServiceClient) (*ListClustersResponse, error) {
	// GET /v1.0/{project_id}/clusters
	raw, err := client.Get(client.ServiceURL("clusters"), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res ListClustersResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListClustersResponse struct {
	// List of cluster objects
	Clusters []ClusterInfo `json:"clusters,omitempty"`
	// Total number of cluster objects
	Count int `json:"count,omitempty"`
}

type ClusterInfo struct {
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
	// Cluster creation time. Format: ISO8601:YYYY-MM-DDThh:mm:ssZ
	Created string `json:"created"`
	// Service port of a cluster. The value ranges from 8000 to 30000. The default value is 8000.
	Port int `json:"port"`
	// Private network connection information about the cluster.
	Endpoints []Endpoints `json:"endpoints"`
	// Unused
	Nodes []Nodes `json:"nodes"`
	// Tags in a cluster
	Tags []Tags `json:"tags"`
	// Cluster administrator name
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
	// Public network connection information about the cluster.
	// If the parameter is not specified, the public network connection information is not used by default.
	PublicEndpoints []PublicEndpoints `json:"public_endpoints"`
	// Task information, consisting of a key and a value. The key indicates an ongoing task, and the value indicates the progress of the ongoing task.
	// Valid key values include:
	// GROWING
	// RESTORING
	// SNAPSHOTTING
	// REPAIRING
	// CREATING
	// The value indicates the task progress.
	// Example:
	// "action_progress":
	// {"SNAPSHOTTING":"16%"}
	ActionProgress map[string]string `json:"action_progress"`
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
	// Security group ID
	SecurityGroupId string `json:"security_group_id"`
	// Cause of failure. If the parameter is left empty, the cluster is in the normal state.
	FailedReasons FailedReason `json:"failed_reasons,omitempty"`
}
