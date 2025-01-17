package parameter_configuration

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, clusterId string) (*Configurations, error) {
	// GET /v1.0/{project_id}/clusters/{cluster_id}/ymls/template
	raw, err := client.Get(client.ServiceURL("clusters", clusterId, "ymls", "template"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Configurations
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Configurations struct {
	Templates map[string]Template `json:"configurations"`
}

type Template struct {
	// parameter ID
	ID string `json:"id"`
	// parameter name
	Key string `json:"key"`
	// parameter value
	Value string `json:"value"`
	// parameter default value.
	DefaultValue string `json:"defaultValue"`
	// parameter constraint
	Regex string `json:"regex"`
	// parameter type description
	Type string `json:"type"`
	// indicates whether a parameter can be modified.
	// The value can be true (parameter value can be changed) and false (parameter value cannot be changed).
	ModifyEnable string `json:"modifyEnable"`
	// parameter value that can be changed
	EnableValue string `json:"enableValue"`
	// name of the file where parameters exist. The default value is elasticsearch.yml
	FileName string `json:"fileName"`
	// version information
	Version string `json:"version"`
	// parameter description
	Description string `json:"descENG"`
	// parameter function description
	ModuleDescription string `json:"moduleDescENG"`
}
