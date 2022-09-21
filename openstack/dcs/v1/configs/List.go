package configs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, instanceID string) (*ConfigParam, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceID, "configs"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ConfigParam
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ConfigParam struct {
	Status       string              `json:"status"`
	InstanceID   string              `json:"instance_id"`
	RedisConfigs []ResultRedisConfig `json:"redis_config"`
	ConfigStatus string              `json:"config_status"`
	ConfigTime   string              `json:"config_time"`
}

type ResultRedisConfig struct {
	Description  string `json:"description"`
	ParamID      string `json:"param_id"`
	ParamName    string `json:"param_name"`
	ParamValue   string `json:"param_value"`
	DefaultValue string `json:"default_value"`
	ValueType    string `json:"value_type"`
	ValueRange   string `json:"value_range"`
}
