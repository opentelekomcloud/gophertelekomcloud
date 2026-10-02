package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// ListResponse is returned by List.
type ListResponse struct {
	// Total number of tags.
	Total int `json:"total"`
	// Offset of the next page.
	NextOffset int `json:"next_offset"`
	// Offset of the previous page.
	PreviousOffset int `json:"previous_offset"`
	// Tag keys with all their values used in the project.
	Tags []tags.ListedTag `json:"tags"`
}

// List the tags of all RocketMQ instances in the project.
// Send GET to /v2/{project_id}/rocketmq/tags
func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("rocketmq", "tags").
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
