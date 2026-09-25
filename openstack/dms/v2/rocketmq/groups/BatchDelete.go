package groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchDeleteOpts struct {
	// List of consumer groups to be deleted.
	Groups []string `json:"groups" required:"true"`
}

type deleteAction struct {
	Action string `q:"action"`
}

// BatchDeleteResponse is returned by BatchDelete.
type BatchDeleteResponse struct {
	// ID of the consumer group deletion task.
	JobID string `json:"job_id"`
}

// BatchDelete deletes consumer groups of a RocketMQ instance in batches.
// Send POST to /v2/{project_id}/instances/{instance_id}/groups?action=delete
func BatchDelete(client *golangsdk.ServiceClient, instanceID string, opts BatchDeleteOpts) (*BatchDeleteResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "groups").
		WithQueryParams(&deleteAction{Action: "delete"}).
		Build()
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

	var res BatchDeleteResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
