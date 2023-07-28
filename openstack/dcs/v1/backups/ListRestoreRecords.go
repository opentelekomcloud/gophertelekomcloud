package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListRestoreRecords(client *golangsdk.ServiceClient, instancesId string, opts ListBackupOpts) (*ListRestoreRecordsResponse, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("instances", instancesId, "restores")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListRestoreRecordsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListRestoreRecordsResponse struct {
	// Array of the restoration records.
	RestoreRecordResponse []InstanceRestoreInfo `json:"restore_record_response"`
	// Number of obtained backup records.
	TotalNum int `json:"total_num"`
}

type InstanceRestoreInfo struct {
	// ID of the backup record
	BackupId string `json:"backup_id"`
	// ID of the restoration record
	RestoreId string `json:"restore_id"`
	// Name of the backup record
	BackupName string `json:"backup_name"`
	// Time at which DCS instance restoration completed
	UpdatedAt string `json:"updated_at"`
	// Description of DCS instance restoration
	RestoreRemark string `json:"restore_remark"`
	// Time at which the restoration task is created
	CreatedAt string `json:"created_at"`
	// Restoration progress
	Progress string `json:"progress"`
	// Error code returned if DCS instance restoration fails.
	ErrorCode string `json:"error_code"`
	// Name of the restoration record
	RestoreName string `json:"restore_name"`
	// Description of DCS instance backup
	BackupRemark string `json:"backup_remark"`
	// Restoration status
	// waiting: DCS instance restoration is waiting to begin.
	// restoring: DCS instance restoration is in progress.
	// succeed: DCS instance restoration succeeded.
	// failed: DCS instance restoration failed.
	Status string `json:"status"`
	// Source instance ID.
	SourceInstanceID string `json:"sourceInstanceID"`
	// Source instance name.
	SourceInstanceName string `json:"sourceInstanceName"`
}
