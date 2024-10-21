package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

func List(client *golangsdk.ServiceClient, resourceType, resourceId string) ([]tags.ResourceTag, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(resourceType, resourceId, "tags").Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []tags.ResourceTag
	err = extract.IntoSlicePtr(raw.Body, &res, "tags")
	return res, err
}
