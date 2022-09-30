package public

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"

type BatchModifyJobReq struct {
	Jobs []ModifyJobReq `json:"jobs"`
}

type ModifyJobReq struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Task description. This parameter is mandatory when you modify the task description.
	// Minimum length: 0 character
	// Maximum length: 256
	Description string `json:"description,omitempty"`
	// Task name.
	Name string `json:"name,omitempty"`
	// Set exception notification.
	AlarmNotify AlarmNotifyInfo `json:"alarm_notify,omitempty"`
	// Task mode. FULL_TRANS: full. FULL_INCR_TRANS: full + incremental. INCR_TRANS: incremental.
	// Values:
	// FULL_TRANS
	// INCR_TRANS
	// FULL_INCR_TRANS
	TaskType string `json:"task_type,omitempty"`
	// Source database information. This parameter is mandatory for calling the API after the connection test.
	SourceEndpoint Endpoint `json:"source_endpoint,omitempty"`
	// Destination database information. This parameter is mandatory for calling the API after the connection test.
	TargetEndpoint Endpoint `json:"target_endpoint,omitempty"`
	// Node specification type. This parameter is mandatory when this API is invoked to modify a task after the connection test.
	// Default value: high
	// Values: high
	NodeType string `json:"node_type,omitempty"`
	// Engine type. This parameter is mandatory when this API is invoked to modify a task after the connection test.
	// mysql: used for migration and synchronization of MySQL databases
	// mongodb: used for migration.
	// cloudDataGuard-mysql: used for disaster recovery
	// mysql-to-taurus: used for synchronization from MySQL to GaussDB(for MySQL) primary/standby.
	// postgresql: used for PostgreSQL synchronization.
	// Values:
	// mysql
	// mongodb
	// cloudDataGuard-mysql
	// mysql-to-taurus
	// postgresql
	EngineType string `json:"engine_type,omitempty"`
	// Network type. This parameter is mandatory after the connection test. Values:
	// vpn
	// vpc
	// eip
	NetType string `json:"net_type,omitempty"`
	// Whether to save the database information. This parameter is mandatory when the API is called after the connection test.
	StoreDbInfo bool `json:"store_db_info,omitempty"`
	// 是否为重建任务。
	IsRecreate bool `json:"is_recreate,omitempty"`
	// The migration direction. The value can be up (to the cloud and current cloud as standby in disaster recovery),
	// down (out of cloud and current cloud as active in disaster recovery), or non-dbs (for self-built databases).
	// Values:
	// up
	// down
	// non-dbs
	JobDirection string `json:"job_direction,omitempty"`
	// Whether the destination DB instance can be read-only.
	IsTargetReadonly bool `json:"is_target_readonly,omitempty"`
	// Whether to migrate all Definers to the user. MySQL databases support this setting.
	// This parameter is mandatory when this API is invoked to modify a task after the connection test. Values:
	// true: The Definers of all source database objects will be migrated to the user.
	// Other users do not have permissions on database objects unless they are authorized.
	// false: The Definers of all source database objects will not be changed.
	// You need to migrate all accounts and permissions of the source database in the next step.
	ReplaceDefiner bool `json:"replace_definer,omitempty"`
	// Specifies the tag information.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
	// Migration type. The options are migration (real-time migration),
	// sync (real-time synchronization), and cloudDataGuard (real-time DR).
	// Values:
	// migration
	// sync
	// cloudDataGuard
	DbUseType string `json:"db_use_type,omitempty"`
	// Product ID.
	ProductId string `json:"product_id,omitempty"`
}

type AlarmNotifyInfo struct {
	// Subscription delay, in seconds.
	// Minimum value: 1
	// Maximum value: 3600
	// Default value: 0
	DelayTime int64 `json:"delay_time,omitempty"`
	// RTO delay.
	// Minimum value: 1
	// Maximum value: 3600
	// Default value: 0
	RtoDelay int64 `json:"rto_delay,omitempty"`
	// RPO delay.
	// Minimum value: 1
	// Maximum value: 3600
	// Default value: 0
	RpoDelay int64 `json:"rpo_delay,omitempty"`
	// Whether to notify users of alarms. The default value is false.
	AlarmToUser bool `json:"alarm_to_user,omitempty"`
	// Receiving method and message body. Up to two receiving modes and message bodies are supported.
	Subscriptions []SubscriptionInfo `json:"subscriptions,omitempty"`
}

type SubscriptionInfo struct {
	// List of mobile numbers or email addresses.
	// Use commas (,) to separate multiple mobile numbers or email addresses.
	// Up to 10 mobile numbers or email addresses are supported.
	Endpoints []string `json:"endpoints,omitempty"`
	// Receiving method. Values:
	// sms: SMS message
	// email: email.
	Protocol string `json:"protocol,omitempty"`
}

// BatchJobsResponse

// PUT /v3/{project_id}/jobs/batch-modification
