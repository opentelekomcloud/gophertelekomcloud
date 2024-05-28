package configs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ModifyConfigOpt struct {
	InstanceId  string         `json:"-"`
	RedisConfig []RedisConfigs `json:"redis_config"`
}

type RedisConfigs struct {
	ParamID    string `json:"param_id" required:"true"`
	ParamName  string `json:"param_name" required:"true"`
	ParamValue string `json:"param_value" required:"true"`
}

func Update(client *golangsdk.ServiceClient, opts ModifyConfigOpt) (err error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "configs"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
