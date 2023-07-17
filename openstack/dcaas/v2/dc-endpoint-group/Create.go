package dc_endpoint_group

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Specifies the project ID.
	TenantId string `json:"tenant_id" required:"true"`
	// Specifies the name of the Direct Connect endpoint group.
	Name string `json:"name,omitempty"`
	// Provides supplementary information about the Direct Connect endpoint group.
	Description string `json:"description,omitempty"`
	// Specifies the list of the endpoints in a Direct Connect endpoint group.
	Endpoints []string `json:"endpoints" required:"true"`
	// Specifies the type of the Direct Connect endpoints. The value can only be cidr.
	Type string `json:"type,omitempty"`
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
