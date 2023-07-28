package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"net/url"
)

type ListOpts struct {
	// Specifies the DB instance ID.
	InstanceID string `q:"instance_id"`
	// Specifies the backup ID.
	BackupID string `q:"backup_id"`
	// Specifies the backup type. Value:
	//
	// auto: automated full backup
	// manual: manual full backup
	// fragment: differential full backup
	// incremental: automated incremental backup
	BackupType string `q:"backup_type"`
	// Specifies the index position. If offset is set to N, the resource query starts from the N+1 piece of data. The value is 0 by default, indicating that the query starts from the first piece of data. The value must be a positive number.
	Offset int `q:"offset"`
	// Specifies the number of records to be queried. The default value is 100. The value cannot be a negative number. The minimum value is 1 and the maximum value is 100.
	Limit int `q:"limit"`
	// Specifies the start time for obtaining the backup list. The format of the start time is "yyyy-mm-ddThh:mm:ssZ".
	// T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset.
	// NOTE:
	// When begin_time is not empty, end_time is mandatory.
	BeginTime string `q:"begin_time"`
	// Specifies the end time for obtaining the backup list. The format of the end time is "yyyy-mm-ddThh:mm:ssZ" and the end time must be later than the start time.
	// T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset.
	// NOTE:
	// When end_time is not empty, begin_time is mandatory.
	EndTime string `q:"end_time"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Backup, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/backups
	raw, err := client.Get(client.ServiceURL("backups")+q.String(), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res []Backup
	err = extract.IntoSlicePtr(raw.Body, &res, "backups")
	return res, err
}
