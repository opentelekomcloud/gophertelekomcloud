package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateInstanceConfigurationOpts struct {
	// Specifies the DB instance ID.
	InstanceId string `json:"-"`
	// Specifies the parameter values defined by users based on the default parameter templates.
	// For example, "max_connections": "10"
	Values map[string]interface{} `json:"values"`
}

func UpdateInstanceConfiguration(client *golangsdk.ServiceClient, opts UpdateInstanceConfigurationOpts) (*UpdateInstanceConfigurationResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/configurations
	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "configurations"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateInstanceConfigurationResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type UpdateInstanceConfigurationResponse struct {
	// Indicates whether a reboot is required.
	RestartRequired bool   `json:"restart_required"`
	JobId           string `json:"job_id"`
	// List of ignored parameters.
	// If a parameter does not exist or is read-only, the parameter cannot be modified and the names of all ignored parameters are returned by ignored_params.
	IgnoredParams []string `json:"ignored_params"`
}
