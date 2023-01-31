package images

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"

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
}

// POST /v1/cloudimages/wholeimages/action

// 200 JobResponse
