package resource

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type RestoreOpts struct {
	// Backup provider ID, which specifies whether the backup object is a server or disk.
	// This parameter has a fixed value. For CSBS, the value is fc4d5750-22e7-4798-8a46-f48f62c4c1da.
	ProviderId string `json:"provider_id"`
	// Backup record ID
	CheckpointId string `json:"checkpoint_id"`
	// Restoration target
	RestoreTarget string `json:"restore_target,omitempty"`
	// Restoration parameters
	Parameters RestoreParam `json:"parameters"`
}

type RestoreParam struct {
	// Backup ID
	CheckpointItemId string `json:"checkpoint_item_id"`
	// Whether to instantly power on the VM after restoration
	PowerOn bool `json:"power_on"`
	// Restoration target
	Targets RestoreTarget `json:"targets"`
}

type RestoreTarget struct {
	// ID of the ECS to be restored
	ServerId string `json:"server_id"`
	// List of the mappings between disk backups and target disks
	Volumes []RestoreVolumeMapping `json:"volumes"`
}

func CreateRestorationTask(client *golangsdk.ServiceClient, opts RestoreOpts) (*RestoreResp, error) {
	b, err := build.RequestBody(opts, "restore")
	if err != nil {
		return nil, err
	}

	// POST https://{endpoint}/v1/{project_id}/restores
	raw, err := client.Post(client.ServiceURL("restores"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res RestoreResp
	err = extract.IntoStructPtr(raw.Body, &res, "restore")
	return &res, err
}

type RestoreResp struct {
	// Restoration target
	RestoreTarget string `json:"restore_target"`
	// Status
	Status string `json:"status"`
	// Project ID
	ProviderId string `json:"provider_id"`
	// Resource status after the resource is restored, for example, available
	ResourcesStatus interface{} `json:"resources_status"`
	// Restoration parameters
	Parameters RestoreParam `json:"parameters"`
	// Backup record ID
	CheckpointId string `json:"checkpoint_id"`
	// Project ID
	ProjectId string `json:"project_id"`
	// Restoration ID
	Id string `json:"id"`
	// Cause of the resource restoration failure
	ResourcesReason interface{} `json:"resources_reason"`
}
