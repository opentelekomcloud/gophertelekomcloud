package dc_endpoint_group

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Endpoints   []string `json:"endpoints" required:"true"`
	Type        string   `json:"type" required:"true"`
}

func Create(c *golangsdk.ServiceClient, opts CreateOpts) (*DCEndpointGroup, error) {
	b, err := build.RequestBody(opts, "dc_endpoint_group")
	if err != nil {
		return nil, err
	}
	raw, err := c.Post(c.ServiceURL("dcaas", "dc-endpoint-groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res DCEndpointGroup
	err = extract.IntoStructPtr(raw.Body, &res, "dc_endpoint_group")
	return &res, err
}
