package migration

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateMigrationTaskOpts struct {
	// Name of the migration task.
	TaskName string `json:"task_name"`
	// Description of the migration task.
	Description string `json:"description,omitempty"`
	// Mode of the migration.
	// Options:
	// 	backupfile_import: indicates importing backup files.
	// 	online_migration: indicates migrating data online.
	MigrationType string `json:"migration_type"`
	// Type of the migration.
	// Options:
	// 	full_amount_migration: indicates a full migration.
	// 	incremental_migration: indicates an incremental migration.
	MigrationMethod string `json:"migration_method"`
	// Backup files to be imported when the migration mode is importing backup files.
	BackupFiles BackupFilesBody `json:"backup_files,omitempty"`
	// Type of the network for communication between the source and
	// destination Redis when the migration mode is online data migration.
	// Options:
	// 	vpc
	// 	vpn
	NetworkType string `json:"network_type,omitempty"`
	// Source Redis information. This parameter is mandatory when the migration mode is online data migration.
	SourceInstance SourceInstanceBody `json:"source_instance,omitempty"`
	// Destination Redis instance information.
	TargetInstance TargetInstanceBody `json:"target_instance"`
}

type BackupFilesBody struct {
	// Data source. Currently, only OBS buckets are supported. The value is self_build_obs.
	FileSource string `json:"file_source,omitempty"`
	// OBS bucket name.
	BucketName string `json:"bucket_name"`
	// List of backup files to be imported.
	Files []Files `json:"files"`
}

type Files struct {
	// Name of a backup file.
	FileName string `json:"file_name"`
	// File size in bytes.
	Size string `json:"size,omitempty"`
	// Time when the file is last modified. The format is YYYY-MM-DD HH:MM:SS.
	UpdateAt string `json:"update_at,omitempty"`
}

type SourceInstanceBody struct {
	// Source Redis name (specified in the source_instance parameter).
	Addrs string `json:"addrs"`
	// Redis password. If a password is set, this parameter is mandatory.
	Password string `json:"password,omitempty"`
}

type TargetInstanceBody struct {
	// Destination Redis instance ID (mandatory in the target_instance parameter).
	Id string `json:"id"`
	// Destination Redis instance name (specified in the target_instance parameter).
	Name string `json:"name,omitempty"`
	// Redis password. If a password is set, this parameter is mandatory.
	Password string `json:"password,omitempty"`
}

func CreateMigrationTask(client *golangsdk.ServiceClient, opts CreateMigrationTaskOpts) (*CreateMigrationTaskResponse, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CreateMigrationTaskResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateMigrationTaskResponse struct {
	// ID of the migration task.
	Id string `json:"id,omitempty"`
	// Name of the migration task.
	Name string `json:"name,omitempty"`
	// Migration task status. The value can be:
	// 	SUCCESS: Migration succeeded.
	// 	FAILED: Migration failed.
	// 	MIGRATING: Migration is in progress.
	// 	TERMINATED: Migration has been stopped.
	// 	TERMINATING: Migration is being stopped.
	// 	RUNNING: The migration task has been created and is waiting to be executed.
	// 	CREATING: The migration task is being created.
	// 	FULLMIGRATING: Full migration is in progress.
	// 	INCRMIGEATING: Incremental migration is in progress.
	// 	ERROR: faulty
	// 	DELETED: faulty
	// 	RELEASED: automatically released
	// 	MIGRATION_SUCCESS: The migration is successful, and resources are to be cleared.
	// 	MIGRATION_FAILED: The migration failed, and resources are to be cleared.
	Status string `json:"status,omitempty"`
}
