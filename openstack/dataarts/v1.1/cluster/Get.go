package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get is used to query cluster details.
// Send request GET /v1.1/{project_id}/clusters/{cluster_id}
func Get(client *golangsdk.ServiceClient, clusterId string) (*ClusterQuery, error) {
	raw, err := client.Get(client.ServiceURL(clustersEndpoint, clusterId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ClusterQuery
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ClusterQuery struct {
	// EIP bound to the cluster
	PublicEndpoint string `json:"publicEndpoint"`
	// Cluster node information. For details, see the descriptions of instances parameters.
	Instances []DetailedInstances `json:"instances"`
	// Security group ID
	SecurityGroupId string `json:"security_group_id"`
	// Subnet ID
	SubnetId string `json:"subnet_id"`
	// VPC ID
	VpcId string `json:"vpc_id"`
	// User configuration
	CustomerConfig CustomerConfig `json:"customerConfig"`
	// CDM information
	Datastore Datastore `json:"datastore"`
	// Auto shutdown
	IsAutoOff bool `json:"isAutoOff"`
	// Domain name for the EIP bound to the cluster
	PublicEndpointDomainName string `json:"publicEndpointDomainName"`
	// Start time
	BakExpectedStartTime string `json:"bakExpectedStartTime"`
	// Retention duration
	BakKeepDay string `json:"bakKeepDay"`
	// Maintenance window
	MaintainWindow MaintainWindow `json:"maintainWindow"`
	// Number of events
	RecentEvent int `json:"recentEvent"`
	// Flavor name
	FlavorName string `json:"flavorName"`
	// AZ name
	AzName string `json:"azName"`
	// Peer domain name
	EndpointDomainName string `json:"endpointDomainName"`
	// EIP status
	PublicEndpointStatus PublicEndpointStatus `json:"publicEndpointStatus"`
	// Whether to enable scheduled startup/shutdown. The scheduled startup/shutdown and auto shutdown functions cannot be enabled at the same time.
	IsScheduleBootOff bool `json:"isScheduleBootOff"`
	// Namespace
	Namespace string `json:"namespace"`
	// EIP ID
	EipId string `json:"eipId"`
	// Failure cause. If this parameter is left empty, the cluster is in normal state.
	FailedReasons FailedReasons `json:"failedReasons"`
	// Database user
	DbUser string `json:"dbuser"`
	// Cluster link information
	Links []ClusterLinks `json:"links"`
	// Cluster mode: sharding
	ClusterMode string `json:"clusterMode"`
	// Task information
	Task ClusterTask `json:"task"`
	// Cluster creation time in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	Created string `json:"created"`
	// Cluster status: normal
	StatusDetail string `json:"statusDetail"`
	// Cluster configuration status
	//    In-Sync: The configuration has been synchronized.
	//    Applying: The configuration is in progress.
	//    Sync-Failure: The configuration fails.
	ConfigStatus string `json:"config_status"`
	// Cluster operation progress, which consists of a key and a value.
	// The key indicates an ongoing task, and the value indicates the progress of the ongoing task.
	// An example is "action_progress":{"SNAPSHOTTING":"16%"}.
	ActionProgress ActionProgress `json:"actionProgress"`
	// Cluster name
	Name string `json:"name"`
	// Cluster ID
	Id string `json:"id"`
	// Whether the cluster is frozen. The value can be 0 (not frozen) or 1 (frozen).
	IsFrozen string `json:"isFrozen"`
	// Cluster configuration status. Options: - In-Sync: The cluster configuration has been synchronized. - Applying: The cluster is being configured. - Sync-Failure: The cluster configuration failed.
	Actions []string `json:"actions"`
	// Cluster update time in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	Updated string `json:"updated"`
	// Cluster status
	//    100: creating
	//    200: normal
	//    300: failed
	//    303: failed to be created
	//    800: frozen
	//    900: stopped
	//    910: stopping
	//    920: starting
	Status string `json:"status"`
}

type DetailedInstances struct {
	// VM flavor of a node. For details, see the descriptions of flavor parameters.
	Flavor Flavor `json:"flavor"`
	// Disk information of a node. For details, see the descriptions of volume parameters.
	Volume Volume `json:"volume"`
	// Node status
	//    100: creating
	//    200: normal
	//    300: failed
	//    303: failed to be created
	//    400: deleted
	//    800: frozen
	Status string `json:"status"`
	// Node operation status
	//    REBOOTING: restarting
	//    RESTORING: restoring
	//    REBOOT_FAILURE: failed to restart
	Actions []string `json:"actions"`
	// Node type. Currently, only cdm is available.
	Type string `json:"type"`
	// Node VM ID
	Id string `json:"id"`
	// Name of the VM on the node
	Name string `json:"name"`
	// Whether the node is frozen. The value can be 0 (not frozen) or 1 (frozen).
	IsFrozen string `json:"isFrozen"`
	// Component
	Components string `json:"components"`
	// Node configuration status. The value is null when the cluster list is queried.
	//    In-Sync: The configuration has been synchronized.
	//    Applying: The configuration is in progress.
	//    Sync-Failure: The configuration fails.
	ConfigStatus string `json:"config_status"`
	// Instance role
	Role string `json:"role"`
	// Group
	Group string `json:"group"`
	// Link information
	Links []ClusterLinks `json:"links"`
	// Group ID
	ParamsGroupId string `json:"paramsGroupId"`
	// Public IP address
	PublicIp string `json:"publicIp"`
	// Management IP address
	ManageIp string `json:"manageIp"`
	// Traffic IP address
	TrafficIp string `json:"trafficIp"`
	// Slice ID
	ShardId string `json:"shard_id"`
	// Management fix IP address
	ManageFixIp string `json:"manage_fix_ip"`
	// Private IP address
	PrivateIp string `json:"private_ip"`
	// Internal IP address
	InternalIp string `json:"internal_ip"`
	// Resource information (null is returned for querying the cluster list)
	Resource []Resource `json:"resource"`
}

type Flavor struct {
	// VM flavor ID
	Id string `json:"id"`
	// Link information
	Links []ClusterLinks `json:"links"`
}

type Volume struct {
	// Type of disks on the node. Only local disks are supported.
	Type string `json:"type"`
	// Size of the disk on the node (GB)
	Size int64 `json:"size"`
}

type Resource struct {
	// Resource ID
	ResourceId string `json:"resource_id"`
	// Resource type: server
	ResourceType string `json:"resource_type"`
}

type CustomerConfig struct {
	// Failure notification
	FailureRemind string `json:"failureRemind"`
	// Cluster type
	ClusterName string `json:"clusterName"`
	// Service provisioning
	ServiceProvider string `json:"serviceProvider"`
	// Whether the disk is a local disk
	LocalDisk string `json:"localDisk"`
	// Whether to enable SSL
	Ssl string `json:"ssl"`
	// Source
	CreateFrom string `json:"createFrom"`
	// Resource ID
	ResourceId string `json:"resourceId"`
	// Flavor type
	FlavorType string `json:"flavorType"`
	// Workspace ID
	WorkSpaceId string `json:"workSpaceId"`
	// Trial
	Trial string `json:"trial"`
}

type MaintainWindow struct {
	// Day of a week
	Day string `json:"day"`
	// Start time
	StartTime string `json:"startTime"`
	// End time
	EndTime string `json:"endTime"`
}

type PublicEndpointStatus struct {
	// Status
	Status string `json:"status"`
	// Error message
	ErrorMessage string `json:"errorMessage"`
}

type FailedReasons struct {
	// Cause of the cluster creation failure
	CreateFailed CreateFailed `json:"CREATE_FAILED"`
}

type CreateFailed struct {
	// Error code
	ErrorCode string `json:"errorCode"`
	// Failure cause
	ErrorMsg string `json:"errorMsg"`
}

type ClusterLinks struct {
	// Relationship
	Rel string `json:"rel"`
	// Link address
	Href string `json:"href"`
}

type ClusterTask struct {
	// Task description
	Description string `json:"description"`
	// Task ID
	Id string `json:"id"`
	// Task name
	Name string `json:"name"`
}

type ActionProgress struct {
	// Cluster creation progress, for example, 29%
	Creating string `json:"creating"`
	// Cluster expansion progress, for example, 29%
	Growing string `json:"growing"`
	// Cluster restoration progress, for example, 29%
	Restoring string `json:"restoring"`
	// Cluster snapshotting progress, for example, 29%
	Snapshotting string `json:"snapshotting"`
	// Cluster repairing progress, for example, 29%
	Repairing string `json:"repairing"`
}
