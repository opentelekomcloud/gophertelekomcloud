package configs

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type ConfigParam struct {
	Status       string              `json:"status"`
	InstanceID   string              `json:"instance_id"`
	RedisConfigs []ResultRedisConfig `json:"redis_config"`
	ConfigStatus string              `json:"config_status"`
	ConfigTime   string              `json:"config_time"`
}

type ResultRedisConfig struct {
	Description  string `json:"description"`
	ParamID      int    `json:"param_id"`
	ParamName    string `json:"param_name"`
	ParamValue   string `json:"param_value"`
	DefaultValue string `json:"default_value"`
	ValueType    string `json:"value_type"`
	ValueRange   string `json:"value_range"`
}

type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (*ConfigParam, error) {
	s := new(ConfigParam)
	err := r.ExtractIntoStructPtr(s, "")
	return s, err
}

type UpdateResult struct {
	golangsdk.Result
}
