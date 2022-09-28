package parameter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateInstanceConfigurationOpts struct {
	InstanceId string
	// The parameter values defined by users based on the default parameter template
	Values map[string]string `json:"values"`
}

func UpdateInstanceConfiguration(client *golangsdk.ServiceClient, opts UpdateInstanceConfigurationOpts) (*UpdateInstanceConfigurationResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/configurations
	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "configurations"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res UpdateInstanceConfigurationResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type UpdateInstanceConfigurationResponse struct {
	// ID of the asynchronous task for modifying instance parameters
	JobId string `json:"job_id,omitempty"`
	// Whether the instance needs to be restarted.
	RestartRequired bool `json:"restart_required,omitempty"`
}
