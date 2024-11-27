package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// AddExistingNodeOpts defines the structure of the API request body.
type AddExistingNodeOpts struct {
	// API version. The value is fixed at v3.
	APIVersion string `json:"apiVersion" required:"true"`
	// API type. The value is fixed at List.
	Kind string `json:"kind" required:"true"`
	// List of the nodes to be accepted.
	// Nodes must have 2-core or higher CPU, 4 GB or larger memory.
	NodeList []AddNode `json:"nodeList" required:"true"`
}

// AddNode defines the parameters for a node to be added.
type AddNode struct {
	// Server ID. For details about how to obtain the server ID, see the ECS or BMS documentation.
	ServerID string `json:"serverID" required:"true"`
	// Node reinstallation configuration parameters.
	// Currently, accepted nodes cannot be added into node pools.
	Spec *ReinstallNodeSpec `json:"spec" required:"true"`
}

// ReinstallNodeSpec defines the parameters for node reinstallation.
type ReinstallNodeSpec struct {
	// Operating system.
	// If you specify a custom image, the actual OS version in the IMS image is used.
	// Select an OS version supported by the current cluster, for example, EulerOS 2.5, EulerOS 2.9, Ubuntu 22.04, or HCE OS 2.0.
	OS string `json:"os" required:"true"`
	// Node login mode.
	Login Login `json:"login" required:"true"`
	// Node name.
	// Specifying this field during reinstallation will change the node name, and the server name will change accordingly.
	// By default, the current server name is used as the node name. Enter 1 to 56 characters starting with a letter
	// and not ending with a hyphen (-). Only lowercase letters, digits, and hyphens (-) are allowed.
	Name string `json:"name,omitempty"`
	// Server configuration.
	ServerConfig *ReinstallServerConfig `json:"serverConfig,omitempty"`
	// Volume management configuration.
	VolumeConfig *ReinstallVolumeConfig `json:"volumeConfig,omitempty"`
	// Container runtime configuration.
	RuntimeConfig *ReinstallRuntimeConfig `json:"runtimeConfig,omitempty"`
	// Kubernetes node configuration.
	K8sOptions *ReinstallK8sOptionsConfig `json:"k8sOptions,omitempty"`
	// Customized lifecycle configuration of a node.
	Lifecycle *NodeLifecycleConfig `json:"lifecycle,omitempty"`
	// Custom initialization flags.
	// Before CCE nodes are initialized, they are tainted with `node.cloudprovider.kubernetes.io/Uninitialized`
	// to prevent pods from being scheduled to them. Maximum 20 characters and 2 flags allowed.
	InitializedConditions []string `json:"initializedConditions,omitempty"`
	// Extended reinstallation parameter, which is discarded.
	ExtendParam *ReinstallExtendParam `json:"extendParam,omitempty"`
}

// Login defines the parameters for node login mode.
type Login struct {
	// Name of the key pair used for login.
	SSHKey string `json:"sshKey" required:"true"`
	// Password used for node login.
	// This field is not supported for the current version.
	UserPassword string `json:"userPassword,omitempty"`
}

// ReinstallServerConfig defines the server configuration for reinstallation.
type ReinstallServerConfig struct {
	// Cloud server labels. The key of a label must be unique.
	// The maximum number of user-defined labels supported by CCE depends on the region.
	// In the region that supports the least number of labels, you can still create up to 5 labels for a cloud server.
	UserTags []UserTag `json:"userTags,omitempty"`
	// System disk configurations used in reinstallation.
	RootVolume *ReinstallVolumeSpec `json:"rootVolume,omitempty"`
}

// UserTag defines the key-value pair for a cloud server label.
type UserTag struct {
	// Key of the cloud server label.
	// The value cannot start with `CCE-` or `__type_baremetal`.
	Key string `json:"key,omitempty"`
	// Value of the cloud server label.
	Value string `json:"value,omitempty"`
}

// ReinstallVolumeSpec defines the volume specification for reinstallation.
type ReinstallVolumeSpec struct {
	// Custom image ID.
	ImageID string `json:"imageID,omitempty"`
	// User master key ID.
	// If this parameter is left blank by default, the EVS disk is not encrypted.
	CmkID string `json:"cmkID,omitempty"`
}

