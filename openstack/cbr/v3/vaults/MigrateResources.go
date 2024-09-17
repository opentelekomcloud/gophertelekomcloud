package vaults

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type MigrateOpts struct {
	DestinationVaultId string   `json:"destination_vault_id" required:"true"`
	ResourceIds        []string `json:"resource_ids" required:"true"`
}

func MigrateResources(client *golangsdk.ServiceClient, vaultID string, opts MigrateOpts) ([]string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/vaults/{vault_id}/migrateresources
	raw, err := client.Post(client.ServiceURL("vaults", vaultID, "migrateresources"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []string
	return res, extract.IntoSlicePtr(raw.Body, &res, "migrated_resources")
}
