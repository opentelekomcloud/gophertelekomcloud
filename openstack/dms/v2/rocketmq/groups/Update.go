package groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Whether a message can be consumed.
	Enabled *bool `json:"enabled" required:"true"`
	// Whether to broadcast.
	Broadcast *bool `json:"broadcast" required:"true"`
	// List of associated brokers.
	Brokers []string `json:"brokers,omitempty"`
	// Consumer group whose parameters are to be modified.
	// The consumer group name cannot be modified.
	Name string `json:"name,omitempty"`
	// Maximum number of retries. Range: 1-16.
	RetryMaxTime int `json:"retry_max_time" required:"true"`
}

// Update the parameters of a specified consumer group of a RocketMQ instance.
// Send PUT to /v2/{project_id}/instances/{instance_id}/groups/{group}
func Update(client *golangsdk.ServiceClient, instanceID, group string, opts UpdateOpts) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "groups", group).Build()
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
