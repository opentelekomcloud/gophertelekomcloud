package access_config

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type DeleteOpts struct {
	AccessConfigIds []string `json:"access_config_id_list" required:"true"`
}

// Delete a host group by id
func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (*DeleteResult, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// DELETE /v3/{project_id}/lts/access-config
	raw, err := client.DeleteWithBody(client.ServiceURL("lts", "access-config"), b, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResult
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// DeleteResult represents the API response after deletion
type DeleteResult struct {
	// Host group details.
	Result []AccessConfigInfo `json:"result"`
	// Number of deleted host groups.
	Total int64 `json:"total"`
}

type AccessConfigInfo struct {
	// Cross-account log ingestion ID.
	ID string `json:"access_config_id"`
	// Cross-account log ingestion name.
	Name string `json:"access_config_name"`
	// Cross-account log ingestion type.
	Type string `json:"access_config_type"`
	// Creation time.
	CreatedAt int64 `json:"create_time"`
	// Log splitting.
	LogSplit bool `json:"log_split"`
	// Binary collection.
	BinaryCollect bool `json:"binary_collect"`
	// CCE cluster ID
	ClusterId string `json:"cluster_id"`
	// Ingestion configuration details.
	AccessConfigDetail *AccessConfigDetailResponse `json:"access_config_detail"`
	// Log details.
	LogInfo *AccessConfigQueryLogResponse `json:"log_info"`
	// Host group ID list.
	HostGroupInfo *AccessConfigHostGroupIdsResponse `json:"host_group_info"`
	// Tag information.
	Tags []tags.ResourceTag `json:"access_config_tag"`
}

type AccessConfigDetailResponse struct {
	// Collection paths.
	Paths []string `json:"paths"`
	// Collection path blacklist.
	BlackPaths []string `json:"black_paths"`
	// Log format.
	Format *AccessConfigFormatResponse `json:"format"`
	// Windows event logs.
	WindowsLogInfo *AccessConfigWindowsLogInfoResponse `json:"windows_log_info"`
	// Standard output switch. This parameter is used only for CCE log ingestion.
	Stdout bool `json:"stdout"`
	// Standard error switch. This parameter is used only for CCE log ingestion.
	Stderr bool `json:"stderr"`
	// CCE log ingestion type. This parameter is used only for CCE log ingestion.
	PathType string `json:"pathType"`
	// Regular expression matching of Kubernetes namespaces. This parameter is used only for CCE log ingestion.
	NamespaceRegex string `json:"namespaceRegex"`
	// Regular expression matching of Kubernetes pods. This parameter is used only for CCE log ingestion.
	PodNameRegex string `json:"podNameRegex"`
	// Regular expression matching of Kubernetes container names. This parameter is used only for CCE log ingestion.
	ContainerNameRegex string `json:"containerNameRegex"`
	// Container label whitelist. You can create up to 30 whitelists. The key names must be unique. This parameter is used only for CCE log ingestion.
	IncludeLabels map[string]string `json:"includeLabels"`
	// Container label blacklist. You can create up to 30 blacklists. The key names must be unique. This parameter is used only for CCE log ingestion.
	ExcludeLabels map[string]string `json:"excludeLabels"`
	// Environment variable whitelist. You can create up to 30 whitelists. The key names must be unique. This parameter is used only for CCE log ingestion.
	IncludeEnvs map[string]string `json:"includeEnvs"`
	// Environment variable blacklist. You can create up to 30 blacklists. The key names must be unique. This parameter is used only for CCE log ingestion.
	ExcludeEnvs map[string]string `json:"excludeEnvs"`
	// Container label. You can create up to 30 labels. The key names must be unique. This parameter is used only for CCE log ingestion.
	LogLabels map[string]string `json:"logLabels"`
	// Environment variable label. You can create up to 30 labels. The key names must be unique. This parameter is used only for CCE log ingestion.
	LogEnvs map[string]string `json:"logEnvs"`
	// Kubernetes label whitelist. You can create up to 30 whitelists. The key names must be unique. This parameter is used only for CCE log ingestion.
	IncludeK8sLabels map[string]string `json:"includeK8sLabels"`
	// Kubernetes label blacklist. You can create up to 30 blacklists. The key names must be unique. This parameter is used only for CCE log ingestion.
	ExcludeK8sLabels map[string]string `json:"excludeK8sLabels"`
	// Kubernetes label. You can create up to 30 labels. The key names must be unique. This parameter is used only for CCE log ingestion.
	LogK8s map[string]string `json:"logK8s"`
}

type AccessConfigFormatResponse struct {
	// Single-line logs.
	Single *AccessConfigFormatBody `json:"single"`
	// Multi-line logs.
	Multi *AccessConfigFormatBody `json:"multi"`
}

type AccessConfigWindowsLogInfoResponse struct {
	// Type of Windows event logs to be collected.
	// Application: application event logs.
	// System: system event logs.
	// Security: security event logs.
	// Setup: startup event logs.
	Categories []string `json:"categorys"`
	// Offset from first collection time.
	TimeOffset *AccessConfigTimeOffsetResponse `json:"time_offset"`
	EventLevel []string                        `json:"event_level"`
}

type AccessConfigTimeOffsetResponse struct {
	// Time offset.
	// When unit is day, the value ranges from 1 to 7.
	// When unit is hour, the value ranges from 1 to 168.
	// When unit is sec, the value ranges from 1 to 604800.
	Offset int64 `json:"offset"`
	// Unit of the time offset.
	// day
	// hour
	// sec
	Unit string `json:"unit"`
}

type AccessConfigQueryLogResponse struct {
	// Log group ID.
	LogGroupId string `json:"log_group_id"`
	// Log stream ID.
	LogStreamId string `json:"log_stream_id"`
	// Log group name.
	LogGroupName string `json:"log_group_name"`
	// Log stream name.
	LogStreamName string `json:"log_stream_name"`
	// Log group alias.
	LogGroupNameAlias string `json:"log_group_name_alias"`
	// Log stream alias.
	LogStreamNameAlias string `json:"log_stream_name_alias"`
}

type AccessConfigHostGroupIdsResponse struct {
	HostGroupIds []string `json:"host_group_id_list"`
}
