package groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Consumer group name.
	Group string `q:"group,omitempty"`
	// Number of records to query. Default: 10.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
}

// ListResponse is returned by List.
type ListResponse struct {
	// Total number of consumer groups.
	Total int `json:"total"`
	// Consumer group list.
	Groups []ConsumerGroup `json:"groups"`
	// Maximum number of consumer groups that can be created.
	Max int `json:"max"`
	// Remaining number of consumer groups that can be created.
	Remaining int `json:"remaining"`
	// Offset of the next page.
	NextOffset int `json:"next_offset"`
	// Offset of the previous page.
	PreviousOffset int `json:"previous_offset"`
}

// List the consumer groups of a RocketMQ instance.
// Send GET to /v2/{project_id}/instances/{instance_id}/groups
func List(client *golangsdk.ServiceClient, instanceID string, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "groups").
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
