package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BackupDatabase struct {
	// Specifies the names of self-built databases.
	Name string `json:"name"`
}

type CreateOpts struct {
	// Specifies the DB instance ID.
	InstanceID string `json:"instance_id" required:"true"`
	// Specifies the backup name. It must be 4 to 64 characters in length and start with a letter. It is case-sensitive and can contain only letters, digits, hyphens (-), and underscores (_).
	// The backup name must be unique.
	Name string `json:"name" required:"true"`
	// Specifies the backup description. It contains a maximum of 256 characters and cannot contain the following special characters: >!<"&'=
	Description string `json:"description,omitempty"`
	// Specifies a list of self-built Microsoft SQL Server databases that are partially backed up. (Only Microsoft SQL Server support partial backups.)
	Databases []BackupDatabase `json:"databases,omitempty"`
}

func Create(c *golangsdk.ServiceClient, opts CreateOpts) (*Backup, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/backups
	raw, err := c.Post(c.ServiceURL("backups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	var res Backup
	err = extract.IntoStructPtr(raw.Body, &res, "backup")
	return &res, err
}
