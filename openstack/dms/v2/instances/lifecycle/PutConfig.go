package lifecycle

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateInstanceConfOpts is a struct which represents the parameters of update function
type UpdateInstanceConfOpts struct {
	// Configurations to be modified.
	KafkaConfigs []KafkaConfig `json:"kafka_configs,omitempty"`
}

type KafkaConfig struct {
	// Names of configurations to be modified.
	Name string `json:"name"`
	// New value of the modified configuration.
	Value string `json:"value"`
}

// UpdateInstanceConf is used to modify instance configurations.
// Send PUT /v2/{project_id}/instances/{instance_id}/configs
func UpdateInstanceConf(client *golangsdk.ServiceClient, id string, opts UpdateInstanceConfOpts) (*UpdateInstanceConfResp, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", id, "configs"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateInstanceConfResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type UpdateInstanceConfResp struct {
	// Configuration modification task ID.
	JobID string `json:"job_id"`
	// Number of dynamic configuration parameters to be modified.
	DynamicConfig int `json:"dynamic_config"`
	// Number of static configuration parameters to be modified.
	StaticConfig int `json:"static_config"`
}
