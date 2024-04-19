package gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// TagsUpdateOpts is the structure used to modify instance tags.
type TagsUpdateOpts struct {
	InstanceId string             `json:"-" required:"true"`
	Action     string             `json:"action" required:"true"`
	Tags       []tags.ResourceTag `json:"tags" required:"true"`
}

func UpdateTags(client *golangsdk.ServiceClient, opts *TagsUpdateOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	_, err = client.Post(client.ServiceURL("apigw", "instances", opts.InstanceId, "instance-tags/action"),
		b, nil, &golangsdk.RequestOpts{
			MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
			OkCodes:     []int{200, 201, 204},
		})
	return err
}
