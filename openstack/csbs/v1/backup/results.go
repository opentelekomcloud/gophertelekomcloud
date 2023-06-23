package backup

import (
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

const ProviderID = "fc4d5750-22e7-4798-8a46-f48f62c4c1da"

type Backup struct {
	// Backup record ID
	CheckpointId string `json:"checkpoint_id"`
	// Creation time, for example, 2017-04-18T01:21:52.701973
	CreatedAt string `json:"created_at"`
	// Extension information
	ExtendInfo ExtendInfo `json:"extend_info"`
	// Backup ID
	Id string `json:"id"`
	// Backup name
	Name string `json:"name"`
	// ID of the object to be backed up
	ResourceId string `json:"resource_id"`
	// Backup status. Possible values are waiting_protect, protecting, available, waiting_restore, restoring, error, waiting_delete, deleting, and deleted.
	// Enum:[ waiting_protect, protecting, available, waiting_restore, restoring, error, waiting_delete, deleting,deleted]
	Status string `json:"status"`
	// Modification time, for example, 2017-04-18T01:21:52.701973
	UpdatedAt string `json:"updated_at"`
	// VM metadata
	VMMetadata VMMetadata `json:"backup_data"`
	// Backup description
	Description string `json:"description"`
	// List of backup tags
	// Keys in the tag list must be unique.
	Tags []tags.ResourceTag `json:"tags"`
	// Type of the backup object
	ResourceType string `json:"resource_type"`
}

type ImagesData struct {
	ImageId string `json:"image_id"`
}

type ExtendInfo struct {
	// Whether automatic trigger is enabled
	AutoTrigger bool `json:"auto_trigger"`
	// Average rate. The unit is kb/s
	AverageSpeed float32 `json:"average_speed"`
	// The destination region of a backup replication. The default value is empty.
	CopyFrom string `json:"copy_from"`
	// Backup replication status. The default value is na.
	// Possible values are na, waiting_copy, copying, success, and fail.
	CopyStatus string `json:"copy_status"`
	// Error code
	FailCode FailCode `json:"fail_code"`
	// Type of the failed operation
	// Enum: [backup, restore, delete]
	FailOp string `json:"fail_op"`
	// Description of the failure cause
	FailReason string `json:"fail_reason"`
	// Backup type. For example, backup
	ImageType string `json:"image_type"`
	// Whether the backup is an enhanced backup
	Incremental bool `json:"incremental"`
	// Backup progress. The value is an integer ranging from 0 to 100.
	Progress int `json:"progress"`
	// AZ to which the backup resource belongs
	ResourceAz string `json:"resource_az"`
	// Backup object name
	ResourceName string `json:"resource_name"`
	// Type of the backup object. For example, OS::Nova::Server
	ResourceType string `json:"resource_type"`
	// Backup capacity. The unit is MB.
	Size int `json:"size"`
	// Space saving rate
	SpaceSavingRatio float32 `json:"space_saving_ratio"`
	// Volume backup list
	VolumeBackups []VolumeBackup `json:"volume_backups"`
	// Backup completion time, for example, 2017-04-18T01:21:52.701973
	FinishedAt string `json:"finished_at"`
	// Image data. This parameter has a value if an image has been created for the VM.
	OsImagesData []ImagesData `json:"os_images_data"`
	// Job ID
	TaskId string `json:"taskid"`
	// Virtualization type
	// The value is fixed at QEMU.
	HypervisorType string `json:"hypervisor_type"`
	// Restoration mode. Possible values are na, snapshot, and backup.
	// backup: Data is restored from backups of the EVS disks of the server.
	// na: Restoration is not supported.
	SupportedRestoreMode string `json:"supported_restore_mode"`
	// Whether to allow lazyloading for fast restoration
	Supportlld bool `json:"support_lld"`
}

type VMMetadata struct {
	// Name of the AZ where the server is located. If this parameter is left blank, such information about the server has not been obtained.
	RegionName string `json:"__openstack_region_name"`
	// Server type
	// The value is fixed at server (ECSs).
	CloudServiceType string `json:"cloudservicetype"`
	// System disk size corresponding to the server specifications
	Disk int `json:"disk"`
	// Image type
	// The value can be:
	// gold: public image
	// private: private image
	// market: market image
	ImageType string `json:"imagetype"`
	// Memory size of the server, in MB
	Ram int `json:"ram"`
	// CPU cores corresponding to the server
	Vcpus int `json:"vcpus"`
	// Elastic IP address of the server. If this parameter is left blank, such information about the server has not been obtained.
	Eip string `json:"eip"`
	// Internal IP address of the server. If this parameter is left blank, such information about the server has not been obtained.
	PrivateIp string `json:"private_ip"`
}

type FailCode struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
}

type VolumeBackup struct {
	// Average rate, in MB/s
	AverageSpeed int `json:"average_speed"`
	// Whether the disk is bootable
	// The value can be true or false.
	Bootable bool `json:"bootable"`
	// Cinder backup ID
	Id string `json:"id"`
	// Backup set type: backup
	// Enum:[ backup]
	ImageType string `json:"image_type"`
	// Whether incremental backup is used
	Incremental bool `json:"incremental"`
	// ID of the snapshot from which the backup is generated
	SnapshotID string `json:"snapshot_id"`
	// EVS disk backup name
	Name string `json:"name"`
	// Accumulated size (MB) of backups
	Size int `json:"size"`
	// Source disk ID
	SourceVolumeId string `json:"source_volume_id"`
	// Source volume size in GB
	SourceVolumeSize int `json:"source_volume_size"`
	// Space saving rate
	SpaceSavingRatio int `json:"space_saving_ratio"`
	// Status
	Status string `json:"status"`
	// Source volume name
	SourceVolumeName string `json:"source_volume_name"`
}