// ReinstallVolumeConfig defines the configuration for Docker data disk and disk initialization.
type ReinstallVolumeConfig struct {
	// Docker data disk configurations.
	// Example default configuration:
	// "lvmConfig":"dockerThinpool=vgpaas/90%VG;kubernetesLV=vgpaas/10%VG;diskType=evs;lvType=linear"
	// Fields included:
	// - userLV: size of the user space, e.g., vgpaas/20%VG.
	// - userPath: mount path of the user space, e.g., /home/wqt-test.
	// - diskType: disk type. Supported values: evs, hdd, and ssd.
	// - lvType: type of a logical volume. Values: linear or striped.
	// - dockerThinpool: Docker space size, e.g., vgpaas/60%VG.
	// - kubernetesLV: kubelet space size, e.g., vgpaas/20%VG.
	LvmConfig string `json:"lvmConfig,omitempty"`
	// Disk initialization management parameter.
	// This parameter is complex to configure. For details, see Attaching Disks to a Node.
	// If not specified, disks are managed based on the DockerLVMConfigOverride (discarded) parameter in extendParam.
	// Supported by clusters of version 1.15.11 and later.
	Storage Storage `json:"storage,omitempty"`
}

// Storage defines the disk initialization management parameters.
type Storage struct {
	// Disk selection. Matched disks are managed according to matchLabels and storageType.
	StorageSelectors []StorageSelector `json:"storageSelectors" required:"true"`
	// A storage group consists of multiple storage devices.
	// It is used to divide storage space.
	StorageGroups []StorageGroup `json:"storageGroups" required:"true"`
}

// StorageSelector defines the parameters for disk selection and matching.
type StorageSelector struct {
	// Selector name, used as the index of selectorNames in storageGroup.
	// The name of each selector must be unique.
	Name string `json:"name" required:"true"`
	// Specifies the storage type. Currently, only evs (EVS volumes) and local (local volumes) are supported.
	// The local storage does not support disk selection. All local disks will form a VG.
	// Therefore, only one storageSelector of the local type is allowed.
	StorageType string `json:"storageType" required:"true"`
	// Matching field of an EVS volume. The size, volumeType, metadataEncrypted, metadataCmkid, and count fields are supported.
	MatchLabels *MatchLabels `json:"matchLabels,omitempty"`
}

// MatchLabels defines the matching criteria for EVS disks.
type MatchLabels struct {
	// Matched disk size. If this parameter is left unspecified, the disk size is not limited.
	// Example: "100".
	Size string `json:"size,omitempty"`
	// EVS disk type.
	VolumeType string `json:"volumeType,omitempty"`
	// Disk encryption identifier. "0" indicates that the disk is not encrypted,
	// and "1" indicates that the disk is encrypted.
	MetadataEncrypted string `json:"metadataEncrypted,omitempty"`
	// Customer master key ID of an encrypted disk. The value is a 36-byte string.
	MetadataCmkID string `json:"metadataCmkid,omitempty"`
	// Number of disks to be selected. If this parameter is left blank, all disks of this type are selected.
	Count string `json:"count,omitempty"`
}

// StorageGroups represents a virtual storage group configuration.
type StorageGroup struct {
	// Name of a virtual storage group, which must be unique.
	Name string `json:"name" required:"true"`

	// Storage space for Kubernetes and runtime components.
	// Only one group can be set to true. Defaults to false if not specified.
	CceManaged *bool `json:"cceManaged,omitempty"`

	// This parameter corresponds to "name" in storageSelectors.
	// A group can match multiple selectors, but a selector can match only one group.
	SelectorNames []string `json:"selectorNames" required:"true"`

	// Detailed management of space configuration in a group.
	VirtualSpaces []VirtualSpace `json:"virtualSpaces" required:"true"`

	// Number of disks to be selected. If this parameter is left blank, all disks of this type are selected.
	Count string `json:"count,omitempty"`
}

// VirtualSpace represents the configuration for virtual space management.
type VirtualSpace struct {
	// Name of a virtualSpace.
	// Kubernetes: Kubernetes space configuration. lvmConfig needs to be configured.
	// runtime: runtime space configuration. runtimeConfig needs to be configured.
	// user: user space configuration. lvmConfig needs to be configured
	Name string `json:"name" required:"true"`
	// Size of a virtualSpace. The value must be an integer in percentage. Example: 90%.
	Size string `json:"size" required:"true"`
	// LVM configurations, applicable to kubernetes and user spaces. Note that one virtual space supports only one config.
	LvmConfig LvmConfig `json:"lvmConfig,omitempty"`
	// Runtime configurations, applicable to the runtime space. Note that one virtual space supports only one config.
	RuntimeConfig RuntimeConfig `json:"runtimeConfig,omitempty"`
}

