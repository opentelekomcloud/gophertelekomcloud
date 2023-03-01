package members

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/members"
)

// UpdateOpts represents options to an Update request.
type UpdateOpts struct {
	MemberOpts
	// Specifies whether a shared image will be accepted or declined.
	//
	// Available values include:
	//
	// accepted: indicates that a shared image is accepted. After an image is accepted, the image is displayed in the image list. You can use the image to create ECSs.
	//
	// rejected: indicates that a shared image is declined. After an image is rejected, the image is not displayed in the image list. However, you can still use the image to create ECSs.
	Status string `json:"status" required:"true"`
	// Specifies the ID of a vault.
	//
	// This parameter is mandatory if you want to accept a shared full-ECS image created from a CBR backup.
	//
	// You can obtain the vault ID from the CBR console or section "Querying the Vault List" in Cloud Backup and Recovery API Reference.
	VaultID string `json:"vault_id,omitempty"`
}

// Update function updates member.
func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*members.Member, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/images/{image_id}/members/{member_id}
	raw, err := client.Put(client.ServiceURL("images", opts.ImageId, "members", opts.MemberId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
