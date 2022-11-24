package configs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Array of configuration items of the DCS instance.
	RedisConfigs []RedisConfig `json:"redis_config" required:"true"`
}

type RedisConfig struct {
	// Configuration item ID.
	ParamID string `json:"param_id" required:"true"`
	// Configuration item name.
	ParamName string `json:"param_name" required:"true"`
	// Value of the configuration item.
	ParamValue string `json:"param_value" required:"true"`
}

func Update(client *golangsdk.ServiceClient, instanceID string, opts UpdateOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	_, err = client.Put(client.ServiceURL("instances", instanceID, "configs"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
