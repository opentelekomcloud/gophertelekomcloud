package event

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Offset from which the query starts. If the value is less than 0, it is automatically converted to 0.
	Offset *int `q:"offset"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 10,
	// and a value greater than 200 will be automatically converted to 200.
	Limit int `q:"limit"`
	// Event category. Its value can be:
	// host: host security event
	// container: container security event
	Category string `q:"category" required:"true"`
	// Enterprise project ID.
	// The value 0 indicates the default enterprise project.
	// To query all enterprise projects, set this parameter to all_granted_eps.
	EnterpriseProjectId string `q:"enterprise_project_id"`
	// Number of days to be queried.
	// This parameter is mutually exclusive with begin_time and end_time.
	Days int `q:"last_days"`
	// Server name
	HostName string `q:"host_name"`
	// Host ID
	HostID string `q:"host_id"`
	// Server IP address
	PrivateIP string `q:"private_ip"`
	// Container instance name
	ContainerName string `q:"container_name"`
	// Intrusion type. Its value can be:
	// 1001: Malware
	// 1010: Rootkit
	// 1011: Ransomware
	// 1015: Web shell
	// 1017: Reverse shell
	// 2001: Common vulnerability exploit
	// 3002: File privilege escalation
	// 3003: Process privilege escalation
	// 3004: Important file change
	// 3005: File/Directory change
	// 3007: Abnormal process behavior
	// 3015: High-risk command execution
	// 3018: Abnormal shell
	// 3027: Suspicious crontab tasks
	// 4002: Brute-force attack
	// 4004: Abnormal login
	// 4006: Invalid system account
	EventTypes []string `q:"event_types"`
	// Status. Its value can be:
	// unhandled
	// handled
	HandleStatus string `q:"handle_status"`
	// Threat level. Its value can be:
	// Security
	// Low
	// Medium
	// High
	// Critical
	Severity string `q:"severity"`
	// Customized start time of a segment. The timestamp is accurate to seconds.
	// The begin_time should be no more than two days earlier than the end_time.
	// This parameter is mutually exclusive with the queried duration.
	BeginTime string `q:"begin_time"`
	// Customized end time of a segment. The timestamp is accurate to seconds.
	// The begin_time should be no more than two days earlier than the end_time.
	// This parameter is mutually exclusive with the queried duration.
	EndTime string `q:"end_time"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]EventResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("event", "events").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v5/{project_id}/event/events
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return EventPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractEvents(pages)
}

type EventPage struct {
	pagination.NewSinglePageBase
}

func ExtractEvents(r pagination.NewPage) ([]EventResp, error) {
	var s struct {
		Events []EventResp `json:"data_list"`
	}
	err := extract.Into(bytes.NewReader((r.(EventPage)).Body), &s)
	return s.Events, err
}

