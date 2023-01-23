package tag

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateTagOpts struct {
	// Resource ID. Currently, you can only add tags to a cluster, so specify this parameter to the cluster ID.
	ClusterId string
	Tag       tags.ResourceTag `json:"tag"`
}

func CreateClusterTag(client *golangsdk.ServiceClient, opts CreateTagOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /v1.0/{project_id}/clusters/{resource_id}/tags
	_, err = client.Post(client.ServiceURL("clusters", opts.ClusterId, "tags"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
