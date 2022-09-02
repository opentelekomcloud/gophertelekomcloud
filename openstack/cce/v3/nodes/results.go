package nodes

import (
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// Nodes of the cluster
type Nodes struct {
	//  API type, fixed value " Host "
	Kind string `json:"kind"`
	// API version, fixed value v3
	Apiversion string `json:"apiVersion"`
	// Node metadata
	Metadata Metadata `json:"metadata"`
	// Node detailed parameters
	Spec Spec `json:"spec"`
	// Node status information
	Status Status `json:"status"`
}

// Metadata required to create a node
type Metadata struct {
	// Node name
	Name string `json:"name"`
	// Node ID
	Id string `json:"uid"`
	// Node tag, key value pair format
	Labels map[string]string `json:"labels,omitempty"`
	// Node annotation, key/value pair format
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Spec describes Nodes specification
type Spec struct {
	// Node specifications
	Flavor string `json:"flavor" required:"true"`
	// The value of the available partition name
	Az string `json:"az" required:"true"`
	// The OS of the node
	Os string `json:"os,omitempty"`
	// ID of the dedicated host to which nodes will be scheduled
	DedicatedHostID string `json:"dedicatedHostId,omitempty"`
	// Node login parameters
	Login LoginSpec `json:"login" required:"true"`
	// System disk parameter of the node
	RootVolume VolumeSpec `json:"rootVolume" required:"true"`
	// The data disk parameter of the node must currently be a disk
	DataVolumes []VolumeSpec `json:"dataVolumes" required:"true"`
	// Elastic IP parameters of the node
	PublicIP PublicIPSpec `json:"publicIP,omitempty"`
	// The billing mode of the node: the value is 0 (on demand)
	BillingMode int `json:"billingMode,omitempty"`
	// Number of nodes when creating in batch
	Count int `json:"count" required:"true"`
	// The node nic spec
	NodeNicSpec NodeNicSpec `json:"nodeNicSpec,omitempty"`
	// Extended parameter
	ExtendParam ExtendParam `json:"extendParam,omitempty"`
	// UUID of an ECS group
	EcsGroupID string `json:"ecsGroupId,omitempty"`
	// Tag of a VM, key value pair format
	UserTags []tags.ResourceTag `json:"userTags,omitempty"`
	// Tag of a Kubernetes node, key value pair format
	K8sTags map[string]string `json:"k8sTags,omitempty"`
	// taints to created nodes to configure anti-affinity
	Taints []TaintSpec `json:"taints,omitempty"`
}

// NodeNicSpec spec of the node
type NodeNicSpec struct {
	// The primary Nic of the Node
	PrimaryNic PrimaryNic `json:"primaryNic,omitempty"`
}

// PrimaryNic of the node
type PrimaryNic struct {
	// The Subnet ID of the primary Nic
	SubnetId string `json:"subnetId,omitempty"`

	// FixedIPs define list of private IPs
	FixedIPs []string `json:"fixedIps,omitempty"`
}

// TaintSpec to created nodes to configure anti-affinity
type TaintSpec struct {
	Key   string `json:"key" required:"true"`
	Value string `json:"value" required:"true"`
	// Available options are NoSchedule, PreferNoSchedule, and NoExecute
	Effect string `json:"effect" required:"true"`
}

// Status gives the current status of the node
type Status struct {
	// The state of the Node
	Phase string `json:"phase"`
	// The virtual machine ID of the node in the ECS
	ServerID string `json:"ServerID"`
	// Elastic IP of the node
	PublicIP string `json:"PublicIP"`
	// Private IP of the node
	PrivateIP string `json:"privateIP"`
	// The ID of the Job that is operating asynchronously in the Node
	JobID string `json:"jobID"`
	// Reasons for the Node to become current
	Reason string `json:"reason"`
	// Details of the node transitioning to the current state
	Message string `json:"message"`
	// The status of each component in the Node
	Conditions Conditions `json:"conditions"`
}

type LoginSpec struct {
	// Select the key pair name when logging in by key pair mode
	SshKey string `json:"sshKey,omitempty"`
	// Select the user/password when logging in
	UserPassword UserPassword `json:"userPassword,omitempty"`
}

type UserPassword struct {
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
}

type VolumeSpec struct {
	// Disk Size in GB
	Size int `json:"size" required:"true"`
	// Disk VolumeType
	VolumeType string `json:"volumetype" required:"true"`
	// Metadata contains data disk encryption information
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Disk extension parameter
	ExtendParam string `json:"extendParam,omitempty"`
}

type ExtendParam struct {
	// Node charging mode, 0 is on-demand charging.
	ChargingMode int `json:"chargingMode,omitempty"`
	// Classification of cloud server specifications.
	EcsPerformanceType string `json:"ecs:performancetype,omitempty"`
	// Order ID, mandatory when the node payment type is the automatic payment package period type.
	OrderID string `json:"orderID,omitempty"`
	// The Product ID.
	ProductID string `json:"productID,omitempty"`
	// The Public Key.
	PublicKey string `json:"publicKey,omitempty"`
	// The maximum number of instances a node is allowed to create.
	MaxPods int `json:"maxPods,omitempty"`
	// Script required before the installation.
	PreInstall string `json:"alpha.cce/preInstall,omitempty"`
	// Script required after the installation.
	PostInstall string `json:"alpha.cce/postInstall,omitempty"`
	// Whether auto-renew is enabled.
	IsAutoRenew *bool `json:"isAutoRenew,omitempty"`
	// Whether to deduct fees automatically.
	IsAutoPay *bool `json:"isAutoPay,omitempty"`
	// Available disk space of a single Docker container on the node using the device mapper.
	DockerBaseSize int `json:"dockerBaseSize,omitempty"`
	// ConfigMap of the Docker data disk.
	DockerLVMConfigOverride string `json:"DockerLVMConfigOverride,omitempty"`
}

type PublicIPSpec struct {
	// List of existing elastic IP IDs
	Ids []string `json:"ids,omitempty"`
	// The number of elastic IPs to be dynamically created
	Count int `json:"count,omitempty"`
	// Elastic IP parameters
	Eip EipSpec `json:"eip,omitempty"`
}

type EipSpec struct {
	// The value of the iptype keyword
	IpType string `json:"iptype,omitempty"`
	// Elastic IP bandwidth parameters
	Bandwidth BandwidthOpts `json:"bandwidth,omitempty"`
}

type BandwidthOpts struct {
	ChargeMode string `json:"chargemode,omitempty"`
	Size       int    `json:"size,omitempty"`
	ShareType  string `json:"sharetype,omitempty"`
}

type Conditions struct {
	// The type of component
	Type string `json:"type"`
	// The state of the component
	Status string `json:"status"`
	// The reason that the component becomes current
	Reason string `json:"reason"`
}
