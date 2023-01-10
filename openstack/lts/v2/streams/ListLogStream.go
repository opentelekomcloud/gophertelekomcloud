package streams

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListLogStream(client *golangsdk.ServiceClient, groupId string) ([]LogStream, error) {
	// GET /v2/{project_id}/groups/{log_group_id}/streams
	raw, err := client.Get(client.ServiceURL("groups", groupId, "streams"), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res []LogStream
	err = extract.IntoSlicePtr(raw.Body, &res, "log_streams")
	return res, err
}

type LogStream struct {
	// Creation time.
	// Minimum value: 1577808000000
	// Maximum value: 4102416000000
	CreationTime int64 `json:"creation_time"`
	// Log stream name.
	// Value length: 36 characters
	LogStreamName string `json:"log_stream_name"`
	// Log stream ID.
	// Value length: 36 characters
	LogStreamId string `json:"log_stream_id"`
	// Number of filters.
	// Minimum value: 0
	// Maximum value: 5
	FilterCount int32 `json:"filter_count"`
	// Log stream tag.
	Tag map[string]string `json:"tag,omitempty"`
}
