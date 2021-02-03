package backups

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type GetResult struct {
	golangsdk.Result
}

type BackupStatus string

type RestoreMode string

type BackupExtendInfo struct {
	AutoTrigger bool   `json:"auto_trigger"`
	Bootable    bool   `json:"bootable"`
	Incremental bool   `json:"incremental"`
	SnapshotID  string `json:"snapshot_id"`
	SupportID   string `json:"support_id"`

	OSImagesData      string `json:"os_images_data"`
	ContainSystemDisk bool   `json:"contain_system_disk"`
	Encrypted         bool   `json:"encrypted"`
	SystemDisk        bool   `json:"system_disk"`

	SupportedRestoreMode RestoreMode `json:"supported_restore_mode"`
}

type ImageType string

type BackupResp struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CheckpointID string    `json:"checkpoint_id"`
	CreatedAt    string    `json:"created_at"`
	ExpiredAt    string    `json:"expired_at"`
	ImageType    ImageType `json:"image_type"`
	ParentID     string    `json:"parent_id"`
	ProjectID    string    `json:"project_id"`
	ProtectedAt  string    `json:"protected_at"`
	ResourceAZ   string    `json:"resource_az"`
	ResourceID   string    `json:"resource_id"`
	ResourceName string    `json:"resource_name"`
	ResourceSize int       `json:"resource_size"`
	ResourceType string    `json:"resource_type"`
	UpdatedAt    string    `json:"updated_at"`
	VaultID      string    `json:"vault_id"`

	EnterpriseProjectID string `json:"enterprise_project_id"`

	Status BackupStatus `json:"backup_status"`
}

type Backup struct {
	BackupResp
	Children []BackupResp `json:"children"`
}

func (r GetResult) Extract() (*Backup, error) {
	var s struct {
		Backup *Backup `json:"backup"`
	}
	if r.Err != nil {
		return nil, r.Err
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, fmt.Errorf("error extracting backup details from response: %s", err)
	}
	return s.Backup, err
}

type BackupPage struct {
	pagination.MarkerPageBase
}

// ExtractInto is a function that takes a ListResult and returns the
// backups' information.
func ExtractInto(r pagination.Page) ([]BackupResp, error) {
	var s []BackupResp
	err := (r.(BackupPage)).ExtractInto(&s)
	return s, err
}
