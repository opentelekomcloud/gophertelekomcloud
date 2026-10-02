package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type ListOpts struct {
	// Number of records to query. Default: 10.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
}

// GetResponse is returned by Get.
type GetResponse struct {
	// Total number of tags.
	Total int `json:"total"`
	// Offset of the next page.
	NextOffset int `json:"next_offset"`
	// Offset of the previous page.
	PreviousOffset int `json:"previous_offset"`
	// Tag list.
	Tags []tags.ResourceTag `json:"tags"`
}

// Get the tags of a RocketMQ instance.
// Send GET to /v2/{project_id}/rocketmq/{instance_id}/tags
func Get(client *golangsdk.ServiceClient, instanceID string, opts ListOpts) (*GetResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("rocketmq", instanceID, "tags").
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

	var res GetResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
