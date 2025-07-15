package streams

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type UpdateLogStreamOpts struct {
	GroupId  string `json:"-" required:"true"`
	StreamId string `json:"-" required:"true"`
	// Log retention duration, in days (fixed to 7 days).
	TTLInDays int `json:"ttl_in_days" required:"true"`
	// Tags - tags of a group.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateLogStreamOpts) (*LogStreamUpdateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/groups/{log_group_id}/streams-ttl/{log_stream_id}
	raw, err := client.Put(client.ServiceURL("groups", opts.GroupId, "streams-ttl", opts.StreamId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res LogStreamUpdateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type LogStreamUpdateResponse struct {
	// Creation time.
	// Minimum value: 1577808000000
	// Maximum value: 4102416000000
	CreationTime int64 `json:"creation_time"`
	// Log stream name.
	// Value length: 36 characters
	LogStreamName string `json:"log_topic_name"`
	// Log stream ID.
	// Value length: 36 characters
	LogStreamId string `json:"log_topic_id"`
	// Log retention duration, in days.
	TtlInDays int `json:"ttl_in_days"`
}
