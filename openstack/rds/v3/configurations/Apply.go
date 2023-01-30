package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ApplyOpts contains all the instances needed to apply another template.
type ApplyOpts struct {
	// Specifies the parameter template ID.
	ConfigId string
	// Specifies the DB instance ID list object.
	InstanceIDs []string `json:"instance_ids" required:"true"`
}

// Apply is used to apply a parameter template to one or more DB instances.
func Apply(client *golangsdk.ServiceClient, opts ApplyOpts) (*ApplyResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT https://{Endpoint}/v3/{project_id}/configurations/{config_id}/apply
	raw, err := client.Put(client.ServiceURL("configurations", opts.ConfigId, "apply"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}

	var res ApplyResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ApplyResponse struct {
	// Specifies the parameter template ID.
	ConfigurationID string `json:"configuration_id"`
	// Specifies the parameter template name.
	ConfigurationName string `json:"configuration_name"`
	// Specifies the result of applying the parameter template.
	ApplyResults []ApplyResult `json:"apply_results"`
	// Specifies whether each parameter template is applied to DB instances successfully.
	Success bool   `json:"success"`
	JobId   string `json:"job_id"`
}

type ApplyResult struct {
	// Indicates the DB instance ID.
	InstanceID string `json:"instance_id"`
	// Indicates the DB instance name.
	InstanceName string `json:"instance_name"`
	// Indicates whether a reboot is required.
	RestartRequired bool `json:"restart_required"`
	// Indicates whether each parameter template is applied to DB instances successfully.
	Success bool `json:"success"`
}
