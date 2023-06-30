package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

func doAction(client *golangsdk.ServiceClient, id string, opts tags.ActionOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("scaling_group_tag", id, "tags/action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return
}

func Create(client *golangsdk.ServiceClient, id string, tag []tags.ResourceTag) (err error) {
	opts := tags.ActionOpts{
		Tags:   tag,
		Action: "create",
	}
	return doAction(client, id, opts)
}

func Update(client *golangsdk.ServiceClient, id string, tag []tags.ResourceTag) (err error) {
	opts := tags.ActionOpts{
		Tags:   tag,
		Action: "update",
	}
	return doAction(client, id, opts)
}

func Delete(client *golangsdk.ServiceClient, id string, tag []tags.ResourceTag) (err error) {
	opts := tags.ActionOpts{
		Tags:   tag,
		Action: "delete",
	}
	return doAction(client, id, opts)
}
