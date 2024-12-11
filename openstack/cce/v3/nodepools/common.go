package nodepools

import (
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/nodes"
)

// NodePool - Individual node pools of the cluster
type NodePool struct {
	// Node pool type
	Type string `json:"type" required:"true"`
	//  API type, fixed value " Host "
	Kind string `json:"kind"`
	// API version, fixed value v3
	Apiversion string `json:"apiVersion"`
	// Node Pool metadata
	Metadata Metadata `json:"metadata"`
	// Node Pool detailed parameters
	Spec Spec `json:"spec"`
	// Node Pool status information
	Status Status `json:"status"`
}

// Metadata of the node pool
type Metadata struct {
	// Node Pool name
	Name string `json:"name"`
	// Node Pool ID
	Id string `json:"uid"`
}

// Status - Gives the current status of the node pool
type Status struct {
	// The state of the node pool
	Phase string `json:"phase"`
	// Number of nodes in the node pool
	CurrentNode int `json:"currentNode"`
}

// Spec describes Node pools specification
type Spec struct {
	// Node type. Currently, only VM nodes are supported.
	Type string `json:"type" required:"true"`
	// Node Pool template
	NodeTemplate nodes.Spec `json:"nodeTemplate" required:"true"`
	// Initial number of expected node pools
	InitialNodeCount int `json:"initialNodeCount" required:"true"`
	// Auto scaling parameters
	Autoscaling AutoscalingSpec `json:"autoscaling"`
	// Node pool management parameters
	NodeManagement NodeManagementSpec `json:"nodeManagement"`
	// Custom security group settings for a node pool
	CustomSecurityGroupIds []string `json:"customSecurityGroups,omitempty"`
}

type AutoscalingSpec struct {
	// Whether to enable auto scaling
	Enable bool `json:"enable"`
	// Minimum number of nodes allowed if auto scaling is enabled
	MinNodeCount int `json:"minNodeCount"`
	// This value must be greater than or equal to the value of minNodeCount
	MaxNodeCount int `json:"maxNodeCount"`
	// Interval between two scaling operations, in minutes
	ScaleDownCooldownTime int `json:"scaleDownCooldownTime"`
	// Weight of a node pool
	Priority int `json:"priority"`
}

type NodeManagementSpec struct {
	// ECS group ID
	ServerGroupReference string `json:"serverGroupReference"`
}
