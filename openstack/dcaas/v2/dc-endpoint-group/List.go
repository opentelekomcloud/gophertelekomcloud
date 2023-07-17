package dc_endpoint_group

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

type DCEndpointGroup struct {
	// Specifies the project ID.
	TenantId string `json:"tenant_id"`
	// Specifies the ID of the Direct Connect endpoint group.
	ID string `json:"id"`
	// Specifies the name of the Direct Connect endpoint group.
	Name string `json:"name"`
	// Provides supplementary information about the Direct Connect endpoint group.
	Description string `json:"description"`
	// Specifies the list of the endpoints in a Direct Connect endpoint group.
	Endpoints []string `json:"endpoints"`
	// Specifies the type of the Direct Connect endpoints. The value can only be cidr.
	Type string `json:"type"`
}

type ListOpts struct {
	ID string `q:"id"`
}

// List is used to obtain the DirectConnects list
func List(c *golangsdk.ServiceClient, opts ListOpts) ([]DCEndpointGroup, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v2.0/dcaas/dc-endpoint-groups?id={id}
	raw, err := c.Get(c.ServiceURL("dcaas", "dc-endpoint-groups")+q.String(), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res []DCEndpointGroup
	err = extract.IntoSlicePtr(raw.Body, &res, "dc_endpoint_groups")
	return res, err
}
