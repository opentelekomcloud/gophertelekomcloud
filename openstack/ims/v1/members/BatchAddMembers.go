package members

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/others"
)

type BatchAddMembersOpts struct {
	// Specifies the image IDs.
	Images []string `json:"images"`
	// Specifies the project IDs.
	Projects []string `json:"projects"`
}

func BatchAddMembers(client *golangsdk.ServiceClient, opts BatchAddMembersOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/cloudimages/members
	raw, err := client.Post(client.ServiceURL("cloudimages", "members"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return others.ExtractJobId(err, raw)
}
