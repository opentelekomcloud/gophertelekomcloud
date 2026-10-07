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
	// Extended scaling groups with node flavors or AZs that differ from the default node template.
	ExtensionScaleGroups []ExtensionScaleGroup `json:"extensionScaleGroups"`
	// Custom security group settings for a node pool
	CustomSecurityGroupIds []string `json:"customSecurityGroups,omitempty"`
}

// ExtensionScaleGroup describes an additional scaling group in a node pool.
// The default scaling group is described by Spec.NodeTemplate.
type ExtensionScaleGroup struct {
	// Basic information about the extended scaling group.
	Metadata *ExtensionScaleGroupMetadata `json:"metadata,omitempty"`
	// Configuration that differs from the default scaling group.
	Spec *ExtensionScaleGroupSpec `json:"spec,omitempty"`
}

// ExtensionScaleGroupMetadata contains the identity of an extended scaling group.
type ExtensionScaleGroupMetadata struct {
	// UID is generated when the scaling group is created. When updating a node
	// pool, it identifies the scaling group to update.
	Uid string `json:"uid,omitempty"`
	// Name of the extended scaling group. The name cannot be "default".
	Name string `json:"name,omitempty"`
}

// ExtensionScaleGroupSpec describes the node flavor, AZ, and scaling behavior
// of an extended scaling group.
type ExtensionScaleGroupSpec struct {
	// Node flavor for the extended scaling group.
	Flavor string `json:"flavor,omitempty"`
	// Availability zone. If omitted, the default scaling group AZ is used.
	Az string `json:"az,omitempty"`
	// Capacity reservation configuration for the extended scaling group.
	CapacityReservationSpecification *CapacityReservationSpecification `json:"capacityReservationSpecification,omitempty"`
	// Auto scaling configuration for the extended scaling group.
	Autoscaling *ScaleGroupAutoscaling `json:"autoscaling,omitempty"`
}

// CapacityReservationSpecification describes whether an extended scaling
// group uses capacity from a private pool.
type CapacityReservationSpecification struct {
	// ID of the private pool. It is required when Preference is "targeted".
	ID string `json:"id,omitempty"`
	// Capacity reservation preference. Supported values are "none" and "targeted".
	Preference string `json:"preference,omitempty"`
}

// ScaleGroupAutoscaling describes auto scaling for an extended scaling group.
type ScaleGroupAutoscaling struct {
	// Whether auto scaling is enabled for the scaling group.
	Enable bool `json:"enable,omitempty"`
	// Scaling group priority. A larger value indicates a higher priority.
	ExtensionPriority int `json:"extensionPriority,omitempty"`
	// Minimum number of nodes retained during auto scaling.
	MinNodeCount int `json:"minNodeCount,omitempty"`
	// Maximum number of nodes retained during auto scaling.
	MaxNodeCount int `json:"maxNodeCount,omitempty"`
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
