package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

func ListImagesTags(client *golangsdk.ServiceClient) ([]tags.ListedTag, error) {
	// GET /v2/{project_id}/images/tags
	raw, err := client.Get(client.ServiceURL("images", "tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []tags.ListedTag
	err = extract.IntoSlicePtr(raw.Body, &res, "tags")
	return res, err
}
