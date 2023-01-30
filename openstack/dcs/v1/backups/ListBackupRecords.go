package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListBackupOpts struct {
	// Start sequence number of the backup record that is to be queried. By default, this parameter is set to 1.
	Start int32 `q:"start"`
	// Start time of the period to be queried. Format: yyyyMMddHHmmss, for example, 20170718235959.
	BeginTime string `q:"begin_time"`
	// End time of the period to be queried. Format: yyyyMMddHHmmss, for example, 20170718235959.
	EndTime string `q:"end_time"`
	// Number of backup records displayed on each page. The minimum value of this parameter is 1.
	// If this parameter is not set, 10 backup records are displayed on each page by default.
	Limit int32 `q:"limit"`
}

func ListBackupRecords(client *golangsdk.ServiceClient, instancesId string, opts ListBackupOpts) (*ListBackupRecordsResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("instances", instancesId, "backups")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListBackupRecordsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListBackupRecordsResponse struct {
	// Number of obtained backup records.
	TotalNum int `json:"total_num"`
	// Array of the backup records. For details about backup_record_response,
	BackupRecordResponse []BackupRecordResponse `json:"backup_record_response"`
}

type BackupRecordResponse struct {
	// ID of the backup record
	BackupId string `json:"backup_id"`
	// Time segment in which DCS instance backup was performed
	Period string `json:"period"`
	// Name of the backup record
	BackupName string `json:"backup_name"`
	// DCS instance ID
	InstanceId string `json:"instance_id"`
	// Size of the backup file. Unit: byte.
	Size int64 `json:"size"`
	// Backup type. Options:
	// manual: manual backup
	// auto: automatic backup
	BackupType string `json:"backup_type"`
	// Time at which the backup task is created
	CreatedAt string `json:"created_at"`
	// Time at which DCS instance backup is completed
	UpdatedAt string `json:"updated_at"`
	// Backup progress
	Progress string `json:"progress"`
	// Error code returned if DCS instance backup fails.
	ErrorCode string `json:"error_code"`
	// Description of DCS instance backup
	Remark string `json:"remark"`
	// Backup status. Options:
	// waiting: DCS instance restoration is waiting to begin.
	// backuping: DCS instance backup is in progress.
	// succeed: DCS instance backup succeeded.
	// failed: DCS instance backup failed.
	// expired: The backup file expires.
	// deleted: The backup file has been deleted manually.
	Status string `json:"status"`
	// An indicator of whether restoration is supported. Options: TRUE or FALSE.
	IsSupportRestore string `json:"is_support_restore"`
	// Time at which the backup starts.
	ExecutionAt string `json:"execution_at"`
	// Backup format.
	BackupFormat string `json:"backup_format"`
}
