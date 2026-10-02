package configs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// RocketMQ configurations to modify.
	Configs []UpdateConfig `json:"rocketmq_configs" required:"true"`
}

type UpdateConfig struct {
	// Configuration name.
	Name string `json:"name" required:"true"`
	// Target value.
	Value string `json:"value" required:"true"`
}

// Update the configurations of a RocketMQ instance.
// Send PUT to /v2/{project_id}/rocketmq/instances/{instance_id}/configs
func Update(client *golangsdk.ServiceClient, instanceID string, opts UpdateOpts) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("rocketmq", "instances", instanceID, "configs").Build()
	if err != nil {
		return err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
