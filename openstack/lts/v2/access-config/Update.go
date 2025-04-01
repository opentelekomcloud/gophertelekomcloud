package access_config

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type UpdateOpts struct {
	// Ingestion configuration ID.
	ID string `json:"access_config_id" required:"true"`
	// Access configuration details.
	Details *AccessConfigDetailsUpdate `json:"access_config_detail,omitempty"`
	// Host group information
	HostGroupInfo *HostGroupInfo `json:"host_group_info,omitempty"`
	// Tag information. A tag key must be unique. Up to 20 tags are allowed.
	Tags *[]tags.ResourceTag `json:"access_config_tag,omitempty"`
	// Binary collection.
	BinaryCollect *bool `json:"binary_collect,omitempty"`
	// Log splitting.
	LogSplit *bool `json:"log_split,omitempty"`
	// Cluster ID
	ClusterId string `json:"cluster_id,omitempty"`
}

type AccessConfigDetailsUpdate struct {
	// Collection paths.
	Paths []string `json:"paths,omitempty"`
	// Collection path blacklist.
	BlackPaths []string `json:"black_paths,omitempty"`
	// Log format.
	Format *AccessConfigFormat `json:"format,omitempty"`
	// Windows event logs.
	WindowsLogInfo *AccessConfigWindowsLogInfoUpdate `json:"windows_log_info,omitempty"`
	// Standard output switch. This parameter is used only for CCE log ingestion.
	Stdout bool `json:"stdout,omitempty"`
	// Standard error switch. This parameter is used only for CCE log ingestion.
	Stderr bool `json:"stderr,omitempty"`
	// CCE log ingestion type. This parameter is used only for CCE log ingestion.
	PathType string `json:"pathType,omitempty"`
	// Regular expression matching of Kubernetes namespaces. This parameter is used only for CCE log ingestion.
	NamespaceRegex string `json:"namespaceRegex,omitempty"`
	// Regular expression matching of Kubernetes pods. This parameter is used only for CCE log ingestion.
	PodNameRegex string `json:"podNameRegex,omitempty"`
	// Regular expression matching of Kubernetes container names. This parameter is used only for CCE log ingestion.
	ContainerNameRegex string `json:"containerNameRegex,omitempty"`
	// Container label whitelist. You can create up to 30 whitelists. The key names must be unique. This parameter is used only for CCE log ingestion.
	IncludeLabels map[string]string `json:"includeLabels,omitempty"`
	// Container label blacklist. You can create up to 30 blacklists. The key names must be unique. This parameter is used only for CCE log ingestion.
	ExcludeLabels map[string]string `json:"excludeLabels,omitempty"`
	// Environment variable whitelist. You can create up to 30 whitelists. The key names must be unique. This parameter is used only for CCE log ingestion.
	IncludeEnvs map[string]string `json:"includeEnvs,omitempty"`
	// Environment variable blacklist. You can create up to 30 blacklists. The key names must be unique. This parameter is used only for CCE log ingestion.
	ExcludeEnvs map[string]string `json:"excludeEnvs,omitempty"`
	// Container label. You can create up to 30 labels. The key names must be unique. This parameter is used only for CCE log ingestion.
	LogLabels map[string]string `json:"logLabels,omitempty"`
	// Environment variable label. You can create up to 30 labels. The key names must be unique. This parameter is used only for CCE log ingestion.
	LogEnvs map[string]string `json:"logEnvs,omitempty"`
	// Kubernetes label whitelist. You can create up to 30 whitelists. The key names must be unique. This parameter is used only for CCE log ingestion.
	IncludeK8sLabels map[string]string `json:"includeK8sLabels,omitempty"`
	// Kubernetes label blacklist. You can create up to 30 blacklists. The key names must be unique. This parameter is used only for CCE log ingestion.
	ExcludeK8sLabels map[string]string `json:"excludeK8sLabels,omitempty"`
	// Kubernetes label. You can create up to 30 labels. The key names must be unique. This parameter is used only for CCE log ingestion.
	LogK8s map[string]string `json:"logK8s,omitempty"`
}

type AccessConfigWindowsLogInfoUpdate struct {
	// Type of Windows event logs to be collected.
	// Application: application event logs.
	// System: system event logs.
	// Security: security event logs.
	// Setup: startup event logs.
	Categories []string `json:"categorys,omitempty"`
	// Offset from first collection time.
	TimeOffset *AccessConfigTimeOffset `json:"time_offset,omitempty"`
	// Event level.
	// information: common information events, which do not affect system running.
	// warning: warning events, which may affect system running but do not cause system breakdown.
	// error: error events, which cause system breakdown or prevent the service from running properly.
	// critical: critical events, which may cause system or application failures.
	// verbose: detailed event information, which does not affect the system running.
	EventLevel []string `json:"event_level,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*AccessConfigInfo, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/lts/access-config
	raw, err := client.Put(client.ServiceURL("lts", "access-config"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res AccessConfigInfo
	err = extract.Into(raw.Body, &res)
	return &res, err
}
