package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// UpdateOpts contains all the values needed to update a Configuration.
type UpdateOpts struct {
	// Specifies the parameter template ID.
	ConfigId string
	// Specifies the parameter template name. It contains a maximum of 64 characters and can contain only uppercase letters, lowercase letters, digits, hyphens (-), underscores (_), and periods (.).
	Name string `json:"name,omitempty"`
	// Specifies the parameter template description. It contains a maximum of 256 characters and does not support the following special characters: !<>='&" Its value is left blank by default.
	Description string `json:"description,omitempty"`
	// Specifies the parameter values defined by users based on the default parameter template. If this parameter is left blank, the parameter value cannot be changed.
	Values map[string]string `json:"values,omitempty"`
}

// Update accepts a UpdateOpts struct and uses the values to update a Configuration.The response code from api is 200
func Update(c *golangsdk.ServiceClient, opts UpdateOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT https://{Endpoint}/v3/{project_id}/configurations/{config_id}
	_, err = c.Put(c.ServiceURL("configurations", opts.ConfigId), b, nil,
		&golangsdk.RequestOpts{OkCodes: []int{200}, MoreHeaders: openstack.StdRequestOpts().MoreHeaders})

	return
}
