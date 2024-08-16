package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
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

// UpdateInstanceConf is  used to modify instance configurations.
// via accessing to the service with Put method and parameters
// Send PUT /v2/{project_id}/instances/{instance_id}
func UpdateInstanceConf(client *golangsdk.ServiceClient, id string, opts UpdateInstanceConfOpts) error {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL(ResourcePath, id, configsEndpoint), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}

type UpdateInstanceConfResp struct {
	// Configuration modification task ID.
	JobID string `json:"job_id"`
	// Number of dynamic configuration parameters to be modified.
	DynamicConfig int `json:"dynamic_config"`
	// Number of static configuration parameters to be modified.
	StaticConfig int `json:"static_config"`
}
