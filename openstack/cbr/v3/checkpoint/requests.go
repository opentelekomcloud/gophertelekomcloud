package checkpoint

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// ID of the vault
	VaultID string `json:"vault_id" required:"true"`
	// Checkpoint parameters
	Parameters CheckpointParam `json:"parameters,omitempty"`
}

type CheckpointParam struct {
	// Describes whether automatic triggering is enabled
	// Default: false
	AutoTrigger bool `json:"auto_trigger,omitempty"`
	// Backup description
	Description string `json:"description,omitempty"`
	// Whether bacup is incremental or not
	// Default: true
	Incremental bool `json:"incremental,omitempty"`
	// Backup name
	Name string `json:"name,omitempty"`
	// UUID list of resources to be backed up
	Resources []string `json:"resources,omitempty"`
	// Additional information on Resource
	ResourceDetails []Resource `json:"resource_details,omitempty"`
}

type Resource struct {
	// ID of the resource to be backed up
	ID string `json:"id"`
	// Name of the resource to be backed up
	Name string `json:"name,omitempty"`
	// Type of the resource to be backed up
	// OS::Nova::Server | OS::Cinder::Volume
	Type string `json:"type,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Checkpoint, error) {
	b, err := build.RequestBodyMap(opts, "checkpoint")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("checkpoints"), b, nil, &golangsdk.RequestOpts{
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
	CreatedAt string    `json:"created_at"`
	ID        string    `json:"id"`
	ProjectID string    `json:"project_id"`
	Status    string    `json:"status"`
	Vault     Vault     `json:"vault"`
	ExtraInfo ExtraInfo `json:"extra_info"`
}

type Vault struct {
	ID               string                `json:"id"`
	Name             string                `json:"name"`
	Resources        []CheckpointResources `json:"resources"`
	SkippedResources []SkippedResources    `json:"skipped_resources"`
}

type ExtraInfo struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	RetentionDuration int    `json:"retention_duration"`
}

type CheckpointResources struct {
	ExtraInfo     string `json:"extra_info"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	ProtectStatus string `json:"protect_status"`
	ResourceSize  string `json:"resource_size"`
	Type          string `json:"type"`
	BackupSize    string `json:"backup_size"`
	BackupCount   string `json:"backup_count"`
}

type SkippedResources struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Reason string `json:"reason"`
}
