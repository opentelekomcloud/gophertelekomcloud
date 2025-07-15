package access_config

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOpts is a struct that contains all the parameters.
type CreateOpts struct {
	// Ingestion configuration name.
	Name string `json:"access_config_name" required:"true"`
	// Ingestion configuration type. AGENT: ECS access; K8S_CCE: CCE access
	Type string `json:"access_config_type" required:"true"`
	// Access configuration details.
	Details *AccessConfigDetails `json:"access_config_detail" required:"true"`
	// Log information
	LogInfo *LogInfo `json:"log_info" required:"true"`
	// Host group information
	HostGroupInfo *HostGroupInfo `json:"host_group_info,omitempty"`
	// Tag information. A tag key must be unique. Up to 20 tags are allowed.
	Tags []tags.ResourceTag `json:"access_config_tag,omitempty"`
	// Binary collection.
	BinaryCollect *bool `json:"binary_collect,omitempty"`
	// Log splitting.
	LogSplit *bool `json:"log_split,omitempty"`
	// Cluster ID
	ClusterId string `json:"cluster_id,omitempty"`
}

type AccessConfigDetails struct {
	// Collection paths.
	Paths []string `json:"paths,omitempty"`
	// Collection path blacklist.
	BlackPaths []string `json:"black_paths,omitempty"`
	// Log format.
	Format *AccessConfigFormat `json:"format" required:"true"`
	// Windows event logs.
	WindowsLogInfo *AccessConfigWindowsLogInfo `json:"windows_log_info,omitempty"`
	// Standard output switch. This parameter is used only for CCE log ingestion.
	Stdout *bool `json:"stdout,omitempty"`
	// Standard error switch. This parameter is used only for CCE log ingestion.
	Stderr *bool `json:"stderr,omitempty"`
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

type AccessConfigFormat struct {
	// Single-line logs.
	Single *AccessConfigFormatBody `json:"single,omitempty"`
	// Multi-line logs.
	Multi *AccessConfigFormatBody `json:"multi,omitempty"`
}

type AccessConfigFormatBody struct {
	// Mode
	Mode string `json:"mode,omitempty"`
	// Log time.If mode is system, the value is the current timestamp.
	// If mode is wildcard, the value is a time wildcard, which is used by ICAgent
	// to look for the log printing time as the beginning of a log event.
	// If the time format in a log event is 2019-01-01 23:59:59, the time wildcard is YYYY-MM-DD hh:mm:ss.
	// If the time format in a log event is 19-1-1 23:59:59, the time wildcard is YY-M-D hh:mm:ss.
	Value string `json:"value,omitempty"`
}

type AccessConfigWindowsLogInfo struct {
	// Type of Windows event logs to be collected.
	// Application: application event logs.
	// System: system event logs.
	// Security: security event logs.
	// Setup: startup event logs.
	Categories []string `json:"categorys" required:"true"`
	// Offset from first collection time.
	TimeOffset *AccessConfigTimeOffset `json:"time_offset" required:"true"`
	// Event level.
	// information: common information events, which do not affect system running.
	// warning: warning events, which may affect system running but do not cause system breakdown.
	// error: error events, which cause system breakdown or prevent the service from running properly.
	// critical: critical events, which may cause system or application failures.
	// verbose: detailed event information, which does not affect the system running.
	EventLevel []string `json:"event_level" required:"true"`
}

type AccessConfigTimeOffset struct {
	// Time offset.
	// When unit is day, the value ranges from 1 to 7.
	// When unit is hour, the value ranges from 1 to 168.
	// When unit is sec, the value ranges from 1 to 604800.
	Offset int64 `json:"offset" required:"true"`
	// Unit of the time offset.
	// day
	// hour
	// sec
	Unit string `json:"unit" required:"true"`
}

type LogInfo struct {
	// Log group ID.
	LogGroupId string `json:"log_group_id" required:"true"`
	// Log stream ID.
	LogStreamId string `json:"log_stream_id" required:"true"`
}

type HostGroupInfo struct {
	// List of host group IDs.
	HostGroupIds *[]string `json:"host_group_id_list" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*AccessConfigInfo, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/lts/access-config
	raw, err := client.Post(client.ServiceURL("lts", "access-config"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}

	var res AccessConfigInfo
	err = extract.Into(raw.Body, &res)
	return &res, err
}
