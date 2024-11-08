package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/instances"
)

func GetBackupPolicy(client *golangsdk.ServiceClient, instanceId string) (*instances.BackupStrategy, error) {
	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/backups/policy
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "backups", "policy"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res instances.BackupStrategy
	err = extract.IntoStructPtr(raw.Body, &res, "backup_policy")
	return &res, err
}
