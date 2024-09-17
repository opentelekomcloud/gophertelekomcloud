package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	BackupID string `json:"_"`
	// Status of a shared backup
	Status string `json:"status" required:"true"`
	// Vault in which the shared backup is to be stored
	VaultId string `json:"vault_id,omitempty"`
}

func UpdateSharingMember(client *golangsdk.ServiceClient, memberID string, opts UpdateOpts) (*Member, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("backups", opts.BackupID, "members", memberID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Member

	err = extract.IntoStructPtr(raw.Body, &res, "member")
	return &res, err
}
