package backups

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Backup struct {
	CheckpointID string            `json:"checkpoint_id"`
	CreatedAt    string            `json:"created_at"`
	Description  string            `json:"description"`
	ExpiredAt    string            `json:"expired_at"`
	ExtendInfo   *BackupExtendInfo `json:"extend_info"`
	ID           string            `json:"id"`
	ImageType    string            `json:"image_type"`
	Name         string            `json:"name"`
	ParentID     string            `json:"parent_id"`
	ProjectID    string            `json:"project_id"`
	ProtectedAt  string            `json:"protected_at"`
	ResourceAZ   string            `json:"resource_az"`
	ResourceID   string            `json:"resource_id"`
	ResourceName string            `json:"resource_name"`
	ResourceSize int               `json:"resource_size"`
	ResourceType string            `json:"resource_type"`
	Status       string            `json:"status"`
	UpdatedAt    string            `json:"updated_at"`
	VaultId      string            `json:"vault_id"`
	ProviderID   string            `json:"provider_id"`
	Children     []BackupResp      `json:"children"`
}

type BackupExtendInfo struct {
	AutoTrigger          bool        `json:"auto_trigger"`
	Bootable             bool        `json:"bootable"`
	Incremental          bool        `json:"incremental"`
	SnapshotID           string      `json:"snapshot_id"`
	SupportLld           bool        `json:"support_lld"`
	SupportedRestoreMode string      `json:"supported_restore_mode"`
	OsImagesData         []ImageData `json:"os_images_data"`
	ContainSystemDisk    bool        `json:"contain_system_disk"`
	Encrypted            bool        `json:"encrypted"`
	SystemDisk           bool        `json:"system_disk"`
}

type BackupResp struct {
	CheckpointID string            `json:"checkpoint_id"`
	CreatedAt    string            `json:"created_at"`
	Description  string            `json:"description"`
	ExpiredAt    string            `json:"expired_at"`
	ExtendInfo   *BackupExtendInfo `json:"extend_info"`
	ID           string            `json:"id"`
	ImageType    string            `json:"image_type"`
	Name         string            `json:"name"`
	ParentID     string            `json:"parent_id"`
	ProjectID    string            `json:"project_id"`
	ProtectedAt  string            `json:"protected_at"`
	ResourceAZ   string            `json:"resource_az"`
	ResourceID   string            `json:"resource_id"`
	ResourceName string            `json:"resource_name"`
	ResourceSize int               `json:"resource_size"`
	ResourceType string            `json:"resource_type"`
	Status       string            `json:"status"`
	UpdatedAt    string            `json:"updated_at"`
	VaultID      string            `json:"vault_id"`
	ProviderID   string            `json:"provider_id"`
}

type ImageData struct {
	ImageID string `json:"image_id"`
}

type BackupPage struct {
	pagination.LinkedPageBase
}

func (r BackupPage) NextPageURL() (string, error) {
	var res []golangsdk.Link
	err := extract.IntoSlicePtr(bytes.NewReader(r.Result.Body), &res, "backups_links")
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res)
}

func (r BackupPage) IsEmpty() (bool, error) {
	is, err := ExtractBackups(r)
	return len(is) == 0, err
}

func ExtractBackups(r pagination.Page) ([]Backup, error) {
	var res []Backup
	err := extract.IntoSlicePtr(bytes.NewReader(r.(BackupPage).Body), &res, "backups")
	if err != nil {
		return nil, err
	}
	return res, nil
}
