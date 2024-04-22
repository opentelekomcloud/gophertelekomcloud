package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

func GetResourceTags(client *golangsdk.ServiceClient, funcURN string) (*TagsResp, error) {
	raw, err := client.Get(client.ServiceURL("functions", funcURN, "tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res TagsResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type TagsResp struct {
	Tags    []tags.ResourceTag `json:"tags"`
	SysTags []tags.ResourceTag `json:"sys_tags"`
}
