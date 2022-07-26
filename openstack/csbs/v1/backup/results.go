package backup

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Checkpoint struct {
	Status         string         `json:"status"`
	CreatedAt      string         `json:"created_at"`
	Id             string         `json:"id"`
	ResourceGraph  string         `json:"resource_graph"`
	ProjectId      string         `json:"project_id"`
	ProtectionPlan ProtectionPlan `json:"protection_plan"`
	ExtraInfo      interface{}    `json:"extra_info"`
}

type ProtectionPlan struct {
	Id              string               `json:"id"`
	Name            string               `json:"name"`
	BackupResources []СsbsBackupResource `json:"resources"`
}

type СsbsBackupResource struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Name      string      `json:"name"`
	ExtraInfo interface{} `json:"extra_info"`
}

type ResourceCapability struct {
	Result       bool   `json:"result"`
	ResourceType string `json:"resource_type"`
	ErrorCode    string `json:"error_code"`
	ErrorMsg     string `json:"error_msg"`
	ResourceId   string `json:"resource_id"`
}

type Backup struct {
	CheckpointId string             `json:"checkpoint_id"`
	CreatedAt    string             `json:"created_at"`
	ExtendInfo   ExtendInfo         `json:"extend_info"`
	Id           string             `json:"id"`
	Name         string             `json:"name"`
	ResourceId   string             `json:"resource_id"`
	Status       string             `json:"status"`
	UpdatedAt    string             `json:"updated_at"`
	VMMetadata   VMMetadata         `json:"backup_data"`
	Description  string             `json:"description"`
	Tags         []tags.ResourceTag `json:"tags"`
	ResourceType string             `json:"resource_type"`
}

type ExtendInfo struct {
	AutoTrigger          bool           `json:"auto_trigger"`
	AverageSpeed         float32        `json:"average_speed"`
	CopyFrom             string         `json:"copy_from"`
	CopyStatus           string         `json:"copy_status"`
	FailCode             FailCode       `json:"fail_code"`
	FailOp               string         `json:"fail_op"`
	FailReason           string         `json:"fail_reason"`
	ImageType            string         `json:"image_type"`
	Incremental          bool           `json:"incremental"`
	Progress             int            `json:"progress"`
	ResourceAz           string         `json:"resource_az"`
	ResourceName         string         `json:"resource_name"`
	ResourceType         string         `json:"resource_type"`
	Size                 int            `json:"size"`
	SpaceSavingRatio     float32        `json:"space_saving_ratio"`
	VolumeBackups        []VolumeBackup `json:"volume_backups"`
	FinishedAt           string         `json:"finished_at"`
	TaskId               string         `json:"taskid"`
	HypervisorType       string         `json:"hypervisor_type"`
	SupportedRestoreMode string         `json:"supported_restore_mode"`
	Supportlld           bool           `json:"support_lld"`
}

type VMMetadata struct {
	RegionName       string `json:"__openstack_region_name"`
	CloudServiceType string `json:"cloudservicetype"`
	Disk             int    `json:"disk"`
	ImageType        string `json:"imagetype"`
	Ram              int    `json:"ram"`
	Vcpus            int    `json:"vcpus"`
	Eip              string `json:"eip"`
	PrivateIp        string `json:"private_ip"`
}

type FailCode struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
}

type VolumeBackup struct {
	AverageSpeed     int    `json:"average_speed"`
	Bootable         bool   `json:"bootable"`
	Id               string `json:"id"`
	ImageType        string `json:"image_type"`
	Incremental      bool   `json:"incremental"`
	SnapshotID       string `json:"snapshot_id"`
	Name             string `json:"name"`
	Size             int    `json:"size"`
	SourceVolumeId   string `json:"source_volume_id"`
	SourceVolumeSize int    `json:"source_volume_size"`
	SpaceSavingRatio int    `json:"space_saving_ratio"`
	Status           string `json:"status"`
	SourceVolumeName string `json:"source_volume_name"`
}

func (r QueryResult) ExtractQueryResponse() ([]ResourceCapability, error) {
	var s []ResourceCapability
	err := r.ExtractIntoSlicePtr(&s, "protectable")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Extract will get the checkpoint object from the golangsdk.Result
func (r CreateResult) Extract() (*Checkpoint, error) {
	s := new(Checkpoint)
	err := r.ExtractIntoStructPtr(s, "checkpoint")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Extract will get the backup object from the golangsdk.Result
func (r GetResult) Extract() (*Backup, error) {
	s := new(Backup)
	err := r.ExtractIntoStructPtr(s, "checkpoint_item")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// СsbsBackupPage is the page returned by a pager when traversing over a
// collection of backups.
type СsbsBackupPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of backups has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r СsbsBackupPage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:"checkpoint_items_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a СsbsBackupPage struct is empty.
func (r СsbsBackupPage) IsEmpty() (bool, error) {
	is, err := ExtractBackups(r)
	return len(is) == 0, err
}

// ExtractBackups accepts a Page struct, specifically a СsbsBackupPage struct,
// and extracts the elements into a slice of Backup structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractBackups(r pagination.Page) ([]Backup, error) {
	var s []Backup
	err := (r.(СsbsBackupPage)).ExtractIntoSlicePtr(&s, "checkpoint_items")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type CreateResult struct {
	golangsdk.Result
}

type DeleteResult struct {
	golangsdk.ErrResult
}

type GetResult struct {
	golangsdk.Result
}

type QueryResult struct {
	golangsdk.Result
}
