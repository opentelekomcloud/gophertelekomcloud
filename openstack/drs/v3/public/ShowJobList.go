package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type QueryJobsOpts struct {
	// Current page. Set the value to 0 to obtain all items.
	// Default value: 1
	CurPage int32 `json:"cur_page"`
	// Number of records on each page.
	// Default value: 10
	// Minimum value: 0
	// Maximum value: 100
	PerPage int32 `json:"per_page"`
	// The migration scenario. The value can be migration (real-time migration), sync (real-time synchronization),
	// or cloudDataGuard (real-time disaster recovery).
	// Values:
	// migration
	// sync
	// cloudDataGuard
	DbUseType string `json:"db_use_type"`
	// The engine type. The value can be mysql (used for migration or synchronization), mongodb (used for migration),
	// mysql-to-taurus (used for synchronization from MySQL to GaussDB(for MySQL) primary/standby),
	// cloudDataGuard-mysql (used for DR), and postgresql (used for PostgreSQL synchronization).
	// Default value: mysql
	// Values:
	// mysql
	// mongodb
	// cloudDataGuard-mysql
	// mysql-to-taurus
	// postgresql
	EngineType string `json:"engine_type,omitempty"`
	// Name or ID.
	Name string `json:"name,omitempty"`
	// Network type. Values:
	// vpn
	// vpc
	// eip
	NetType string `json:"net_type,omitempty"`
	// Service name.
	ServiceName string `json:"service_name,omitempty"`
	// Status. The value can be CREATING, CREATE_FAILED, or CONFIGURATION.
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
	// Values:
	// CREATING
	// CREATE_FAILED
	// CONFIGURATION
	// STARTJOBING
	// WAITING_FOR_START
	// START_JOB_FAILED
	// FULL_TRANSFER_STARTED
	// FULL_TRANSFER_FAILED
	// FULL_TRANSFER_COMPLETE
	// INCRE_TRANSFER_STARTED
	// INCRE_TRANSFER_FAILED
	// RELEASE_RESOURCE_STARTED
	// RELEASE_RESOURCE_FAILED
	// RELEASE_RESOURCE_COMPLETE
	// CHANGE_JOB_STARTED
	// CHANGE_JOB_FAILED
	// CHILD_TRANSFER_STARTING
	// CHILD_TRANSFER_STARTED
	// CHILD_TRANSFER_COMPLETE
	// CHILD_TRANSFER_FAILED
	// RELEASE_CHILD_TRANSFER_STARTED
	// RELEASE_CHILD_TRANSFER_COMPLETE
	Status string `json:"status,omitempty"`
	// Tags.
	Tags map[string]string `json:"tags,omitempty"`
}

func ShowJobList(client *golangsdk.ServiceClient, opts QueryJobsOpts) (*ShowJobListResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs
	raw, err := client.Post(client.ServiceURL("jobs"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res ShowJobListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ShowJobListResponse struct {
	TotalRecord int32     `json:"total_record,omitempty"`
	Jobs        []JobInfo `json:"jobs,omitempty"`
}

type JobInfo struct {
	// Task ID.
	Id string `json:"id"`
	// Task name.
	Name string `json:"name"`
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
	Status string `json:"status"`
	// Task description.
	Description string `json:"description"`
	// Time when a task is created
	CreateTime string `json:"create_time"`
	// Engine type. Values:
	// cloudDataGuard-cassandra
	// cloudDataGuard-ddm
	// cloudDataGuard-taurus-to-mysql
	// cloudDataGuard-mysql
	// cloudDataGuard-mysql-to-taurus
	EngineType string `json:"engine_type"`
	// Network type. Values:
	// vpn
	// vpc
	// eip
	NetType string `json:"net_type"`
	// Billing tag.
	BillingTag bool `json:"billing_tag"`
	// Migration direction. Values:
	// up
	// down
	JobDirection string `json:"job_direction"`
	// Migration scenario. Values:
	// migration: real-time migration.
	// sync: real-time synchronization.
	// cloudDataGuard: real-time disaster recovery.
	DbUseType string `json:"db_use_type"`
	// Migration type. Values:
	// FULL_TRANS: full migration
	// FULL_INCR_TRANS: full+incremental migration
	// INCR_TRANS: incremental migration
	TaskType string `json:"task_type"`
	// Subtask information body.
	Children []ChildrenJobInfo `json:"children,omitempty"`
	// Whether the framework is a new framework.
	NodeNewFramework bool `json:"node_newFramework"`
}

type ChildrenJobInfo struct {
	// Billing tag.
	BillingTag bool `json:"billing_tag"`
	// Time when a task is created
	CreateTime string `json:"create_time"`
	// Replication scenario. Values:
	// migration: real-time migration.
	// sync: real-time synchronization.
	// cloudDataGuard: real-time disaster recovery.
	DbUseType string `json:"db_use_type"`
	// Task description.
	Description string `json:"description"`
	// Engine type. Values:
	// cloudDataGuard-cassandra
	// cloudDataGuard-ddm
	// cloudDataGuard-taurus-to-mysql
	// cloudDataGuard-mysql
	// cloudDataGuard-mysql-to-taurus
	EngineType string `json:"engine_type"`
	// Task failure cause.
	ErrorMsg string `json:"error_msg"`
	// Task ID.
	Id string `json:"id"`
	// Migration direction. Values:
	// up: The current cloud is the standby cloud in the DR and to-the-cloud scenarios.
	// down: The current cloud is the active cloud in the DR and out-of-cloud scenarios.
	// non-dbs: self-built databases.
	JobDirection string `json:"job_direction"`
	// Task name.
	Name string `json:"name"`
	// Network type. Values:
	// vpc
	// vpn
	// eip
	NetType string `json:"net_type"`
	// New framework
	NodeNewFramework bool `json:"node_newFramework"`
	// Task status.
	Status string `json:"status"`
	// Migration type. Values:
	// FULL_TRANS: full migration
	// FULL_INCR_TRANS: full+incremental migration
	// INCR_TRANS: incremental migration
	TaskType string `json:"task_type"`
}
