package tag

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type BatchCreateClusterTagsOpts struct {
	// Resource ID, for example, 7d85f602-a948-4a30-afd4-e84f47471c15.
	ClusterId string
	// Identifies the operation. The value can be created or delete.
	// create: adds tags in batches.
	// delete: deletes tags in batches.
	Action string `json:"action"`
	// Tag list.
	Tags []tags.ResourceTag `json:"tags"`
}

func BatchAddDeleteTags(client *golangsdk.ServiceClient, opts BatchCreateClusterTagsOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /v1.0/{project_id}/clusters/{resource_id}/tags/action
	_, err = client.Post(client.ServiceURL("clusters", opts.ClusterId, "tags", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
