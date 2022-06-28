package backups

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type backupResult struct {
	golangsdk.Result
}

type GetResult struct {
	backupResult
}

type Backup struct {
	CheckpointId string            `json:"checkpoint_id"`
	CreatedAt    string            `json:"created_at"`
	Description  string            `json:"description"`
	ExpiredAt    string            `json:"expired_at"`
	ExtendInfo   *BackupExtendInfo `json:"extend_info"`
	ID           string            `json:"id"`
	ImageType    string            `json:"image_type"`
	Name         string            `json:"name"`
	ParentId     string            `json:"parent_id"`
	ProjectId    string            `json:"project_id"`
	ProtectedAt  string            `json:"protected_at"`
	ResourceAZ   string            `json:"resource_az"`
	ResourceID   string            `json:"resource_id"`
	ResourceName string            `json:"resource_name"`
	ResourceSize int               `json:"resource_size"`
	ResourceType string            `json:"resource_type"`
	Status       string            `json:"status"`
	UpdatedAt    string            `json:"updated_at"`
	VaultId      string            `json:"vault_id"`
	ProviderId   string            `json:"provider_id"`
	Children     *[]BackupResp     `json:"children"`
}
type BackupExtendInfo struct {
	AutoTrigger          bool        `json:"auto_trigger"`
	Bootable             bool        `json:"bootable"`
	Incremental          bool        `json:"incremental"`
	SnapshotId           string      `json:"snapshot_id"`
	SupportLld           bool        `json:"support_lld"`
	SupportedRestoreMode string      `json:"supported_restore_mode"`
	OsImagesData         []ImageData `json:"os_images_data"`
	ContainSystemDisk    bool        `json:"contain_system_disk"`
	Encrypted            bool        `json:"encrypted"`
	SystemDisk           bool        `json:"system_disk"`
}

type BackupResp struct {
	CheckpointId string            `json:"checkpoint_id"`
	CreatedAt    string            `json:"created_at"`
	Description  string            `json:"description"`
	ExpiredAt    string            `json:"expired_at"`
	ExtendInfo   *BackupExtendInfo `json:"extend_info"`
	ID           string            `json:"id"`
	ImageType    string            `json:"image_type"`
	Name         string            `json:"name"`
	ParentId     string            `json:"parent_id"`
	ProjectId    string            `json:"project_id"`
	ProtectedAt  string            `json:"protected_at"`
	ResourceAZ   string            `json:"resource_az"`
	ResourceID   string            `json:"resource_id"`
	ResourceName string            `json:"resource_name"`
	ResourceSize int               `json:"resource_size"`
	ResourceType string            `json:"resource_type"`
	Status       string            `json:"status"`
	UpdatedAt    string            `json:"updated_at"`
	VaultId      string            `json:"vault_id"`
	ProviderId   string            `json:"provider_id"`
}

type ImageData struct {
	ImageId string `json:"image_id"`
}

func (r backupResult) Extract() (*Backup, error) {
	var s struct {
		Backup *Backup `json:"backup"`
	}
	if r.Err != nil {
		return nil, r.Err
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, fmt.Errorf("error extracting backup from get response: %s", err)
	}
	return s.Backup, err
}

type BackupPage struct {
	pagination.SinglePageBase
}

func ExtractBackups(r pagination.Page) ([]Backup, error) {
	var s []Backup
	err := r.(BackupPage).Result.ExtractIntoSlicePtr(&s, "backups")
	if err != nil {
		return nil, err
	}
	return s, nil
}
