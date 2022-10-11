package streams

type ListStreamsRequest struct {
	// The maximum number of DIS streams to list in a single API call.
	// Value range: 1-100 Default value: 10
	// Minimum: 1
	// Maximum: 500
	// Default: 10
	Limit *int32 `json:"limit,omitempty"`
	// Name of the DIS stream to start the stream list with. The returned stream list does not contain this DIS stream name.
	// If pagination query is required, this parameter is not transferred for query on the first page.
	// If the value of has_more_streams is true, the query is performed on the next page.
	// The value of start_stream_name is the name of the last stream in the query result of the first page.
	StartStreamName *string `json:"start_stream_name,omitempty"`
}

// GET /v2/{project_id}/streams

type ListStreamsResponse struct {
	// Total number of all the DIS streams created by the current tenant.
	TotalNumber *int64 `json:"total_number,omitempty"`
	// List of the streams meeting the current requests.
	StreamNames *[]string `json:"stream_names,omitempty"`
	// Specify whether there are more matching DIS streams to list. Possible values:
	// true: yes
	// false: no
	// Default: false
	HasMoreStreams *bool `json:"has_more_streams,omitempty"`
	// Stream details.
	StreamInfoList *[]StreamInfo `json:"stream_info_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}
type StreamInfo struct {
	// Name of the stream.
	StreamName *string `json:"stream_name,omitempty"`
	// Time when the stream is created. The value is a 13-bit timestamp.
	CreateTime *int64 `json:"create_time,omitempty"`
	// Period for storing data in units of hours.
	RetentionPeriod *int32 `json:"retention_period,omitempty"`
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
	Status *StreamInfoStatus `json:"status,omitempty"`
	// Stream type.
	// COMMON: a common stream. The bandwidth is 1 MB/s.
	// ADVANCED: an advanced stream. The bandwidth is 5 MB/s.
	// Enumeration values:
	// COMMON
	// ADVANCED
	StreamType *StreamInfoStreamType `json:"stream_type,omitempty"`
	// Source data type.
	// BLOB: a collection of binary data stored as a single entity in a database management system.
	// Default value: BLOB
	// Enumeration values:
	// BLOB
	DataType *StreamInfoDataType `json:"data_type,omitempty"`
	// Quantity of partitions. Partitions are the base throughput unit of a DIS stream.
	PartitionCount *int32 `json:"partition_count,omitempty"`
	// List of tags for the newly created DIS stream.
	Tags *[]Tag `json:"tags,omitempty"`
	// Specifies whether to enable auto scaling.
	// true: auto scaling is enabled.
	// false: auto scaling is disabled.
	// This function is disabled by default.
	// Default: false
	AutoScaleEnabled *bool `json:"auto_scale_enabled,omitempty"`
	// Minimum number of partitions for automatic scale-down when auto scaling is enabled.
	// Minimum: 1
	AutoScaleMinPartitionCount *int32 `json:"auto_scale_min_partition_count,omitempty"`
	// Maximum number of partitions for automatic scale-up when auto scaling is enabled.
	AutoScaleMaxPartitionCount *int32 `json:"auto_scale_max_partition_count,omitempty"`
}
type StreamInfoStatus struct {
	value string
}
type StreamInfoStreamType struct {
	value string
}
type StreamInfoDataType struct {
	value string
}
