package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Number of records to query.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
}

// ListResponse is returned by List.
type ListResponse struct {
	// User list.
	Users []User `json:"users"`
	// Total number of users.
	Total int `json:"total"`
}

// List the users of a RocketMQ instance.
// Send GET to /v2/{project_id}/instances/{instance_id}/users
func List(client *golangsdk.ServiceClient, instanceID string, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "users").
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
