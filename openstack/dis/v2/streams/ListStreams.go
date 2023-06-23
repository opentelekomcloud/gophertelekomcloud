package streams

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type ListStreamsOpts struct {
	// The maximum number of DIS streams to list in a single API call.
	// Value range: 1-100 Default value: 10
	// Minimum: 1
	// Maximum: 100
	// Default: 10
	Limit *int `q:"limit,omitempty"`
	// Name of the DIS stream to start the stream list with. The returned stream list does not contain this DIS stream name.
	// If pagination query is required, this parameter is not transferred for query on the first page.
	// If the value of has_more_streams is true, the query is performed on the next page.
	// The value of start_stream_name is the name of the last stream in the query result of the first page.
	StartStreamName string `q:"start_stream_name,omitempty"`
}

func ListStreams(client *golangsdk.ServiceClient, opts ListStreamsOpts) (*ListStreamsResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/streams
	raw, err := client.Get(client.ServiceURL("streams")+q.String(), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"}, JSONBody: nil,
	})
	if err != nil {
		return nil, err
	}

	var res ListStreamsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListStreamsResponse struct {
	// Total number of all the DIS streams created by the current tenant.
	TotalNumber *int64 `json:"total_number,omitempty"`
	// List of the streams meeting the current requests.
	StreamNames []string `json:"stream_names,omitempty"`
	// Specify whether there are more matching DIS streams to list. Possible values:
	// true: yes
	// false: no
	// Default: false
	HasMoreStreams *bool `json:"has_more_streams,omitempty"`
	// Stream details.
	StreamInfoList []StreamInfo `json:"stream_info_list,omitempty"`
}

type StreamInfo struct {
	// Name of the stream.
	StreamName string `json:"stream_name,omitempty"`
	// Time when the stream is created. The value is a 13-bit timestamp.
	CreateTime *int64 `json:"create_time,omitempty"`
	// Period for storing data in units of hours.
	RetentionPeriod *int `json:"retention_period,omitempty"`
	// Current status of the stream. Possible values:
	// CREATING: The stream is being created.
	// RUNNING: The stream is running.
	// TERMINATING: The stream is being deleted.
	// TERMINATED: The stream has been deleted.
	// Enumeration values:
	// CREATING
	// RUNNING
	// TERMINATING
	// FROZEN
	Status string `json:"status,omitempty"`
	// Stream type.
	// COMMON: a common stream. The bandwidth is 1 MB/s.
	// ADVANCED: an advanced stream. The bandwidth is 5 MB/s.
	// Enumeration values:
	// COMMON
	// ADVANCED
	StreamType string `json:"stream_type,omitempty"`
	// Source data type.
	// BLOB: a collection of binary data stored as a single entity in a database management system.
	// JSON: an open-standard file format that uses human-readable text to transmit data objects
	// consisting of attribute-value pairs and array data types.
	// CSV: a simple text format for storing tabular data in a plain text file. Commas are used as delimiters.
	// Default value: BLOB
	// Enumeration values:
	// BLOB
	// JSON
	// CSV
	DataType string `json:"data_type,omitempty"`
	// Quantity of partitions. Partitions are the base throughput unit of a DIS stream.
	PartitionCount *int `json:"partition_count,omitempty"`
	// Specifies whether to enable auto-scaling.
	// true: auto-scaling is enabled.
	// false: auto-scaling is disabled.
	// This function is disabled by default.
	// Default: false
	AutoScaleEnabled *bool `json:"auto_scale_enabled,omitempty"`
	// Minimum number of partitions for automatic scale-down when auto scaling is enabled.
	// Minimum: 1
	AutoScaleMinPartitionCount *int `json:"auto_scale_min_partition_count,omitempty"`
	// Maximum number of partitions for automatic scale-up when auto scaling is enabled.
	AutoScaleMaxPartitionCount *int `json:"auto_scale_max_partition_count,omitempty"`
	// List of tags for the newly created DIS stream.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}
