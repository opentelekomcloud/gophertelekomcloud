package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type TagsActionOpts struct {
	Id      string             `json:"-"`
	Action  string             `json:"action"`
	Tags    []tags.ResourceTag `json:"tags"`
	SysTags []tags.ResourceTag `json:"sys_tags"`
}

func CreateResourceTag(client *golangsdk.ServiceClient, opts TagsActionOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("functions", opts.Id, "tags", "create"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
