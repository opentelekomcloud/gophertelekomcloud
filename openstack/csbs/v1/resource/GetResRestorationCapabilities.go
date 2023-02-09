package resource

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type GetRestorationOpts struct {
	// ID of the backup used to restore data
	CheckpointItemId string `json:"checkpoint_item_id" required:"true"`
	// Restoration target
	Target RestorableTarget `json:"target" required:"true"`
}

type RestorableTarget struct {
	// ID of the resource to which the backup is restored
	ResourceId string `json:"resource_id" required:"true"`
	// Type of the target to which the backup is restored, for example, OS::Nova::Server for an ECS
	ResourceType string `json:"resource_type" required:"true"`
	// Disk mapping list for restoring an ECS. Enter the mapping between disks and backups based on the actual situation.
	Volumes []RestoreVolumeMapping `json:"volumes" required:"true"`
}

type RestoreVolumeMapping struct {
	// Disk backup ID. Use the API in 2.3.3 Querying a Single Backup to obtain the disk backup ID.
	BackupId string `json:"backup_id" required:"true"`
	// ID of the destination EVS disk for the restoration
	VolumeId string `json:"volume_id" required:"true"`
}

// GetResRestorationCapabilities check whether target resources can be restored.
func GetResRestorationCapabilities(client *golangsdk.ServiceClient, opts []GetRestorationOpts) ([]ResourceCapability, error) {
	return doAction(client, opts, "check_restorable", "restorable")
}
