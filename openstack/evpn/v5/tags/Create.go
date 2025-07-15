package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type TagsOpts struct {
	// Specifies a tag list.
	// A maximum of 20 tags can be specified.
	Tags []tags.ResourceTag `json:"tags" required:"true"`
}

// Create creates tags
// resourceType Specifies the resource type.
// The value can be vpn-gateway, customer-gateway, or vpn-connection.
func Create(client *golangsdk.ServiceClient, resourceType, resourceId string, opts TagsOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	_, err = client.Post(client.ServiceURL(resourceType, resourceId, "tags", "create"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
	})
	return err
}
