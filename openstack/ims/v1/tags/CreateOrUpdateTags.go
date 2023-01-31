package tags

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"

type CreateOrUpdateTagsOpts struct {
	// Specifies the image ID.
	ImageId string `json:"image_id" required:"true"`
	// Specifies the tag.
	//
	// Use either tag or image_tag.
	Tag string `json:"tag,omitempty"`
	// Lists the image tags. For detailed description, see Image Tag Data Formats. This parameter is left blank by default.
	//
	// Use either tag or image_tag.
	ImageTag tags.ResourceTag `json:"image_tag,omitempty"`
}

// PUT /v1/cloudimages/tags

// 204
