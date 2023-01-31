package tags

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"

// GET /v2/{project_id}/images/tags

type ListImagesTagsResponse struct {
	Tags []tags.ListedTag `json:"tags,omitempty"`
}
