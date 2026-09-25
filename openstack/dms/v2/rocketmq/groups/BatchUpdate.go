package groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchUpdateOpts struct {
	// Consumer group list.
	Groups []BatchUpdateGroup `json:"groups" required:"true"`
}

type BatchUpdateGroup struct {
	// Consumer group name.
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
	Description *string `json:"group_desc,omitempty"`
}

// BatchUpdateResponse is returned by BatchUpdate.
type BatchUpdateResponse struct {
	// Task ID.
	JobID string `json:"job_id"`
}

// BatchUpdate modifies consumer groups of a RocketMQ instance in batches.
// Send PUT to /v2/{project_id}/instances/{instance_id}/groups
func BatchUpdate(client *golangsdk.ServiceClient, instanceID string, opts BatchUpdateOpts) (*BatchUpdateResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "groups").Build()
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res BatchUpdateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
