package nodes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	tag "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type AcceptOpts struct {
	ClusterID string `json:"-"`
	// API type, fixed value List
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// List of nodes to be accepted
	NodeList []AddNode `json:"nodeList" required:"true"`
}

// AddNode contains the parameters for a single node to be accepted
type AddNode struct {
	// Server ID
	ServerID string `json:"serverID" required:"true"`
	// Node reinstallation configuration parameters
	Spec ReinstallNodeSpec `json:"spec" required:"true"`
}

// ReinstallNodeSpec contains the configuration for node reinstallation
type ReinstallNodeSpec struct {
	// Operating system
	OS string `json:"os" required:"true"`
	// Node login mode
	Login LoginSpec `json:"login" required:"true"`
	// Node name
	Name string `json:"name,omitempty"`
	// Server configuration
	ServerConfig *ReinstallServerConfig `json:"serverConfig,omitempty"`
	// Volume management configuration
	VolumeConfig *ReinstallVolumeConfig `json:"volumeConfig,omitempty"`
	// Container runtime configuration
	RuntimeConfig *ReinstallRuntimeConfig `json:"runtimeConfig,omitempty"`
	// Kubernetes node configuration
	K8sOptions *ReinstallK8sOptionsConfig `json:"k8sOptions,omitempty"`
	// Node lifecycle configuration
	Lifecycle *NodeLifecycleConfig `json:"lifecycle,omitempty"`
	// Custom initialization flags
	InitializedConditions []string `json:"initializedConditions,omitempty"`
	// Extended parameters
	ExtendParam *ReinstallExtendParam `json:"extendParam,omitempty"`
}

// ReinstallServerConfig contains server configuration parameters
type ReinstallServerConfig struct {
	// Cloud server labels
	UserTags []tag.ResourceTag `json:"userTags,omitempty"`
	// System disk configurations
	RootVolume *ReinstallVolumeSpec `json:"rootVolume,omitempty"`
}

// UserTag represents a key-value pair label for cloud servers
type UserTag struct {
	// Key of the cloud server label
	Key string `json:"key,omitempty"`
	// Value of the cloud server label
	Value string `json:"value,omitempty"`
}

// ReinstallVolumeSpec contains volume specifications
type ReinstallVolumeSpec struct {
	// Custom image ID
	ImageID string `json:"imageID,omitempty"`
	// User master key ID
	CmkID string `json:"cmkID,omitempty"`
}

// ReinstallVolumeConfig contains volume management configuration
type ReinstallVolumeConfig struct {
	// Docker data disk configurations
	LvmConfig string `json:"lvmConfig,omitempty"`
	// Disk initialization management
	Storage *Storage `json:"storage,omitempty"`
}

// ReinstallRuntimeConfig contains container runtime configuration
type ReinstallRuntimeConfig struct {
	// Available disk space of a single container
	DockerBaseSize int `json:"dockerBaseSize,omitempty"`
	// Container runtime
	Runtime *RuntimeSpec `json:"runtime,omitempty"`
}

// ReinstallK8sOptionsConfig contains Kubernetes configuration options
type ReinstallK8sOptionsConfig struct {
	// Node labels
	Labels map[string]string `json:"labels,omitempty"`
	// Node taints
	Taints []TaintSpec `json:"taints,omitempty"`
	// Maximum number of pods
	MaxPods int `json:"maxPods,omitempty"`
}

// NodeLifecycleConfig contains node lifecycle configuration
type NodeLifecycleConfig struct {
	// Pre-installation script
	PreInstall string `json:"preInstall,omitempty"`
	// Post-installation script
	PostInstall string `json:"postInstall,omitempty"`
}

// ReinstallExtendParam contains extended parameters
type ReinstallExtendParam struct {
	// ID of the user image
	AlphaCCENodeImageID string `json:"alpha.cce/NodeImageID,omitempty"`
}

// JobResult contains the response from accepting nodes
type JobResult struct {
	// Job ID for tracking the acceptance process
	JobID string `json:"jobid"`
}

// Accept sends a request to accept nodes into the specified cluster
func Accept(client *golangsdk.ServiceClient, opts AcceptOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /api/v3/projects/{project_id}/clusters/{cluster_id}/nodes/add
	raw, err := client.Post(client.ServiceURL("clusters", opts.ClusterID, "nodes", "add"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res JobResult
	err = extract.Into(raw.Body, &res)
	return res.JobID, err
}
