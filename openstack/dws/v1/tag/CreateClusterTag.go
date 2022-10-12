package tag

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateTagOpts struct {
	// Resource ID. Currently, you can only add tags to a cluster, so specify this parameter to the cluster ID.
	ClusterId string
	Tag       Tag `json:"tag"`
}

type Tag struct {
	// Tag key. A tag key can contain a maximum of 36 Unicode characters, which cannot be null. The first and last characters cannot be spaces.
	// It can contain uppercase letters (A to Z), lowercase letters (a to z), digits (0-9), hyphens (-), and underscores (_).
	Key string `json:"key"`
	// Key value. A tag value can contain a maximum of 43 Unicode characters, which can be null. The first and last characters cannot be spaces.
	// It can contain uppercase letters (A to Z), lowercase letters (a to z), digits (0-9), hyphens (-), and underscores (_).
	Value string `json:"value"`
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
