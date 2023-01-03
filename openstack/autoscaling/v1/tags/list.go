package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

func List(client *golangsdk.ServiceClient) (*ResourceTags, error) {
	raw, err := client.Get(client.ServiceURL("scaling_group_tag", "tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ResourceTags
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ResourceTags struct {
	Tags    []tags.ResourceTag `json:"tags"`
	SysTags []tags.ResourceTag `json:"sys_tags"`
}
