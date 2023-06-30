package members

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/others"
)

func BatchDeleteMembers(client *golangsdk.ServiceClient, opts BatchMembersOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// DELETE /v1/cloudimages/members
	raw, err := client.DeleteWithBody(client.ServiceURL("cloudimages", "members"), b, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return others.ExtractJobId(err, raw)
}
