package event

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
	// Handling method. Its value can be:
	// mark_as_handled
	// ignore
	// add_to_alarm_whitelist
	// add_to_login_whitelist
	// isolate_and_kill
	// unhandle
	// do_not_ignore
	// remove_from_alarm_whitelist
	// remove_from_login_whitelist
	// do_not_isolate_or_kill
	OperateType string `json:"operate_type" required:"true"`
	// Remarks. This API is available only for handled alarms.
	Handler string `json:"handler,omitempty"`
	// Operated event list
	OperateEventList []OperateRequestInfo `json:"operate_event_list" required:"true"`
}

type OperateRequestInfo struct {
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
	EventClassID string `json:"event_class_id" required:"true"`
	// Event ID
	EventID string `json:"event_id" required:"true"`
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
	EventType int `json:"event_type" required:"true"`
	// Occurrence time, accurate to milliseconds.
	OccurrenceTime int64 `json:"occur_time" required:"true"`
	// Operation details list. If operate_type is set to add_to_alarm_whitelist or remove_from_alarm_whitelist,
	// keyword and hash are mandatory. If operate_type is set to add_to_login_whitelist or remove_from_login_whitelist,
	// the login_ip, private_ip, and login_user_name parameters are mandatory.
	// If operate_type is set to isolate_and_kill or do_not_isolate_or_kill, the agent_id, file_hash, file_path,
	// and process_pid parameters are mandatory. In other cases, the parameters are optional.
	OperateDetailList []DetailRequestInfo `json:"operate_detail_list" required:"true"`
}

type DetailRequestInfo struct {
	// Agent ID
	AgentID string `json:"agent_id,omitempty"`
	// Process ID
	ProcessPID *int `json:"process_pid,omitempty"`
	// File hash
	FileHash string `json:"file_hash,omitempty"`
	// File path
	FilePath string `json:"file_path,omitempty"`
	// File attribute
	FileAttr string `json:"file_attr,omitempty"`
	// Alarm event keyword, which is used only for the alarm whitelist.
	Keyword string `json:"keyword,omitempty"`
	// Alarm event hash, which is used only for the alarm whitelist.
	Hash string `json:"hash,omitempty"`
	// Server private IP address
	PrivateIP string `json:"private_ip,omitempty"`
	// Login source IP address
	LoginIP string `json:"login_ip,omitempty"`
	// Login username
	LoginUserName string `json:"login_user_name,omitempty"`
}

func Operate(client *golangsdk.ServiceClient, opts CreateOpts, epsID string) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	requestPath := client.ServiceURL("event", "operate")
	if epsID != "" {
		requestPath += "?enterprise_project_id=" + epsID
	}
	// POST /v5/{project_id}/event/operate
	_, err = client.Post(requestPath, b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"region": client.RegionID},
	})
	return err
}
