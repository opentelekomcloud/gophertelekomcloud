package tags

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"

type AddImageTagOpts struct {
	// Specifies the image ID.
	ImageId string `json:"-" required:"true"`
	// Specifies the tag to be added or updated.
	Tag tags.ResourceTag `json:"tag" required:"true"`
}

// POST /v2/{project_id}/images/{image_id}/tags

// 204
