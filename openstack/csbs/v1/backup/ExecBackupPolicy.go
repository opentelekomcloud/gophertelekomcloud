package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CheckpointOpts struct {
	// Backup policy ID. Refer to the backup policy ID that is returned by the API of 2.2.5 Querying the Backup Policy List.
	PlanId string `json:"plan_id"`
	// Backup parameters
	Parameters CheckpointParam `json:"parameters"`
}

type CheckpointParam struct {
	// Whether automatic trigger is enabled
	AutoTrigger bool `json:"auto_trigger"`
	// ID list of resources to be backed up
	Resources []string `json:"resources"`
}

func ExecBackupPolicy(client *golangsdk.ServiceClient, opts CheckpointOpts) (*Checkpoint, error) {
	b, err := build.RequestBody(opts, "checkpoint")
	if err != nil {
		return nil, err
	}

	// POST https://{endpoint}/v1/{project_id}/providers/{provider_id}/checkpoints
	raw, err := client.Post(client.ServiceURL("providers", ProviderID, "checkpoints"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Checkpoint
	err = extract.IntoStructPtr(raw.Body, &res, "checkpoint")
	return &res, err
}
