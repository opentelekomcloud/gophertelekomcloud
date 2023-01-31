package tags

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"

// ImageId

// GET /v2/{project_id}/images/{image_id}/tags

type ListImageTagsResponse struct {
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}
