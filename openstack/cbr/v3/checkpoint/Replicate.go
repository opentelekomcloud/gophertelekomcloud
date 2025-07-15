package checkpoint

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ReplicateOpts struct {
	// Parameters in the request body of performing a replication
	Replicate ReplicateParam `json:"replicate" required:"true"`
}

type ReplicateParam struct {
	// Parameters in the request body of performing a replication
	AutoTrigger bool `json:"auto_trigger"`
	// ID of the replication destination project
	DestinationProjectId string `json:"destination_project_id" required:"true"`
	// ID of the replication destination region
	DestinationRegion string `json:"destination_region" required:"true"`
	// ID of the vault in the replication destination region
	DestinationVaultId string `json:"destination_vault_id" required:"true"`
	// Vault ID
	VaultId string `json:"vault_id" required:"true"`
}

func Replicate(client *golangsdk.ServiceClient, opts ReplicateOpts) (*Replication, error) {
	b, err := build.RequestBody(opts, "checkpoint")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/checkpoints/replicate
	raw, err := client.Post(client.ServiceURL("checkpoints", "replicate"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Replication
	err = extract.IntoStructPtr(raw.Body, &res, "replication")
	return &res, err
}

type Replication struct {
	VaultId              string             `json:"vault_id"`
	DestinationProjectId string             `json:"destination_project_id"`
	DestinationRegion    string             `json:"destination_region"`
	DestinationVaultId   string             `json:"destination_vault_id"`
	ProjectId            string             `json:"project_id"`
	ProviderId           string             `json:"provider_id"`
	SourceRegion         string             `json:"source_region"`
	Backups              []ReplicateBackups `json:"backups"`
}

type ReplicateBackups struct {
	BackupId            string `json:"backup_id"`
	ReplicationRecordId string `json:"replication_record_id"`
}
