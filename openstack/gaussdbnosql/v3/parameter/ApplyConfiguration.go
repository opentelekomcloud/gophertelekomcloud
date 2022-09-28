package parameter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ApplyConfigurationRequest struct {
	// Parameter template ID
	ConfigId string
	// Instance IDs
	InstanceIds []string `json:"instance_ids"`
}

func ApplyConfiguration(client *golangsdk.ServiceClient, opts UpdateConfigurationOpts) (*ApplyConfigurationResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT https://{Endpoint}/v3/{project_id}/configurations/{config_id}/apply
	raw, err := client.Put(client.ServiceURL("configurations", opts.ConfigId, "apply"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res ApplyConfigurationResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ApplyConfigurationResponse struct {
	// Asynchronous task ID of the application parameter template
	JobId string `json:"job_id,omitempty"`
	// Whether the parameter template application task is successfully submitted.
	// If the value is true, the task is successfully submitted.
	// If the value is false, the task fails to be submitted.
	Success bool `json:"success,omitempty"`
}
