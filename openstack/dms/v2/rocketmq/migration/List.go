package migration

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Current page, starting from 1.
	Offset int `q:"offset,omitempty"`
	// Current page size.
	Limit int `q:"limit,omitempty"`
}

// ListResponse is returned by List.
type ListResponse struct {
	// Total number of metadata migration tasks.
	Total int `json:"total"`
	// Metadata migration tasks.
	Tasks []Task `json:"task"`
}

// List the metadata migration tasks of a RocketMQ instance.
// Send GET to /v2/{project_id}/instances/{instance_id}/metadata
func List(client *golangsdk.ServiceClient, instanceID string, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "metadata").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
