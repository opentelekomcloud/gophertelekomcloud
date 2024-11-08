package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ReplicateBackupOpts struct {
	// Replica description
	Description string `json:"description,omitempty"`
	// ID of the replication destination project
	DestinationProjectId string `json:"destination_project_id" required:"true"`
	// Replication destination region
	DestinationRegion string `json:"destination_region" required:"true"`
	// ID of the vault in the replication destination region
	DestinationVaultId string `json:"destination_vault_id" required:"true"`
	// Replica name
	Name string `json:"name,omitempty"`
}

func ReplicateBackup(client *golangsdk.ServiceClient, backupID string, opts ReplicateBackupOpts) (*Replicate, error) {
	b, err := build.RequestBody(opts, "replicate")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("backups", backupID, "replicate"), b, nil, nil)
	if err != nil {
		return nil, err
	}
	var res Replicate
	err = extract.IntoStructPtr(raw.Body, &res, "replication")
	return &res, err
}

type Replicate struct {
	// ID of the source backup used for replication
	BackupId string `json:"backup_id"`
	// ID of the replication destination project
	DestinationProjectId string `json:"destination_project_id"`
	// Replication destination region
	DestinationRegion string `json:"destination_region"`
	// ID of the vault in the replication destination region
	DestinationVaultId string `json:"destination_vault_id"`
	// ID of the project where replication is performed
	ProjectId string `json:"project_id"`
	// Resource type ID
	ProviderId string `json:"provider_id"`
	// Replication record ID
	ReplicationRecordId string `json:"replication_record_id"`
	// Replication source region
	SourceRegion string `json:"source_region"`
}
