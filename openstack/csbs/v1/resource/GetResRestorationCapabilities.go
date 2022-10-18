package resource

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type GetRestorationOpts struct {
	CheckRestorable []CheckRestorable `json:"check_restorable"`
}

type CheckRestorable struct {
	// ID of the backup used to restore data
	CheckpointItemId string `json:"checkpoint_item_id"`
	// Restoration target
	Target RestorableTarget `json:"target"`
}

type RestorableTarget struct {
	// ID of the resource to which the backup is restored
	ResourceType string `json:"resource_type"`
	// Type of the target to which the backup is restored, for example, OS::Nova::Server for an ECS
	ResourceId string `json:"resource_id"`
	// Disk mapping list for restoring an ECS. Enter the mapping between disks and backups based on the actual situation.
	Volumes []RestoreVolumeMapping `json:"volumes"`
}

type RestoreVolumeMapping struct {
	// Disk backup ID. Use the API in 2.3.3 Querying a Single Backup to obtain the disk backup ID.
	BackupId string `json:"backup_id"`
	// ID of the destination EVS disk for the restoration
	VolumeId string `json:"volume_id"`
}

// GetResRestorationCapabilities check whether target resources can be restored.
func GetResRestorationCapabilities(client *golangsdk.ServiceClient, opts GetRestorationOpts) ([]ResourceCapability, error) {
	return doAction(client, opts, "restorable")
}
