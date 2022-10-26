package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*ResourceTags, error) {
	raw, err := client.Get(client.ServiceURL("scaling_group_tag", id, "tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ResourceTags
	err = extract.Into(raw.Body, &res)
	return &res, err
}
