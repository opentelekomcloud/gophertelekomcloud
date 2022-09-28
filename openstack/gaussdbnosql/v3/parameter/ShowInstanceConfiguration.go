package parameter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowInstanceConfiguration(client *golangsdk.ServiceClient, instanceId string) (*ShowInstanceConfigurationResponse, error) {
	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/configurations
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "configurations"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ShowInstanceConfigurationResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ShowInstanceConfigurationResponse struct {
	// DB version name
	DatastoreVersionName string `json:"datastore_version_name,omitempty"`
	// Database name
	DatastoreName string `json:"datastore_name,omitempty"`
	// Creation time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	Created string `json:"created,omitempty"`
	// Update time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	Updated string `json:"updated,omitempty"`
	// The parameters defined by users based on the default parameter templates.
	ConfigurationParameters []ConfigurationParameterResult `json:"configuration_parameters,omitempty"`
}

type ConfigurationParameterResult struct {
	// Parameter name
	Name string `json:"name"`
	// Parameter value
	Value string `json:"value"`
	// Whether the instance needs to be restarted.
	RestartRequired bool `json:"restart_required"`
	// Whether the parameter is read-only.
	Readonly bool `json:"readonly"`
	// Value range. For example, the value of the Integer type ranges from 0 to 1,
	// and the value of the Boolean type is true or false.
	ValueRange string `json:"value_range"`
	// Parameter type. The value can be string, integer, boolean, list, or float.
	Type string `json:"type"`
	// Parameter description
	Description string `json:"description"`
}