type EventResp struct {
	// Event ID
	ID string `json:"event_id"`
	// Event category. Its value can be:
	// container_1001: Container namespace
	// container_1002: Container open port
	// container_1003: Container security option
	// container_1004: Container mount directory
	// containerescape_0001: High-risk system call
	// containerescape_0002: Shocker attack
	// containerescape_0003: Dirty Cow attack
	// containerescape_0004: Container file escape
	// dockerfile_001: Modification of user-defined protected container file
	// dockerfile_002: Modification of executable files in the container file system
	// dockerproc_001: Abnormal container process
	// fileprotect_0001: File privilege escalation
	// fileprotect_0002: Key file change
	// fileprotect_0003: AuthorizedKeysFile path change
	// fileprotect_0004: File directory change
	// login_0001: Brute-force attack attempt
	// login_0002: Brute-force attack succeeded
	// login_1001: Succeeded login
	// login_1002: Remote login
	// login_1003: Weak password
	// malware_0001: Shell change
	// malware_0002: Reverse shell
	// malware_1001: Malicious program
	// procdet_0001: Abnormal process behavior
	// procdet_0002: Process privilege escalation
	// procreport_0001: High-risk command
	// user_1001: Account change
	// user_1002: Unsafe account
	// vmescape_0001: Sensitive command executed on VM
	// vmescape_0002: Sensitive file accessed by virtualization process
	// vmescape_0003: Abnormal VM port access
	// webshell_0001: Web shell
	// network_1001: Mining
	// network_1002: DDoS attacks
	// network_1003: Malicious scanning
	// network_1004: Attack in sensitive areas
	// crontab_1001: Suspicious crontab task
	EventClassId string `json:"event_class_id"`
	// Intrusion type. Its value can be:
	// 1001: Malware
	// 1010: Rootkit
	// 1011: Ransomware
	// 1015: Web shell
	// 1017: Reverse shell
	// 2001: Common vulnerability exploit
	// 3002: File privilege escalation
	// 3003: Process privilege escalation
	// 3004: Important file change
	// 3005: File/Directory change
	// 3007: Abnormal process behavior
	// 3015: High-risk command execution
	// 3018: Abnormal shell
	// 3027: Suspicious crontab tasks
	// 4002: Brute-force attack
	// 4004: Abnormal login
	// 4006: Invalid system account
	EventType int `json:"event_type"`
	// Event name
	EventName string `json:"event_name"`
	// Threat level. Its value can be:
	// Security
	// Low
	// Medium
	// High
	// Critical
	Severity string `json:"severity"`
	// Container instance name. This API is available only for container alarms.
	ContainerName string `json:"container_name"`
	// Image name. This API is available only for container alarms.
	ImageName string `json:"image_name"`
	// Server name
	HostName string `json:"host_name"`
	// Host ID
	HostId string `json:"host_id"`
	// Server private IP address
	PrivateIP string `json:"private_ip"`
	// Elastic IP address
	PublicIP string `json:"public_ip"`
	// OS type. Its value can be:
	// Linux
	// Windows
	OsType string `json:"os_type"`
	// Server status. The options are as follows:
	// ACTIVE
	// SHUTOFF
	// BUILDING
	// ERROR
	HostStatus string `json:"host_status"`
	// Agent status. Its value can be:
	// installed
	// not_installed
	// online
	// offline
	// install_failed
	// installing
	AgentStatus string `json:"agent_status"`
	// Protection status. Its value can be:
	// closed
	// opened
	ProtectStatus string `json:"protect_status"`
	// Asset importance. The options are as follows:
	// important
	// common
	// test
	AssetValue string `json:"asset_value"`
	// Attack phase. Its value can be:
	// reconnaissance
	// weaponization
	// delivery
	// exploit
	// installation
	// command_and_control
	// actions
	AttackPhase string `json:"attack_phase"`
	// Attack tag. Its value can be:
	// attack_success
	// attack_attempt
	// attack_blocked
	// abnormal_behavior
	// collapsible_host
	// system_vulnerability
	AttackTag string `json:"attack_tag"`
	// Occurrence time, accurate to milliseconds.
	OccurrenceTime int64 `json:"occur_time"`
	// Handling time, in milliseconds. This API is available only for handled alarms.
	HandleTime int `json:"handle_time"`
	// Processing status. Its value can be:
	// unhandled
	// handled
	HandleStatus string `json:"handle_status"`
	// Handling method. This API is available only for handled alarms. The options are as follows:
	// mark_as_handled
	// ignore
	// add_to_alarm_whitelist
	// add_to_login_whitelist
	// isolate_and_kill
	HandleMethod string `json:"handle_method"`
	// Remarks. This API is available only for handled alarms.
	Handler string `json:"handler"`
	// Supported processing operation
	OperateAcceptList string `json:"operate_accept_list"`
	// Operation details list (not displayed on the page)
	OperateDetailList []OperateDetailList `json:"operate_detail_list"`
	// Attack information, in JSON format.
	ForensicInfo interface{} `json:"forensic_info"`
	// Resource information
	ResourceInfo *ResourceInfo `json:"resource_info"`
	// Geographical location, in JSON format.
	GeoInfo interface{} `json:"geo_info"`
	// Malware information, in JSON format.
	MalwareInfo interface{} `json:"malware_info"`
	// Network information, in JSON format.
	NetworkInfo interface{} `json:"network_info"`
	// Application information, in JSON format.
	AppInfo interface{} `json:"app_info"`
	// System information, in JSON format.
	SystemInfo interface{} `json:"system_info"`
	// Extended event information, in JSON format
	ExtendInfo interface{} `json:"extend_info"`
	// Handling suggestions
	Recommendation string `json:"recommendation"`
	// Process information list
	ProcessInfoList []ProcessInfoList `json:"process_info_list"`
	// User information list
	UserInfoList []UserInfoList `json:"user_info_list"`
	// File information list
	FileInfoList []FileInfoList `json:"file_info_list"`
	// Brief description of the event.
	EventDetails string `json:"event_details"`
}

