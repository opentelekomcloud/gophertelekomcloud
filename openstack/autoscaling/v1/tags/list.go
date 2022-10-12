package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
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
	Tags    []ResourceTag `json:"tags"`
	SysTags []ResourceTag `json:"sys_tags"`
}
