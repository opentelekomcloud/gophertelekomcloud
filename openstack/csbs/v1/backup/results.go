package backup

import (
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

const providerID = "fc4d5750-22e7-4798-8a46-f48f62c4c1da"

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