type OperateDetailList struct {
	// Agent ID
	AgentId string `json:"agent_id"`
	// Process ID
	ProcessPid int `json:"process_pid"`
	// Whether a process is a parent process
	IsParent bool `json:"is_parent"`
	// File hash
	FileHash string `json:"file_hash"`
	// File path
	FilePath string `json:"file_path"`
	// File attribute
	FileAttr string `json:"file_attr"`
	// Server private IP address
	PrivateIP string `json:"private_ip"`
	// Login source IP address
	LoginIP string `json:"login_ip"`
	// Login username
	LoginUserName string `json:"login_user_name"`
	// Alarm event keyword, which is used only for the alarm whitelist.
	Keyword string `json:"keyword"`
	// Alarm event hash, which is used only for the alarm whitelist.
	Hash string `json:"hash"`
}

type ResourceInfo struct {
	// User account ID
	DomainId string `json:"domain_id"`
	// Project ID
	ProjectId string `json:"project_id"`
	// Enterprise project ID. The value 0 indicates the default enterprise project.
	// To query all enterprise projects, set this parameter to all_granted_eps.
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// Region name
	RegionName string `json:"region_name"`
	// VPC ID
	VpcId string `json:"vpc_id"`
	// ECS ID
	EcsId string `json:"cloud_id"`
	// VM name
	VmName string `json:"vm_name"`
	// Specifies the VM UUID, that is, the server ID.
	VmUuid string `json:"vm_uuid"`
	// Container ID
	ContainerId string `json:"container_id"`
	// Image ID
	ImageId string `json:"image_id"`
	// Image name
	ImageName string `json:"image_name"`
	// Host attribute
	HostAttr string `json:"host_attr"`
	// Service
	Service string `json:"service"`
	// Microservice
	Microservice string `json:"micro_service"`
	// System CPU architecture
	Arch string `json:"sys_arch"`
	// OS bit version
	OsBit string `json:"os_bit"`
	// OS type
	OsType string `json:"os_type"`
	// OS name
	OsName string `json:"os_name"`
	// OS version
	OsVersion string `json:"os_version"`
}

type ProcessInfoList struct {
	// Process name
	ProcessName string `json:"process_name"`
	// Process file path
	ProcessPath string `json:"process_path"`
	// Process ID
	ProcessPid int `json:"process_pid"`
	// Process user ID
	ProcessUid int `json:"process_uid"`
	// Process username
	ProcessUsername string `json:"process_username"`
	// Process file command line
	ProcessCmdline string `json:"process_cmdline"`
	// Process file name
	ProcessFilename string `json:"process_filename"`
	// Process start time
	ProcessStartTime int64 `json:"process_start_time"`
	// Process group ID
	ProcessGid int `json:"process_gid"`
	// Valid process group ID
	ProcessEgid int `json:"process_egid"`
	// Valid process user ID
	ProcessEuid int `json:"process_euid"`
	// Parent process name
	ParentProcessName string `json:"parent_process_name"`
	// Parent process file path
	ParentProcessPath string `json:"parent_process_path"`
	// Parent process ID
	ParentProcessPid int `json:"parent_process_pid"`
	// Parent process user ID
	ParentProcessUid int `json:"parent_process_uid"`
	// Parent process file command line
	ParentProcessCmdline string `json:"parent_process_cmdline"`
	// Parent process file name
	ParentProcessFilename string `json:"parent_process_filename"`
	// Parent process start time
	ParentProcessStartTime int64 `json:"parent_process_start_time"`
	// Parent process group ID
	ParentProcessGid int `json:"parent_process_gid"`
	// Valid parent process group ID
	ParentProcessEgid int `json:"parent_process_egid"`
	// Valid parent process user ID
	ParentProcessEuid int `json:"parent_process_euid"`
	// Subprocess name
	ChildProcessName string `json:"child_process_name"`
	// Subprocess file path
	ChildProcessPath string `json:"child_process_path"`
	// Subprocess ID
	ChildProcessPid int `json:"child_process_pid"`
	// Subprocess user ID
	ChildProcessUid int `json:"child_process_uid"`
	// Subprocess file command line
	ChildProcessCmdline string `json:"child_process_cmdline"`
	// Subprocess file name
	ChildProcessFilename string `json:"child_process_filename"`
	// Subprocess start time
	ChildProcessStartTime int64 `json:"child_process_start_time"`
	// Subprocess group ID
	ChildProcessGid int `json:"child_process_gid"`
	// Valid subprocess group ID
	ChildProcessEgid int `json:"child_process_egid"`
	// Valid subprocess user ID
	ChildProcessEuid int `json:"child_process_euid"`
	// Virtualization command
	VirtCmd string `json:"virt_cmd"`
	// Virtualization process name
	VirtProcessName string `json:"virt_process_name"`
	// Escape mode
	EscapeMode string `json:"escape_mode"`
	// Commands executed after escape
	EscapeCmd string `json:"escape_cmd"`
	// Process startup file hash
	ProcessHash string `json:"process_hash"`
}

