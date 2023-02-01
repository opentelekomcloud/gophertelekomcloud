package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

func ListImageTags(client *golangsdk.ServiceClient, imageId string) ([]tags.ResourceTag, error) {
	// GET /v2/{project_id}/images/{image_id}/tags
	raw, err := client.Get(client.ServiceURL("images", imageId, "tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []tags.ResourceTag
	err = extract.IntoSlicePtr(raw.Body, &res, "tags")
	return res, err
}
