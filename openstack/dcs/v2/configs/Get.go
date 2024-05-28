package configs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*ConfigParam, error) {
	raw, err := client.Get(client.ServiceURL("instances", id, "configs"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ConfigParam
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ConfigParam struct {
	ConfigTime   string              `json:"config_time"`
	InstanceID   string              `json:"instance_id"`
	RedisConfigs []RedisConfigResult `json:"redis_config"`
	ConfigStatus string              `json:"config_status"`
	Status       string              `json:"status"`
}

type RedisConfigResult struct {
	ParamValue   string `json:"param_value"`
	Description  string `json:"description"`
	ValueType    string `json:"value_type"`
	ValueRange   string `json:"value_range"`
	DefaultValue string `json:"default_value"`
	ParamID      string `json:"param_id"`
	ParamName    string `json:"param_name"`
}
