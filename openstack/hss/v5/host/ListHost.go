package hss

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListHostOpts struct {
	// Offset from which the query starts. If the value is less than 0, it is automatically converted to 0.
	Offset *int `q:"offset"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 10,
	// and a value greater than 200 will be automatically converted to 200.
	Limit int `q:"limit"`
	// HSS edition. Its values and their meaning are as follows:
	// hss.version.null: none
	// hss.version.enterprise: enterprise edition
	// hss.version.premium: premium edition
	// hss.version.container.enterprise: container edition
	Version string `q:"version"`
	// Agent status. Its value can be:
	// not_ installed
	// online
	// offline
	// install_failed
	// installing
	// not_online: All status except online, which is used only as a query condition.
	AgentStatus string `q:"agent_status"`
	// Detection result. Its value can be:
	// undetected
	// clean
	// risk
	// scanning
	DetectResult string `q:"detect_result"`
	// Server name
	HostName string `q:"host_name"`
	// Server ID
	HostID string `q:"host_id"`
	// Host status. Its value can be:
	// ACTIVE
	// SHUTOFF
	// BUILDING
	// ERROR
	HostStatus string `q:"host_status"`
	// OS type. Its value can be:
	// Linux
	// Windows
	OsType string `q:"os_type"`
	// Server private IP address
	PrivateIp string `q:"private_ip"`
	// Server public IP address
	PublicIp string `q:"public_ip"`
	// Public or private IP address
	IpAddr string `q:"ip_addr"`
	// Protection status. Its value can be:
	// closed
	// opened
	ProtectStatus string `q:"protect_status"`
	// Server group ID
	GroupId string `q:"group_id"`
	// Server group name
	GroupName string `q:"group_name"`
	// Policy group ID
	PolicyGroupId string `q:"policy_group_id"`
	// Policy group name
	PolicyGroupName string `q:"policy_group_name"`
	// on_demand: pay-per-use
	ChargingMode string `q:"charging_mode"`
	// Whether to forcibly synchronize servers from ECSs
	Refresh *bool `q:"refresh"`
	// Whether to return all the versions later than the current version
	AboveVersion *bool `q:"above_version"`
	// Whether a server is a non-cloud server
	OutsideHost *bool `q:"outside_host"`
	// Asset importance. Its value can be:
	// important
	// common
	// test
	AssetValue string `q:"asset_value"`
	// Asset tag
	Label string `q:"label"`
	// Asset server group
	ServerGroup string `q:"server_group"`
}

func ListHost(client *golangsdk.ServiceClient, opts ListHostOpts) ([]HostResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("host-management", "hosts").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v5/{project_id}/host-management/hosts
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return HostGroupPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractHosts(pages)
}

type HostPage struct {
	pagination.NewSinglePageBase
}

func ExtractHosts(r pagination.NewPage) ([]HostResp, error) {
	var s struct {
		Hosts []HostResp `json:"data_list"`
	}
	err := extract.Into(bytes.NewReader((r.(HostPage)).Body), &s)
	return s.Hosts, err
}

type HostResp struct {
	// Server ID
	ID string `json:"host_id"`
	// Server name
	Name string `json:"host_name"`
	// Agent ID
	AgentId string `json:"agent_id"`
	// Private IP address
	PrivateIp string `json:"private_ip"`
	// Elastic IP address
	PublicIp string `json:"public_ip"`
	// Server status. Its value can be:
	// ACTIVE
	// SHUTOFF
	// BUILDING
	// ERROR
	HostStatus string `json:"host_status"`
	// Agent status. Its value can be:
	// not_ installed
	// online
	// offline
	// install_failed
	// installing
	AgentStatus string `json:"agent_status"`
	// Installation result. This API is available only for agents that are installed in batches. The options are as follows:
	// install_succeed
	// network_access_timeout: Connection timed out. Network error.
	// invalid_port
	// auth_failed: The authentication failed due to incorrect password.
	// permission_denied: Insufficient permissions.
	// no_available_vpc: There is no server with an online agent in the current VPC.
	// install_exception
	// invalid_param: Incorrect parameter.
	// install_failed
	// package_unavailable
	// os_type_not_support: Incorrect OS type
	// os_arch_not_support: Incorrect OS architecture
	InstallResultCode string `json:"install_result_code"`
	// HSS edition. Its values and their meaning are as follows:
	// hss.version.null: none
	// hss.version.enterprise: enterprise edition
	// hss.version.premium: premium edition
	// hss.version.container.enterprise: container edition
	Version string `json:"version"`
	// Protection status. Its value can be:
	// closed
	// opened
	ProtectStatus string `json:"protect_status"`
	// System disk image
	OsImage string `json:"os_image"`
	// OS type. Its value can be:
	// Linux
	// Windows
	OsType string `json:"os_type"`
	// OS bit version
	OsBit string `json:"os_bit"`
	// Server scan result. Its value can be:
	// undetected
	// clean
	// risk
	// scanning
	DetectResult string `json:"detect_result"`
	// on_demand: pay-per-use
	ChargingMode string `json:"charging_mode"`
	// Cloud service resource instance ID (UUID)
	ResourceId string `json:"resource_id"`
	// Whether a server is a non-cloud server
	OutsideHost string `json:"outside_host"`
	// Server group ID
	GroupId string `json:"group_id"`
	// Server group name
	GroupName string `json:"group_name"`
	// Policy group ID
	PolicyGroupId string `json:"policy_group_id"`
	// Policy group name
	PolicyGroupName string `json:"policy_group_name"`
	// Asset risk
	Asset int `json:"asset"`
	// Total number of vulnerabilities, including Linux, Windows, Web-CMS, and application vulnerabilities.
	Vulnerability int `json:"vulnerability"`
	// Total number of baseline risks, including configuration risks and weak passwords.
	Baseline int `json:"baseline"`
	// Total intrusion risks
	Intrusion int `json:"intrusion"`
	// Asset importance. Its value can be:
	// important
	// common
	// test
	AssetValue string `json:"asset_value"`
	// Tag list
	Labels []string `json:"labels"`
	// Agent installation time, which is a timestamp.
	// The default unit is milliseconds.
	AgentCreateTime int64 `json:"agent_create_time"`
	// Time when the agent status is changed. This is a timestamp.
	// The default unit is milliseconds.
	AgentUpdateTime int64 `json:"agent_update_time"`
	// Agent version
	AgentVersion string `json:"agent_version"`
	// Upgrade status. Its value can be:
	// not_upgrade: Not upgraded. This is the default status. The customer has not delivered any upgrade command to the server.
	// upgrading: The upgrade is in progress.
	// upgrade_failed: The upgrade failed.
	// upgrade_succeed
	UpgradeStatus string `json:"upgrade_status"`
	// Upgrade failure cause. This parameter is displayed only if upgrade_status is upgrade_failed. Its value can be:
	// package_unavailable: The upgrade package fails to be parsed because the upgrade file is incorrect.
	// network_access_timeout: Failed to download the upgrade package because the network is abnormal.
	// agent_offline: The agent is offline.
	// hostguard_abnormal: The agent process is abnormal.
	// insufficient_disk_space
	// failed_to_replace_file: Failed to replace the file.
	UpgradeResultCode string `json:"upgrade_result_code"`
	// Whether the agent of the server can be upgraded
	Upgradable bool `json:"upgradable"`
}