type UserInfoList struct {
	// User UID
	UserId int `json:"user_id"`
	// User GID
	UserGid int `json:"user_gid"`
	// User name
	UserName string `json:"user_name"`
	// User group name
	UserGroupName string `json:"user_group_name"`
	// User home directory
	UserHomeDir string `json:"user_home_dir"`
	// User login IP address
	LoginIP string `json:"login_ip"`
	// Service type. The options are as follows:
	// system
	// mysql
	// redis
	ServiceType string `json:"service_type"`
	// Login service port
	ServicePort int `json:"service_port"`
	// Login mode
	LoginMode int `json:"login_mode"`
	// Last login time
	LoginLastTime int64 `json:"login_last_time"`
	// Number of failed login attempts
	LoginFailCount int `json:"login_fail_count"`
	// Password hash
	PwdHash string `json:"pwd_hash"`
	// Masked password
	PwdWithFuzzing string `json:"pwd_with_fuzzing"`
	// Password age (days)
	PwdUsedDays int `json:"pwd_used_days"`
	// Minimum password validity period
	PwdMinDays int `json:"pwd_min_days"`
	// Maximum password validity period
	PwdMaxDays int `json:"pwd_max_days"`
	// Advance warning of password expiration (days)
	PwdWarnLeftDays int `json:"pwd_warn_left_days"`
}

type FileInfoList struct {
	// File path
	FilePath string `json:"file_path"`
	// File alias
	FileAlias string `json:"file_alias"`
	// File size
	FileSize int `json:"file_size"`
	// Time when a file was last modified
	FileMtime int64 `json:"file_mtime"`
	// Time when a file was last accessed
	FileAtime int64 `json:"file_atime"`
	// Time when the status of a file was last changed
	FileCtime int64 `json:"file_ctime"`
	// The hash value calculated using the SHA256 algorithm.
	FileHash string `json:"file_hash"`
	// File MD5
	FileMd5 string `json:"file_md5"`
	// File SHA256
	FileSha256 string `json:"file_sha256"`
	// File type
	FileType string `json:"file_type"`
	// File content
	FileContent string `json:"file_content"`
	// File attribute
	FileAttr string `json:"file_attr"`
	// File operation type
	FileOperation int `json:"file_operation"`
	// File action
	FileAction string `json:"file_action"`
	// Old/New attribute
	FileChangeAttr string `json:"file_change_attr"`
	// New file path
	FileNewPath string `json:"file_new_path"`
	// File description
	FileDesc string `json:"file_desc"`
	// File keyword
	FileKeyWord string `json:"file_key_word"`
	// Whether it is a directory
	IsDir bool `json:"is_dir"`
	// File handle information
	FdInfo string `json:"fd_info"`
	// Number of file handles
	FdCount int `json:"fd_count"`
}
