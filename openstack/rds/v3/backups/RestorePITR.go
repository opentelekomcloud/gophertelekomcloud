package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
)

type RestorePITROpts struct {
	// Specifies the restoration information.
	Source Source `json:"source"`
	// Specifies the restoration target.
	Target Target `json:"target"`
}

type Source struct {
	// Specifies the ID of the backup used to restore data. This parameter must be specified when the backup file is used for restoration.
	BackupID string `json:"backup_id" required:"false"`
	// Specifies the DB instance ID.
	InstanceID string `json:"instance_id" required:"true"`
	// Specifies the time point of data restoration in the UNIX timestamp. The unit is millisecond and the time zone is UTC.
	RestoreTime int64 `json:"restore_time,omitempty"`
	// Specifies the restoration mode. Enumerated values include:
	// backup: indicates using backup files for restoration. In this mode, type is not mandatory and backup_id is mandatory.
	// timestamp: indicates the point-in-time restoration mode. In this mode, type is mandatory and restore_time is no mandatory.
	Type string `json:"type" required:"true"`
}

type Target struct {
	// Specifies the ID of the DB instance to be restored to.
	InstanceID string `json:"instance_id" required:"true"`
}

func RestorePITR(c *golangsdk.ServiceClient, opts RestorePITROpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/recovery
	raw, err := c.Post(c.ServiceURL("instances", "recovery"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return "", err
	}

	var res instances.JobId
	err = extract.Into(raw.Body, &res)
	return res.JobId, err
}
