package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

type ListRestoreTimesOpts struct {
	// Specifies the DB instance ID.
	InstanceId string `json:"-"`
	// Specifies the date to be queried. The value is in the yyyy-mm-dd format, and the time zone is UTC.
	Date string `json:"date,omitempty"`
}

func ListRestoreTimes(client *golangsdk.ServiceClient, opts ListRestoreTimesOpts) ([]RestoreTime, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/restore-time
	raw, err := client.Get(client.ServiceURL("instances", opts.InstanceId, "restore-time")+q.String(), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res []RestoreTime
	err = extract.IntoSlicePtr(raw.Body, &res, "restore_time")
	return res, err
}

type RestoreTime struct {
	// Indicates the start time of the restoration time range in the UNIX timestamp format. The unit is millisecond and the time zone is UTC.
	StartTime int64 `json:"start_time"`
	// Indicates the end time of the restoration time range in the UNIX timestamp format. The unit is millisecond and the time zone is UTC.
	EndTime int64 `json:"end_time"`
}
