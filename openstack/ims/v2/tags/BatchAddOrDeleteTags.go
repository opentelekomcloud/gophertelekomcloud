package tags

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"

type BatchAddOrDeleteTagsOpts struct {
	// Specifies the image ID.
	ImageId string `json:"-" required:"true"`
	// Specifies the tag operation to be performed. The value is case sensitive and can be create or delete. create indicates that tags will be added or updated, while delete indicates that tags will be deleted.
	Action string `json:"action" required:"true"`
	// Lists the tags to be added or deleted.
	Tags []tags.ResourceTag `json:"tags"`
}

// POST /v2/{project_id}/images/{image_id}/tags/action

// 204
