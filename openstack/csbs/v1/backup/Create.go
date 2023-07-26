package backup

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOpts contains the options for create a Backup. This object is passed to backup.Create().
type CreateOpts struct {
	// Backup name. The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).
	BackupName string `json:"backup_name,omitempty"`
	// Backup description. The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).
	Description string `json:"description,omitempty"`
	// Entity object type of the backup object
	// The current value is OS::Nova::Server indicating that the backup object is an ECS.
	// If this parameter is not passed, the backup object type defaults to OS::Nova::Server.
	ResourceType string `json:"resource_type,omitempty"`
	// Backup type. Value True indicates incremental backup and value False indicates full backup.
	// For the initial backup, full backup is always adopted, in spite of which value is set.
	Incremental *bool `json:"incremental,omitempty"`
	// Tag list
	// This list cannot be an empty list.
	// The list can contain up to 10 keys.
	// Keys in this list must be unique.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
	// Additional information about the backup object
	ExtraInfo any `json:"extra_info,omitempty"`
}

// Create will create a new backup based on the values in CreateOpts. To extract
// the checkpoint object from the response, call the Extract method on the CreateResult.
func Create(client *golangsdk.ServiceClient, resourceID string, opts CreateOpts) (*Checkpoint, error) {
	b, err := build.RequestBody(opts, "protect")
	if err != nil {
		return nil, err
	}

	// POST https://{endpoint}/v1/{project_id}/providers/{provider_id}/resources/{resource_id}/action
	raw, err := client.Post(client.ServiceURL("providers", ProviderID, "resources", resourceID, "action"), b,
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
	// Backup status
	// Enum:[ waiting_protect, protecting, available, waiting_restore, restoring, error, waiting_delete, deleting,deleted]
	Status string `json:"status"`
	// Creation time, for example, 2017-04-18T01:21:52.701973
	CreatedAt string `json:"created_at"`
	// Backup record ID
	Id string `json:"id"`
	// Resource diagram, which displays the inclusion relationship between backups and sub-backups
	ResourceGraph string `json:"resource_graph"`
	// Project ID
	ProjectId string `json:"project_id"`
	// Backup plan information
	ProtectionPlan ProtectionPlan `json:"protection_plan"`
	// Additional information
	ExtraInfo any `json:"extra_info"`
}

type ProtectionPlan struct {
	// Backup policy ID
	Id string `json:"id"`
	// Backup policy name
	Name string `json:"name"`
	// Backup object list
	// For details, see Table 2-8.
	BackupResources []СsbsBackupResource `json:"resources"`
}

type СsbsBackupResource struct {
	// ID of the object to be backed up
	ID string `json:"id"`
	// Entity object type of the backup object. The value is fixed at OS::Nova::Server, indicating that the object type is ECSs.
	Type string `json:"type"`
	// Backup object name
	Name string `json:"name"`
	// Additional information about the backup object
	ExtraInfo any `json:"extra_info"`
}
