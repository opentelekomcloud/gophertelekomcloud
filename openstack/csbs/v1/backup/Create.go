package backup

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOpts contains the options for create a Backup. This object is passed to backup.Create().
type CreateOpts struct {
	BackupName   string             `json:"backup_name,omitempty"`
	Description  string             `json:"description,omitempty"`
	ResourceType string             `json:"resource_type,omitempty"`
	Incremental  *bool              `json:"incremental,omitempty"`
	Tags         []tags.ResourceTag `json:"tags,omitempty"`
	ExtraInfo    interface{}        `json:"extra_info,omitempty"`
}

// Create will create a new backup based on the values in CreateOpts. To extract
// the checkpoint object from the response, call the Extract method on the CreateResult.
func Create(client *golangsdk.ServiceClient, resourceID string, opts CreateOpts) (*Checkpoint, error) {
	b, err := golangsdk.BuildRequestBody(opts, "protect")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("providers", providerID, "resources", resourceID, "action"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	if err != nil {
		return nil, err
	}

	var res Checkpoint
	err = extract.IntoStructPtr(raw.Body, &res, "checkpoint")
	return &res, err
}

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
