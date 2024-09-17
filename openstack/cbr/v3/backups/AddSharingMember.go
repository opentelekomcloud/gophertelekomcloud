package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type MembersOpts struct {
	// Project IDs of the backup share members to be added
	Members []string `json:"members" required:"true"`
}

func AddSharingMember(client *golangsdk.ServiceClient, backupID string, opts MembersOpts) ([]Members, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/backups/{backup_id}/members
	raw, err := client.Post(client.ServiceURL("backups", backupID, "members"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res []Members
	err = extract.IntoSlicePtr(raw.Body, &res, "members")
	return res, err
}

type Members struct {
	// Backup sharing status
	Status string `json:"status"`
	// Backup sharing time
	CreatedAt string `json:"created_at"`
	// Update time
	UpdatedAt string `json:"updated_at"`
	// Backup ID
	BackupId string `json:"backup_id"`
	// ID of the image created by using the accepted shared backup
	ImageId string `json:"image_id"`
	// ID of the project with which the backup is shared
	DestProjectId string `json:"dest_project_id"`
	// Replication record ID
	VaultId string `json:"vault_id"`
	// ID of the shared record
	ID string `json:"id"`
}
