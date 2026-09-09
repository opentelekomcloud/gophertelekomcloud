package images

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateWholeImageFromCBRorCSBSOpts Parameters in the request body when a CSBS backup or CBR backup is used to create a full-ECS image
type CreateWholeImageFromCBRorCSBSOpts struct {
	// Specifies the image name. For detailed description, see Image Attributes.
	Name string `json:"name" required:"true"`
	// Provides supplementary information about the image. For detailed description, see Image Attributes.
	Description string `json:"description,omitempty"`
	// Lists the image tags. The value is left blank by default.
	//
	// Use either tags or image_tags.
	Tags []string `json:"tags,omitempty"`
	// Lists the image tags. The value is left blank by default.
	//
	// Use either tags or image_tags.
	ImageTags []tags.ResourceTag `json:"image_tags,omitempty"`
	// Specifies the CSBS backup ID or CBR backup ID.
	//
	// To obtain the CSBS backup ID, perform the following operations:
	//
	// Log in to the management console.
	//
	// Under Storage, click Cloud Server Backup Service.
	//
	// In the backup list, expand details of the backup to obtain its ID.
	//
	// To obtain the CBR backup ID, perform the following operations:
	//
	// Log in to the management console.
	//
	// Under Storage, click Cloud Backup and Recovery.
	//
	// On the displayed Cloud Server Backup page, click the Backups tab and obtain the backup ID from the backup list.
	BackupId string `json:"backup_id" required:"true"`
	// Specifies the maximum memory of the image in the unit of MB. This parameter is not configured by default.
	MaxRam int `json:"max_ram,omitempty"`
	// Specifies the minimum memory of the image in the unit of MB. The default value is 0, indicating that the memory is not restricted.
	MinRam int `json:"min_ram,omitempty"`
	// Specifies the method of creating a full-ECS image.
	//
	// If the value is CBR, a CBR backup is used to create a full-ECS image. In this case, backup_id is the CBR backup ID.
	//
	// If the value is CSBS, a CSBS backup is used to create a full-ECS image. In this case, backup_id is the CSBS backup ID.
	//
	// If you do not specify this parameter, value CSBS is used by default.
	WholeImageType string `json:"whole_image_type,omitempty"`
	// Specifies the enterprise project that an image belongs to.
	//
	// If only enterprise project authorization is used, the enterprise_project_id parameter must be specified. Otherwise, an error may occur, indicating that you do not have the required permissions.
	//
	// If the value is 0 or left blank, an image belongs to the default enterprise project.
	// If the value is a UUID, an image belongs to the enterprise project with this UUID.
	// For more information about enterprise projects and how to obtain enterprise project IDs, see Enterprise Project Service User Guide.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

func CreateWholeImageFromCBRorCSBS(client *golangsdk.ServiceClient, opts CreateWholeImageFromCBRorCSBSOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return wholeImages(client, b)
}
