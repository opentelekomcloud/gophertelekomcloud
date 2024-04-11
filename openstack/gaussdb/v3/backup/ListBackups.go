package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BackupListOpts struct {
	// Instance ID
	InstanceId string `q:"instance_id"`
	// Backup ID
	BackupId string `q:"backup_id"`
	// Backup type
	// auto: automated full backup
	// manual: manual full backup
	BackupType string `q:"backup_type"`
	// Index offset. If offset is set to N, the resource query starts from the N+1 piece of data.
	// The default value is 0, indicating that the query starts from the first piece of data. The value must be a positive integer.
	Offset string `q:"offset"`
	// Number of records to be queried. The default value is 100.
	// The value must be a positive integer. The minimum value is 1 and the maximum value is 100.
	Limit string `q:"limit"`
	// Query start time. The format is "yyyy-mm-ddThh:mm:ssZ".
	BeginTime string `q:"begin_time"`
	// Query end time. The format is "yyyy-mm-ddThh:mm:ssZ" and the end time must be later than the start time.
	EndTime string `q:"end_time"`
}

func ListBackups(client *golangsdk.ServiceClient, opts BackupListOpts) (*BackupListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("backups").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/mysql/v3/{project_id}/backups
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res BackupListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BackupListResponse struct {
	// Backup information
	Backups []Backups `json:"backups"`
	// Total number of backup files
	TotalCount int `json:"total_count"`
}

type Backups struct {
	// Backup ID
	Id string `json:"id"`
	// Backup name
	Name string `json:"name"`
	// Backup start time in the "yyyy-mm-ddThh:mm:ssZ" format.
	BeginTime string `json:"begin_time"`
	// Backup end time in the "yyyy-mm-ddThh:mm:ssZ" format.
	EndTime string `json:"end_time"`
	// Backup status. Value:
	// BUILDING: Backup in progress
	// COMPLETED: Backup completed
	// FAILED: Backup failed
	// AVAILABLE: Backup available
	Status string `json:"status"`
	// Backup duration in minutes
	TakeUpTime int32 `json:"take_up_time"`
	// Backup type
	// auto: automated full backup
	// manual: manual full backup
	Type string `json:"type"`
	// Backup size in MB.
	Size int64 `json:"size"`
	// Database information
	Datastore MysqlDatastore `json:"datastore"`
	// Instance ID
	InstanceId string `json:"instance_id"`
	// Backup level. This parameter is returned when the level-1 backup function is enabled. Value:
	// 1: level-1 backup
	// 2: level-2 backup
	// 0: Backup being created or creation failed
	BackupLevel string `json:"backup_level"`
	// Description of the backup file
	Description string `json:"description"`
}

type MysqlDatastore struct {
	// DB engine. Currently, only gaussdb-mysql is supported.
	Type string `json:"type"`
	// DB engine version
	Version string `json:"version"`
}