// LVMConfig represents the configuration for Logical Volume Management (LVM).
type LvmConfig struct {
	// LVM write mode. "linear" indicates the linear mode.
	// "striped" indicates the striped mode, in which multiple disks are used
	// to form a strip to improve disk performance.
	LvType string `json:"lvType" required:"true"`

	// Path to which the disk is attached. This parameter takes effect only
	// in user configuration. The value is an absolute path. Digits, letters,
	// periods (.), hyphens (-), and underscores (_) are allowed.
	Path string `json:"path,omitempty"`
}

// RuntimeConfig represents the runtime configuration.
type RuntimeConfig struct {
	// LVM write mode. "linear" indicates the linear mode.
	// "striped" indicates the striped mode, in which multiple disks are used
	// to form a strip to improve disk performance.
	LvType string `json:"lvType" required:"true"`
}

// ReinstallRuntimeConfig represents the runtime configuration for node reinstallation.
type ReinstallRuntimeConfig struct {
	// Available disk space of a single container on a node, in GB.
	// If this parameter is left blank or is set to 0, the default value is used.
	// In Device Mapper mode, the default value is 10.
	// In OverlayFS mode, the available space of a single container is not limited by default.
	// The dockerBaseSize setting takes effect only on EulerOS nodes in the cluster of the new version.
	// When Device Mapper is used, it is recommended to set dockerBaseSize to a value
	// less than or equal to 80 GB to avoid long initialization.
	DockerBaseSize int `json:"dockerBaseSize,omitempty"`

	// Container runtime. Defaults to "docker".
	Runtime Runtime `json:"runtime,omitempty"`
}

// Runtime represents the container runtime.
type Runtime struct {
	// 	Container runtime. Defaults to docker.
	// Enumeration values: docker, containerd
	Name string `json:"name,omitempty"`
}

// ReinstallK8sOptionsConfig represents the Kubernetes options configuration for node reinstallation.
type ReinstallK8sOptionsConfig struct {
	// Defined in key-value pairs. A maximum of 20 key-value pairs are allowed.
	// Key: Enter 1 to 63 characters, starting with a letter or digit.
	// Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed.
	// A DNS subdomain can be prefixed to a key and contain a maximum of 253 characters.
	// Example DNS subdomain: example.com/my-key
	// Value: The value can be left blank or contain 1 to 63 characters that start with a letter or digit.
	// Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed.
	Labels map[string]string `json:"labels,omitempty"`

	// Taints can be added for anti-affinity when creating nodes. A maximum of 20 taints can be added.
	Taints []Taint `json:"taints,omitempty"`

	// Maximum number of pods that can be created on a node, including the default system pods.
	// Value range: 16 to 256. This limit prevents the node from being overloaded with pods.
	MaxPods int `json:"maxPods,omitempty"`
}

// Taint represents a taint that can be applied to nodes for anti-affinity purposes.
type Taint struct {
	// Key: A key must contain 1 to 63 characters starting with a letter or digit.
	// Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed.
	// A DNS subdomain name can be used as the prefix of a key.
	Key string `json:"key" required:"true"`

	// Value: A value must start with a letter or digit and can contain a maximum of 63 characters.
	// It can include letters, digits, hyphens (-), underscores (_), and periods (.).
	Value string `json:"value,omitempty"`

	// Effect: Available options are NoSchedule, PreferNoSchedule, and NoExecute.
	Effect string `json:"effect" required:"true"`
}

// NodeLifecycleConfig represents the lifecycle configuration for a node.
type NodeLifecycleConfig struct {
	// Pre-installation script. Must be base-64 encoded.
	PreInstall string `json:"preInstall,omitempty"`
	// Post-installation script. Must be base-64 encoded.
	PostInstall string `json:"postInstall,omitempty"`
}

// ReinstallExtendParam represents the extended parameters for reinstalling.
type ReinstallExtendParam struct {
	// (Discarded) ID of the user image to run the target OS.
	// Specifying this parameter is equivalent to specifying imageID in ReinstallVolumeSpec.
	// The original value will be overwritten.
	AlphaCCE_NodeImageID string `json:"alpha.cce/NodeImageID,omitempty"`
}

// AddExistingNode function  is used to accept a node into a specified cluster. It returns job ID on successful execution.
func AddExistingNode(client *golangsdk.ServiceClient, clusterId string, opts AddExistingNodeOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /api/v3/projects/{project_id}/clusters/{cluster_id}/nodes/add
	raw, err := client.Post(client.ServiceURL("clusters", clusterId, "nodes", "add"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res AddExistingNodeResp
	return &res.JobId, extract.Into(raw.Body, &res)
}

type AddExistingNodeResp struct {
	// Job ID returned after the job is delivered. The job ID can be used to query the job execution status.
	JobId string `json:"jobid"`
}
