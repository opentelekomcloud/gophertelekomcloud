package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/instances"
)

type ListBackupsOpts struct {
	// Index offset.
	// If offset is set to N, the resource query starts from the N+1 piece of data. The default value is 0, indicating that the query starts from the first piece of data.
	// The value must be a positive integer.
	Offset int `q:"offset"`
	// Maximum number of specifications that can be queried
	// Value range: 1-100
	// If this parameter is not transferred, the first 100 pieces of specification information can be queried by default.
	Limit int `q:"limit"`
	// Instance id.
	InstanceId string `q:"instance_id"`
	// Specifies the backup ID.
	BackupId string `q:"backup_id"`
	// Specifies the backup type.
	// Auto: indicates automated full backup.
	// Manual indicates manual full backup.
	BackupType string `q:"backup_type"`
	// Specifies the start time of the query. The format is yyyy-mm-dd hh:mm:ss. The value is in UTC format.
	BeginTime string `q:"begin_time"`
	// Specifies the end time of the query. The format is "yyyy-mm-dd hh:mm:ss". The value is in UTC format.
	EndTime string `q:"end_time"`
	// Specifies the DB instance mode.
	// Valid value:
	// Sharding
	// ReplicaSet
	// Single
	Mode string `q:"mode"`
}

func List(client *golangsdk.ServiceClient, opts ListBackupsOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("backups").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/backups
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	Backups    []BackupsResponse `json:"backups"`
	TotalCount int               `json:"total_count"`
}

type BackupsResponse struct {
	// Indicates the backup ID.
	ID string `json:"id"`
	// Indicates the backup name.
	Name string `json:"name"`
	// Indicates the ID of the DB instance from which the backup was created.
	InstanceId string `json:"instance_id"`
	// Indicates the name of the DB instance for which the backup is created.
	InstanceName string `json:"instance_name"`
	// Indicates the database version.
	Datastore instances.DataStore `json:"datastore"`
	// Indicates the backup type.
	Type string `json:"type"`
	// Indicates the backup start time. The format of the start time is yyyy-mm-dd hh:mm:ss. The value is in UTC format.
	BeginTime string `json:"begin_time"`
	// Indicates the backup end time. The format of the end time is yyyy-mm-dd hh:mm:ss. The value is in UTC format.
	EndTime string `json:"end_time"`
	// Indicates the backup status. Valid value:
	// BUILDING: Backup in progress
	// COMPLETED: Backup completed
	// FAILED: Backup failed
	// DISABLED: Backup being deleted
	Status string `json:"status"`
	// Indicates the backup size in KB.
	Size int64 `json:"size"`
	// Indicates the backup description.
	Description string `json:"description"`
}
