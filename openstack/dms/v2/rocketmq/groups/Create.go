package groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Consumer group name. Enter 3 to 64 characters. Use only letters, digits,
	// percent (%), vertical bars (|), hyphens (-), and underscores (_).
	Name string `json:"name" required:"true"`
	// Whether to broadcast.
	Broadcast *bool `json:"broadcast,omitempty"`
	// Maximum number of retries. Range: 1-16.
	RetryMaxTime int `json:"retry_max_time,omitempty"`
	// Whether consumption is allowed.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether orderly consumption is enabled.
	// Required only for RocketMQ 5.x instances.
	ConsumeOrderly *bool `json:"consume_orderly,omitempty"`
	// Consumer group description. 0-200 characters.
	Description string `json:"group_desc,omitempty"`
}

// CreateResponse is returned by Create.
type CreateResponse struct {
	// Name of the created consumer group.
	Name string `json:"name"`
}

// Create a consumer group of a RocketMQ instance.
// Send POST to /v2/{project_id}/instances/{instance_id}/groups
func Create(client *golangsdk.ServiceClient, instanceID string, opts CreateOpts) (*CreateResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "groups").Build()
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
