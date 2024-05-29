package channel

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID string `json:"-"`
	// VPC channel name.
	// It can contain 3 to 64 characters, starting with a letter.
	// Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed.
	Name string `json:"name" required:"true"`
	// Host port of the VPC channel.
	// Range: 1-65535.
	Port int `json:"port" required:"true"`
	// Distribution algorithm.
	// 1: Weighted round robin (WRR).
	// 2: Weighted least connections (WLC).
	// 3: Source hashing.
	// 4: URI hashing.
	LbAlgorithm int `json:"balance_strategy" required:"true"`
	// Member type of the VPC channel.
	// ip
	// ecs
	MemberType string `json:"member_type" required:"true"`
	// VPC channel type. The default type is server.
	// 2: Server type.
	// 3: Microservice type.
	Type int `json:"type,omitempty"`
	// Dictionary code of the VPC channel.
	// The value can contain letters, digits, hyphens (-), underscores (_), and periods (.).
	// This parameter is currently not supported.
	Code string `json:"dict_code,omitempty"`
	// Backend server groups of the VPC channel.
	MemberGroups []MemberGroups `json:"member_groups,omitempty"`
	// Backend instances of the VPC channel.
	Members []Members `json:"members,omitempty"`
	// Health check details.
	VpcHealthConfig *VpcHealthConfig `json:"vpc_health_config,omitempty"`
	// Microservice details.
	MicroserviceConfig *MicroserviceConfig `json:"microservice_info,omitempty"`
}

type MemberGroups struct {
	// Name of the VPC channel's backend server group.
	// It can contain 3 to 64 characters, starting with a letter.
	// Only letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.
	Name string `json:"member_group_name" required:"true"`
	// Description of the backend server group.
	Description string `json:"member_group_remark,omitempty"`
	// Weight of the backend server group.
	// If the server group contains servers and a weight has been set for it, the weight is automatically used to assign weights to servers in this group.
	Weight *int `json:"member_group_weight,omitempty"`
	// Dictionary code of the backend server group.
	// The value can contain letters, digits, hyphens (-), underscores (_), and periods (.).
	// Currently, this parameter is not supported.
	Code string `json:"dict_code,omitempty"`
	// Version of the backend server group.
	// This parameter is supported only when the VPC channel type is microservice.
	MicroserviceVersion string `json:"microservice_version,omitempty"`
	// Port of the backend server group.
	// This parameter is supported only when the VPC channel type is microservice.
	// If the port number is 0, all addresses in the backend server group use the original load balancing port to inherit logic.
	MicroservicePort int `json:"microservice_port,omitempty"`
	// Tags of the backend server group.
	// This parameter is supported only when the VPC channel type is microservice.
	MicroserviceTags []MicroserviceTags `json:"microservice_labels,omitempty"`
}

type MicroserviceTags struct {
	// Tag name.
	// Start and end with a letter or digit.
	// Use only letters, digits, hyphens (-), underscores (_), and periods (.). (Max. 63 characters.)
	Key string `json:"label_name" required:"true"`
	// Tag value.
	// Start and end with a letter or digit.
	// Use only letters, digits, hyphens (-), underscores (_), and periods (.). (Max. 63 characters.)
	Value string `json:"label_value" required:"true"`
}

type Members struct {
	// Backend server address.
	// This parameter is required when the member type is IP address.
	Host string `json:"host,omitempty"`
	// Weight.
	// The higher the weight is, the more requests a backend service will receive.
	Weight *int `json:"weight,omitempty"`
	// Indicates whether the backend service is a standby node.
	// After you enable this function, the backend service serves as a standby node.
	// It works only when all non-standby nodes are faulty.
	// This function is supported only when your gateway has been upgraded to the corresponding version.
	// If your gateway does not support this function, contact technical support.
	IsBackup *bool `json:"is_backup,omitempty"`
	// Backend server group name. The server group facilitates backend service address modification.
	MemberGroupName string `json:"member_group_name,omitempty"`
	// Backend server status.
	// 1: available
	// 2: unavailable
	Status int `json:"status" required:"true"`
	// Backend server port.
	Port *int `json:"port,omitempty"`
	// Backend server ID.
	// This parameter is required if the backend instance type is ecs.
	// The value can contain 1 to 64 characters, including letters, digits, hyphens (-), and underscores (_).
	EcsId string `json:"ecs_id,omitempty"`
	// Backend server name.
	// This parameter is required if the backend instance type is ecs.
	// The value can contain 1 to 64 characters, including letters, digits, hyphens (-), underscores (_), and periods (.).
	EcsName string `json:"ecs_name,omitempty"`
}

