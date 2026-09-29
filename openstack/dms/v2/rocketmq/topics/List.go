package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Number of records to query. Range: 1-50.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
}

// ListResponse is returned by List.
type ListResponse struct {
	// Total number of topics.
	Total int `json:"total"`
	// Maximum number of topics that can be created.
	Max int `json:"max"`
	// Remaining number of topics that can be created.
	Remaining int `json:"remaining"`
	// Offset of the next page.
	NextOffset int `json:"next_offset"`
	// Offset of the previous page.
	PreviousOffset int `json:"previous_offset"`
	// Topic list.
	Topics []Topic `json:"topics"`
}

// List the topics of a RocketMQ instance.
// Send GET to /v2/{project_id}/instances/{instance_id}/topics
func List(client *golangsdk.ServiceClient, instanceID string, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "topics").
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
