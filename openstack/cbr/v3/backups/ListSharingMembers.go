package backups

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListMemberOpts struct {
	ID             string
	CheckpointID   string `q:"checkpoint_id"`
	DedicatedCloud bool   `q:"dec"`
	EndTime        string `q:"end_time"`
	ImageType      string `q:"image_type"`
	Limit          string `q:"limit"`
	Marker         string `q:"marker"`
	MemberStatus   string `q:"member_status"`
	Name           string `q:"name"`
	Offset         string `q:"offset"`
	OwningType     string `q:"own_type"`
	ParentID       string `q:"parent_id"`
	ResourceAZ     string `q:"resource_az"`
	ResourceID     string `q:"resource_id"`
	ResourceName   string `q:"resource_name"`
	ResourceType   string `q:"resource_type"`
	Sort           string `q:"sort"`
	StartTime      string `q:"start_time"`
	Status         string `q:"status"`
	UsedPercent    string `q:"used_percent"`
	VaultID        string `q:"vault_id"`
}

func ListSharingMembers(client *golangsdk.ServiceClient, backupId string, opts ListMemberOpts) ([]Member, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("backups", backupId, "members").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return MemberPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractMembers(pages)
}

type MemberPage struct {
	pagination.NewSinglePageBase
}

func ExtractMembers(r pagination.NewPage) ([]Member, error) {
	var s struct {
		Members []Member `json:"members"`
	}
	err := extract.Into(bytes.NewReader((r.(MemberPage)).Body), &s)
	return s.Members, err
}

type Member struct {
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
	// ID of the vault where the shared backup is stored
	VaultId string `json:"vault_id"`
	// ID of the shared record
	ID string `json:"id"`
}
