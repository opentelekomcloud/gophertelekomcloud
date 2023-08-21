package dc_endpoint_group

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(c *golangsdk.ServiceClient, id string) (*DCEndpointGroup, error) {
	raw, err := c.Get(c.ServiceURL("dcaas", "dc-endpoint-groups", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res DCEndpointGroup
	err = extract.IntoStructPtr(raw.Body, &res, "dc_endpoint_group")
	return &res, err
}
