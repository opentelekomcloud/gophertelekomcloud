package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type TagOpts struct {
	// Method of adding tags.
	Action string `json:"action" required:"true"`
	// Whether to call external APIs.
	IsOpen bool `json:"is_open" required:"true"`
	// Resource tags
	Tags []tags.ResourceTag `json:"tags" required:"true"`
}

func Manage(client *golangsdk.ServiceClient, resourceType, resourceId string, opts TagOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	// POST /v1/{project_id}/{resource_type}/{resource_id}/tags/action
	_, err = client.Post(client.ServiceURL(resourceType, resourceId, "tags", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return err
}