type VpcHealthConfig struct {
	// Protocol for performing health checks on backend servers in the VPC channel.
	// TCP
	// HTTP
	// HTTPS
	Protocol string `json:"protocol" required:"true"`
	// Destination path for health checks. This parameter is required if protocol is set to http or https.
	Path string `json:"path,omitempty"`
	// Request method for health checks.
	// GET
	// HEAD
	Method string `json:"method,omitempty"`
	// Destination port for health checks.
	// If this parameter is not specified or set to 0, the host port of the VPC channel is used.
	// If this parameter is set to a non-zero value, the corresponding port is used for health checks.
	Port *int `json:"port,omitempty"`
	// Healthy threshold. It refers to the number of consecutive successful checks required for a backend server to be considered healthy.
	// Min: 1, Max: 10
	HealthyThreshold int `json:"threshold_normal" required:"true"`
	// Unhealthy threshold, which refers to the number of consecutive failed checks required for a backend server to be considered unhealthy.
	// Min: 1, Max: 10
	UnhealthyThreshold int `json:"threshold_abnormal" required:"true"`
	// Interval between consecutive checks. Unit: s. The value must be greater than the value of timeout.
	// Min: 1, Max: 300
	Interval int `json:"time_interval" required:"true"`
	// Response codes for determining a successful HTTP response.
	// The value can be any integer within 100-599 in one of the following formats:
	// Multiple values, for example, 200,201,202
	// Range, for example, 200-299
	// Multiple values and ranges, for example, 201,202,210-299.
	// This parameter is required if protocol is set to http.
	HttpCode string `json:"http_code,omitempty"`
	// Indicates whether to enable two-way authentication.
	// If this function is enabled, the certificate specified in the backend_client_certificate configuration item of the gateway is used.
	EnableClientSsl *bool `json:"enable_client_ssl,omitempty"`
	// Health check result.
	// 1: available
	// 2: unavailable
	Status *int `json:"status,omitempty"`
	// Timeout for determining whether a health check fails. Unit: s. The value must be less than the value of time_interval.
	// Min: 1, Max: 30
	Timeout int `json:"timeout" required:"true"`
}

type MicroserviceConfig struct {
	// Microservice type. Options:
	// CSE: CSE microservice registration center
	// CCE: CCE workload
	ServiceType string `json:"service_type,omitempty"`
	// CSE microservice details. This parameter is required if service_type is set to CSE.
	CseInfo *CseInfo `json:"cse_info,omitempty"`
	// CCE workload details.
	// This parameter is required if service_type is set to CCE.
	// Either app_name or any of label_key and label_value must be set.
	// If only app_name is set, label_key='app' and label_value=app_name.
	CceInfo *CceInfo `json:"cce_info,omitempty"`
	// CCE Service details.
	CceServiceInfo *CceServiceInfo `json:"cce_service_info,omitempty"`
}

type CseInfo struct {
	// Microservice engine ID.
	EngineID string `json:"engine_id" required:"true"`
	// Microservice ID.
	ServiceID string `json:"service_id" required:"true"`
}

type CceInfo struct {
	// CCE cluster ID.
	ClusterId string `json:"cluster_id" required:"true"`
	// Namespace.
	Namespace string `json:"namespace" required:"true"`
	// Workload type.
	// deployment
	// statefulset
	// daemonset
	WorkloadType string `json:"workload_type" required:"true"`
	// App name.
	// Start with a letter, and include only letters, digits, periods (.), hyphens (-), and underscores (_). (1 to 64 characters)
	AppName string `json:"app_name,omitempty"`
	// Service label key.
	// Start with a letter or digit, and use only letters, digits, and these special characters: -_./:(). (1 to 64 characters)
	LabelKey string `json:"label_key,omitempty"`
	// Service label value.
	// Start with a letter, and include only letters, digits, periods (.), hyphens (-), and underscores (_). (1 to 64 characters)
	LabelValue string `json:"label_value,omitempty"`
}

