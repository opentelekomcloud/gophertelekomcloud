package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchQueryTaskOpts struct {
	Jobs    []string `json:"jobs"`
	PageReq PageReq  `json:"page_req,omitempty"`
}

type PageReq struct {
	// Current page number, which cannot exceed the maximum number of pages.
	// (Number of pages = Number of transferred job IDs/Number of tasks on each page)
	// Minimum value: 1.
	// Default value: 1
	CurPage int `json:"cur_page,omitempty"`
	// Number of items on each page. If this parameter is set to 0, all items are obtained.
	// Minimum value: 0
	// Maximum value: 100
	// Default value: 5
	PerPage int `json:"per_page,omitempty"`
}

func BatchListTaskDetails(client *golangsdk.ServiceClient, opts BatchQueryTaskOpts) (*BatchListTaskDetailsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/batch-detail
	raw, err := client.Post(client.ServiceURL("jobs", "batch-detail"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}

	var res BatchListTaskDetailsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BatchListTaskDetailsResponse struct {
	Count   int            `json:"count,omitempty"`
	Results []QueryJobResp `json:"results,omitempty"`
}

type QueryJobResp struct {
	// Task ID.
	Id string `json:"id,omitempty"`
	// Parent task ID.
	ParentId string `json:"parent_id,omitempty"`
	// Task name.
	Name string `json:"name,omitempty"`
	// Task status. Values:
	// CREATING: The task is being created.
	// CREATE_FAILED: The task fails to be created.
	// CONFIGURATION: The task is being configured.
	// STARTJOBING: The task is being started.
	// WAITING_FOR_START: The task is waiting to be started.
	// START_JOB_FAILED: The task fails to be started.
	// FULL_TRANSFER_STARTED: Full migration is in progress, and the DR scenario is initialized.
	// FULL_TRANSFER_FAILED: Full migration failed. Initialization failed in the DR scenario.
	// FULL_TRANSFER_COMPLETE: Full migration is complete, and the initialization is complete in the DR scenario.
	// INCRE_TRANSFER_STARTED: Incremental migration is being performed, and the DR task is in progress.
	// INCRE_TRANSFER_FAILED: Incremental migration fails and a DR exception occurs.
	// RELEASE_RESOURCE_STARTED: The task is being stopped.
	// RELEASE_RESOURCE_FAILED: Stop task failed.
	// RELEASE_RESOURCE_COMPLETE: The task is stopped.
	// CHANGE_JOB_STARTED: The task is being changed.
	// CHANGE_JOB_FAILED: Change task failed.
	// CHILD_TRANSFER_STARTING: The subtask is being started.
	// CHILD_TRANSFER_STARTED: The subtask is being migrated.
	// CHILD_TRANSFER_COMPLETE: The subtask migration is complete.
	// CHILD_TRANSFER_FAILED: Migrate subtask failed.
	// RELEASE_CHILD_TRANSFER_STARTED: The subtask is being stopped.
	// RELEASE_CHILD_TRANSFER_COMPLETE: The subtask is complete.
	Status string `json:"status,omitempty"`
	// Description.
	Description string `json:"description,omitempty"`
	// Creation time, in timestamp format.
	CreateTime string `json:"create_time,omitempty"`
	// Migration type. Values:
	// FULL_TRANS: full migration
	// INCR_TRANS: incremental migration
	// FULL_INCR_TRANS: full+incremental migration
	TaskType string `json:"task_type,omitempty"`
	// Source database information.
	SourceEndpoint Endpoint `json:"source_endpoint,omitempty"`
	// DMQ information body.
	DmqEndpoint Endpoint `json:"dmq_endpoint,omitempty"`
	// Information about the physical source database.
	SourceSharding []Endpoint `json:"source_sharding,omitempty"`
	// Information body of the destination database.
	TargetEndpoint Endpoint `json:"target_endpoint,omitempty"`
	// Network type. Values:
	// vpn
	// vpc
	// eip
	NetType string `json:"net_type,omitempty"`
	// Failure cause.
	FailedReason string `json:"failed_reason,omitempty"`
	// Replication instance information.
	InstInfo InstInfo `json:"inst_info,omitempty"`
	// Start time, in timestamp format.
	ActualStartTime string `json:"actual_start_time,omitempty"`
	// Full migration completion time, in timestamp format.
	FullTransferCompleteTime string `json:"full_transfer_complete_time,omitempty"`
	// Update time, in timestamp format.
	UpdateTime string `json:"update_time,omitempty"`
	// Task direction. Values:
	// up: The current cloud is the standby cloud in the DR and to-the-cloud scenarios.
	// down: The current cloud is the active cloud in the DR and out-of-cloud scenarios.
	// non-dbs: self-built databases.
	JobDirection string `json:"job_direction,omitempty"`
	// Migration scenario Values:
	// migration: real-time migration.
	// sync: real-time synchronization.
	// cloudDataGuard: real-time disaster recovery.
	DbUseType string `json:"db_use_type,omitempty"`
	// Whether the instance needs to be restarted.
	NeedRestart *bool `json:"need_restart,omitempty"`
	// Whether the destination instance is restricted to read-only.
	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`
	// Conflict policy. Values:
	// stop: The conflict fails.
	// overwrite: Conflicting data is overwritten.
	// ignore: The conflict is ignored.
	ConflictPolicy string `json:"conflict_policy,omitempty"`
	// DDL filtering policy. Values:
	// drop_database: Filters DDLs.
	// drop_databasefilter_all: Filters out all DLLs.
	// "": No filter.
	FilterDdlPolicy string `json:"filter_ddl_policy,omitempty"`
	// Migration speed limit.
	SpeedLimit []SpeedLimitInfo `json:"speed_limit,omitempty"`
	// Migration schemes. Values:
	// Replication: primary/standby replication.
	// Tungsten: parses logs.
	// PGBaseBackup: PostgreSQL backup.
	SchemaType string `json:"schema_type,omitempty"`
	// The number of nodes.
	NodeNum string `json:"node_num,omitempty"`
	// Whether to select objects.
	ObjectSwitch bool `json:"object_switch,omitempty"`
	// Main task ID
	MasterJobId string `json:"master_job_id,omitempty"`
	// Full snapshot mode.
	FullMode string `json:"full_mode,omitempty"`
	// Whether to migrate the structure.
	StructTrans *bool `json:"struct_trans,omitempty"`
	// Whether to migrate indexes.
	IndexTrans *bool `json:"index_trans,omitempty"`
	// Whether to replace the definer with the user of the destination database.
	ReplaceDefiner *bool `json:"replace_definer,omitempty"`
	// Whether to migrate users.
	MigrateUser *bool `json:"migrate_user,omitempty"`
	// Whether to perform database-level synchronization.
	SyncDatabase *bool `json:"sync_database,omitempty"`

	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`

	// Information about the root node database of the destination instance.
	TargetRootDb *DefaultRootDb `json:"target_root_db,omitempty"`
	// AZ where the node is located.
	AzCode string `json:"az_code,omitempty"`
	// VPC to which the node belongs.
	VpcId string `json:"vpc_id,omitempty"`
	// Subnet where the node is located.
	SubnetId string `json:"subnet_id,omitempty"`
	// Security group to which the node belongs.
	SecurityGroupId string `json:"security_group_id,omitempty"`
	// Whether the task is a multi-active DR task. The value is true when the task is a dual-active DR task.
	MultiWrite *bool `json:"multi_write,omitempty"`
	// Whether IPv6 is supported
	SupportIpV6 *bool `json:"support_ip_v6,omitempty"`
	// Inherited task ID, which is used for the Oracle_Mrskafka link.
	InheritId string `json:"inherit_id,omitempty"`
	// GTID set of breakpoints.
	Gtid string `json:"gtid,omitempty"`
	// Exception notification settings.
	AlarmNotify *QuerySmnInfoResp `json:"alarm_notify,omitempty"`
	// Whether the task is a cross-AZ synchronization task.
	IsMultiAz *bool `json:"is_multi_az"`
	// AZ name of the node.
	AzName string `json:"az_name"`
	// Primary AZ of the cross-AZ task.
	MasterAz string `json:"master_az"`
	// Standby AZ of the cross-AZ task.
	SlaveAz string `json:"slave_az"`
	// Primary/Standby role of a task.
	NodeRole string `json:"node_role"`
	// Start point of an incremental task.
	IncreStartPosition string `json:"incre_start_position,omitempty"`
}

type InstInfo struct {
	// Engine type. Values:
	// mysql
	// mongodb
	// cloudDataGuard-mysql
	// mysql-to-taurus
	// postgresql
	EngineType string `json:"engine_type,omitempty"`
	// DB instance type. Values:
	// high
	InstType string `json:"inst_type,omitempty"`
	// Private IP address of the replication instance.
	Ip string `json:"ip,omitempty"`
	// EIP of the replication instance.
	PublicIp string `json:"public_ip,omitempty"`
	// Scheduled start time of a replication instance task.
	StartTime string `json:"start_time,omitempty"`
	// Replication instance status. Values:
	// active
	// deleted
	Status string `json:"status,omitempty"`
	// Storage space of a replication instance.
	VolumeSize int `json:"volume_size,omitempty"`
}

type DefaultRootDb struct {
	// Database name.
	DbName string `json:"db_name,omitempty"`
	// Encoding format
	DbEncoding string `json:"db_encoding,omitempty"`
}

type QuerySmnInfoResp struct {
	TopicName     string          `json:"topic_name"`
	DelayTime     int             `json:"delay_time"`
	RtoDisplay    int             `json:"rto_delay"`
	RpoDisplay    int             `json:"rpo_delay"`
	AlarmToUser   *bool           `json:"alarm_to_user"`
	Subscriptions []Subscriptions `json:"subscriptions"`
}

type Subscriptions struct {
	Endpoints []string `json:"endpoints"`
	Protocol  string   `json:"protocol"`
}
