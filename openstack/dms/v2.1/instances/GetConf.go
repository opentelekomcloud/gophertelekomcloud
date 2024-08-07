package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const configsEndpoint = "configs"

// GetInstanceConf is used to obtain instance configurations.
// Send GET /v2/{project_id}/instances/{instance_id}/configs
func GetInstanceConf(client *golangsdk.ServiceClient, id string) (*InstanceConfs, error) {
	raw, err := client.Get(client.ServiceURL(resourcePath, id, configsEndpoint), nil, nil)
	if err != nil {
		return nil, err
	}

	var res InstanceConfs
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type InstanceConfs struct {
	KafkaConfigs []InstanceConf `json:"kafka_configs"`
}

type InstanceConf struct {
	// Configuration name.
	Name string `json:"name"`
	// Valid value.
	ValidValues string `json:"valid_values"`
	// Default value.
	DefaultValue string `json:"default_value"`
	// Configuration type. The value can be static or dynamic.
	ConfigType string `json:"config_type"`
	// Current value.
	Value string `json:"value"`
	// Value type.
	ValueType string `json:"value_type"`
}