type CceServiceInfo struct {
	// CCE cluster ID.
	ClusterId string `json:"cluster_id" required:"true"`
	// Namespace.
	Namespace string `json:"namespace" required:"true"`
	// Service name.
	// Start with a letter, and use only letters, digits, periods (.), hyphens (-), and underscores (_). (1 to 64 characters)
	ServiceName string `json:"service_name" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*ChannelResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/apigw/instances/{instance_id}/vpc-channels
	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "vpc-channels"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res ChannelResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ChannelResp struct {
	// VPC channel ID.
	ID string `json:"id"`
	// VPC channel name.
	// It can contain 3 to 64 characters, starting with a letter. Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed.
	Name string `json:"name"`
	// Host port of the VPC channel.
	Port int `json:"port"`
	// Distribution algorithm.
	// 1: Weighted round robin (WRR).
	// 2: Weighted least connections (WLC).
	// 3: Source hashing.
	// 4: URI hashing.
	LbAlgorithm int `json:"balance_strategy"`
	// Member type of the VPC channel.
	// ip
	// ecs
	MemberType string `json:"member_type"`
	// VPC channel type. The default type is server.
	// 2: Server type.
	// 3: Microservice type.
	Type int `json:"type"`
	// Dictionary code of the VPC channel.
	// The value can contain letters, digits, hyphens (-), underscores (_), and periods (.).
	// This parameter is currently not supported.
	Code string `json:"dict_code"`
	// Time when the VPC channel is created.
	CreatedAt string `json:"create_time"`
	// VPC channel status.
	// 1: normal
	// 2: abnormal
	Status int `json:"status"`
	// Backend server groups.
	MemberGroups []MemberGroupsResp `json:"member_groups"`
	// Backend servers.
	Members []MembersResp `json:"members"`
	// Microservice information.
	MicroserviceConfig *MicroserviceConfigResp `json:"microservice_info"`
	// Health check details.
	VpcHealthConfig *VpcHealthConfigResp `json:"vpc_health_config"`
}

type MemberGroupsResp struct {
	// ID of the backend server group of the VPC channel.
	ID string `json:"id"`
	// Name of the VPC channel's backend server group.
	// It can contain 3 to 64 characters, starting with a letter.
	// Only letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.
	Name string `json:"member_group_name"`
	// Description of the backend server group.
	Description string `json:"member_group_remark"`
	// Weight of the backend server group.
	// If the server group contains servers and a weight has been set for it, the weight is automatically used to assign weights to servers in this group.
	Weight int `json:"member_group_weight"`
	// Dictionary code of the backend server group.
	// The value can contain letters, digits, hyphens (-), underscores (_), and periods (.).
	// Currently, this parameter is not supported.
	Code string `json:"dict_code"`
	// Version of the backend server group.
	// This parameter is supported only when the VPC channel type is microservice.
	MicroserviceVersion string `json:"microservice_version"`
	// Port of the backend server group.
	// This parameter is supported only when the VPC channel type is microservice.
	// If the port number is 0, all addresses in the backend server group use the original load balancing port to inherit logic.
	MicroservicePort int `json:"microservice_port"`
	// Tags of the backend server group.
	// This parameter is supported only when the VPC channel type is microservice.
	MicroserviceTags []MicroserviceTags `json:"microservice_labels"`
	// Time when the backend server group is created.
	CreatedAt string `json:"create_time"`
	// Time when the backend server group is updated.
	UpdatedAt string `json:"update_time"`
}

type MicroserviceConfigResp struct {
	// Microservice ID.
	ID string `json:"id"`
	// Gateway ID.
	GatewayID string `json:"instance_id"`
	// Microservice type. Options:
	// CSE: CSE microservice registration center
	// CCE: CCE workload
	ServiceType string `json:"service_type"`
	// CSE microservice details. This parameter is required if service_type is set to CSE.
	CseInfo *CseInfoResp `json:"cse_info"`
	// CCE workload details.
	// This parameter is required if service_type is set to CCE.
	// Either app_name or any of label_key and label_value must be set.
	// If only app_name is set, label_key='app' and label_value=app_name.
	CceInfo *CceInfoResp `json:"cce_info"`
	// CCE Service details.
	CceServiceInfo *CceServiceInfoResp `json:"cce_service_info"`
	// Time when the backend server group is created.
	CreatedAt string `json:"create_time"`
	// Time when the backend server group is updated.
	UpdatedAt string `json:"update_time"`
}

type CseInfoResp struct {
	// Microservice engine ID.
	EngineID string `json:"engine_id"`
	// Microservice ID.
	ServiceID string `json:"service_id"`
	// Microservice engine name.
	EngineName string `json:"engine_name"`
	// Microservice name.
	ServiceName string `json:"service_name"`
	// Registration center address.
	RegisterAddress string `json:"register_address"`
	// App to which the microservice belongs.
	CseAppId string `json:"cse_app_id"`
	// Microservice version, which has been discarded and is reflected in the version of the backend server group.
	Version string `json:"version"`
}

type CceInfoResp struct {
	// CCE cluster ID.
	ClusterId string `json:"cluster_id"`
	// Namespace.
	Namespace string `json:"namespace"`
	// Workload type.
	// deployment
	// statefulset
	// daemonset
	WorkloadType string `json:"workload_type"`
	// App name.
	// Start with a letter, and include only letters, digits, periods (.), hyphens (-), and underscores (_). (1 to 64 characters)
	AppName string `json:"app_name"`
	// Service label key.
	// Start with a letter or digit, and use only letters, digits, and these special characters: -_./:(). (1 to 64 characters)
	LabelKey string `json:"label_key"`
	// Service label value.
	// Start with a letter, and include only letters, digits, periods (.), hyphens (-), and underscores (_). (1 to 64 characters)
	LabelValue string `json:"label_value"`
	// CCE cluster name.
	ClusterName string `json:"cluster_name"`
}

type CceServiceInfoResp struct {
	// CCE cluster ID.
	ClusterId string `json:"cluster_id" required:"true"`
	// Namespace.
	Namespace string `json:"namespace" required:"true"`
	// Service name.
	// Start with a letter, and use only letters, digits, periods (.), hyphens (-), and underscores (_). (1 to 64 characters)
	ServiceName string `json:"service_name" required:"true"`
	// CCE cluster name.
	ClusterName string `json:"cluster_name"`
}

type MembersResp struct {
	// Backend server address.
	// This parameter is required when the member type is IP address.
	Host string `json:"host"`
	// Weight.
	// The higher the weight is, the more requests a backend service will receive.
	Weight *int `json:"weight"`
	// Indicates whether the backend service is a standby node.
	// After you enable this function, the backend service serves as a standby node.
	// It works only when all non-standby nodes are faulty.
	// This function is supported only when your gateway has been upgraded to the corresponding version.
	// If your gateway does not support this function, contact technical support.
	IsBackup *bool `json:"is_backup"`
	// Backend server group name. The server group facilitates backend service address modification.
	MemberGroupName string `json:"member_group_name"`
	// Backend server status.
	// 1: available
	// 2: unavailable
	Status int `json:"status"`
	// Backend server port.
	Port *int `json:"port"`
	// Backend server ID.
	// This parameter is required if the backend instance type is ecs.
	// The value can contain 1 to 64 characters, including letters, digits, hyphens (-), and underscores (_).
	EcsId string `json:"ecs_id"`
	// Backend server name.
	// This parameter is required if the backend instance type is ecs.
	// The value can contain 1 to 64 characters, including letters, digits, hyphens (-), underscores (_), and periods (.).
	EcsName string `json:"ecs_name"`
	// Backend instance ID.
	ID string `json:"id"`
	// VPC channel ID.
	ChannelID string `json:"vpc_channel_id"`
	// Time when the backend server is added to the VPC channel.
	CreatedAt string `json:"create_time"`
	// Backend server group ID.
	MemberGroupId string `json:"member_group_id"`
}

type VpcHealthConfigResp struct {
	// Protocol for performing health checks on backend servers in the VPC channel.
	// TCP
	// HTTP
	// HTTPS
	Protocol string `json:"protocol"`
	// Destination path for health checks. This parameter is required if protocol is set to http or https.
	Path string `json:"path"`
	// Request method for health checks.
	// GET
	// HEAD
	Method string `json:"method"`
	// Destination port for health checks.
	// If this parameter is not specified or set to 0, the host port of the VPC channel is used.
	// If this parameter is set to a non-zero value, the corresponding port is used for health checks.
	Port *int `json:"port"`
	// Healthy threshold. It refers to the number of consecutive successful checks required for a backend server to be considered healthy.
	// Min: 1, Max: 10
	HealthyThreshold int `json:"threshold_normal"`
	// Unhealthy threshold, which refers to the number of consecutive failed checks required for a backend server to be considered unhealthy.
	// Min: 1, Max: 10
	UnhealthyThreshold int `json:"threshold_abnormal"`
	// Interval between consecutive checks. Unit: s. The value must be greater than the value of timeout.
	// Min: 1, Max: 300
	Interval int `json:"time_interval"`
	// Response codes for determining a successful HTTP response.
	// The value can be any integer within 100-599 in one of the following formats:
	// Multiple values, for example, 200,201,202
	// Range, for example, 200-299
	// Multiple values and ranges, for example, 201,202,210-299.
	// This parameter is required if protocol is set to http.
	HttpCode string `json:"http_code"`
	// Indicates whether to enable two-way authentication.
	// If this function is enabled, the certificate specified in the backend_client_certificate configuration item of the gateway is used.
	EnableClientSsl *bool `json:"enable_client_ssl"`
	// Health check result.
	// 1: available
	// 2: unavailable
	Status *int `json:"status"`
	// Timeout for determining whether a health check fails. Unit: s. The value must be less than the value of time_interval.
	// Min: 1, Max: 30
	Timeout int `json:"timeout" required:"true"`
	// Backend instance ID.
	ID string `json:"id"`
	// VPC channel ID.
	ChannelID string `json:"vpc_channel_id"`
	// Time when the backend server is added to the VPC channel.
	CreatedAt string `json:"create_time"`
}
