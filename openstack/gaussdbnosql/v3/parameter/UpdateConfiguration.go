package parameter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateConfigurationOpts struct {
	// Parameter template ID
	ConfigId string `json:"config_id"`
	// Parameter template name. It contains a maximum of 64 characters and can contain only uppercase letters,
	// lowercase letters, digits, hyphens (-), underscores (_), and periods (.).
	Name string `json:"name,omitempty"`
	// Parameter template description. It contains a maximum of 256 characters except the following special characters:
	// >!<"&'= The value is left blank by default.
	Description string `json:"description,omitempty"`
	// The parameter values defined by users based on the default parameter template.
	// If the value is empty, the parameter value is not changed.
	Values map[string]string `json:"values,omitempty"`
}

func UpdateConfiguration(client *golangsdk.ServiceClient, opts UpdateConfigurationOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT https://{Endpoint}/v3/{project_id}/configurations/{config_id}
	_, err = client.Put(client.ServiceURL("configurations", opts.ConfigId), b, nil, nil)
	return
}
