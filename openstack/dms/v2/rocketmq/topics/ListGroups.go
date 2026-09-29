package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListGroupsOpts struct {
	// Maximum number of records returned in a query. Range: 1-50. Default: 10.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
}

// ListGroupsResponse is returned by ListGroups.
type ListGroupsResponse struct {
	// Consumer group list.
	Groups []string `json:"groups"`
	// Total number of consumer groups.
	Total int `json:"total"`
}

// ListGroups queries the consumer group list of a topic.
// Send GET to /v2/{project_id}/instances/{instance_id}/topics/{topic}/groups
func ListGroups(client *golangsdk.ServiceClient, instanceID, topic string, opts ListGroupsOpts) (*ListGroupsResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "topics", topic, "groups").
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

	var res ListGroupsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
