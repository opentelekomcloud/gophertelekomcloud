package configs

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type ConfigParameters struct {
	Status              string               `json:"status"`
	InstanceID          string               `json:"instance_id"`
	RedisConfigurations []RedisConfiguration `json:"redis_config"`
	ConfigurationStatus string               `json:"config_status"`
	ConfigurationTime   string               `json:"config_time"`
}

type RedisConfiguration struct {
	Description    string `json:"description"`
	ParameterID    int    `json:"param_id"`
	ParameterName  string `json:"param_name"`
	ParameterValue string `json:"param_value"`
	DefaultValue   string `json:"default_value"`
	ValueType      string `json:"value_type"`
	ValueRange     string `json:"value_range"`
}

type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (*ConfigParameters, error) {
	s := new(ConfigParameters)
	err := r.ExtractIntoStructPtr(s, "")
	return s, err
}
