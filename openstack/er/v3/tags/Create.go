package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type TagOpts struct {
	// Resource tag
	Tag tags.ResourceTag `json:"tag" required:"true"`
}

// Create creates tag
func Create(client *golangsdk.ServiceClient, resourceType, resourceId string, opts TagOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	_, err = client.Post(client.ServiceURL(resourceType, resourceId, "tags"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
