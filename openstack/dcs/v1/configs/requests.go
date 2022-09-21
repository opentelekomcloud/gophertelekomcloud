package configs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type UpdateOptsBuilder interface {
	ToConfigUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	RedisConfigs []RedisConfig `json:"redis_config" required:"true"`
}

type RedisConfig struct {
	ParamID    string `json:"param_id" required:"true"`
	ParamName  string `json:"param_name" required:"true"`
	ParamValue string `json:"param_value" required:"true"`
}

func (opts UpdateOpts) ToConfigUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}
