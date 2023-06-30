package images

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/others"
)

// CreateWholeImageFromECSOpts Parameters for creating a full-ECS image using an ECS
type CreateWholeImageFromECSOpts struct {
	// Specifies the image name. For detailed description, see Image Attributes.
	Name string `json:"name" required:"true"`
	// Provides supplementary information about the image. For detailed description, see Image Attributes.
	Description string `json:"description,omitempty"`
	// Lists the image tags. The value is left blank by default.
	// Use either tags or image_tags.
	Tags []string `json:"tags,omitempty"`
	// Lists the image tags. The value is left blank by default.
	// Use either tags or image_tags.
	ImageTags []tags.ResourceTag `json:"image_tags,omitempty"`
	// Specifies the ECS ID. This parameter is required when an ECS is used to create a full-ECS image.
	//
	// To obtain the ECS ID, perform the following operations:
	//
	// Log in to management console.
	//
	// Under Computing, click Elastic Cloud Server.
	//
	// In the ECS list, click the name of the ECS and view its ID.
	InstanceId string `json:"instance_id" required:"true"`
	// Specifies the maximum memory of the image in the unit of MB. This parameter is not configured by default.
	MaxRam int `json:"max_ram,omitempty"`
	// Specifies the minimum memory of the image in the unit of MB. The default value is 0.
	MinRam int `json:"min_ram,omitempty"`
	// Specifies the ID of the vault to which an ECS is to be added or has been added.
	//
	// To create a full-ECS image from an ECS, create a backup from the ECS and then use the backup to create a full-ECS image. If a CBR backup is created, vault_id is mandatory. If a CSBS backup is created, vault_id is optional.
	//
	// You can obtain the vault ID from the CBR console or section "Querying the Vault List" in Cloud Backup and Recovery API Reference.
	VaultId string `json:"vault_id,omitempty"`
}

func CreateWholeImageFromECS(client *golangsdk.ServiceClient, opts CreateWholeImageFromECSOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return wholeImages(client, b)
}

func wholeImages(client *golangsdk.ServiceClient, b *build.Body) (*string, error) {
	// POST /v1/cloudimages/wholeimages/action
	raw, err := client.Post(client.ServiceURL("cloudimages", "wholeimages", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return others.ExtractJobId(err, raw)
}
