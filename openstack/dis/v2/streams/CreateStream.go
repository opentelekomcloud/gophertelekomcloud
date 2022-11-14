package streams

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"

type CreateStreamReq struct {
	// Name of the stream.
	// The stream name can contain 1 to 64 characters, including letters, digits, underscores (_), and hyphens (-).
	// Maximum: 64
	StreamName string `json:"stream_name"`
	// Number of partitions. Partitions are the base throughput unit of a DIS stream.
	PartitionCount int32 `json:"partition_count"`
	// Stream type.
	// COMMON: a common stream. The bandwidth is 1 MB/s.
	// ADVANCED: an advanced stream. The bandwidth is 5 MB/s.
	// Enumeration values:
	// COMMON
	// ADVANCED
	StreamType string `json:"stream_type,omitempty"`
	// Source data type.
	// BLOB: a set of binary data stored in a database management system.
	// Default value: BLOB
	// Enumeration values:
	// BLOB
	DataType string `json:"data_type,omitempty"`
	// Period of time for which data is retained in the stream. Value range: 24-72 Unit: hour Default value: 24 If this parameter is left blank, the default value is used.
	// Maximum: 7
	// Default: 24
	DataDuration *int32 `json:"data_duration,omitempty"`
	// Specifies whether to enable auto scaling.
	// true: Auto scaling is enabled.
	// false: Auto scaling is disabled.
	// This function is disabled by default.
	// Default: false
	AutoScaleEnabled *bool `json:"auto_scale_enabled,omitempty"`
	// Minimum number of partitions for automatic scale-down when auto scaling is enabled.
	// Minimum: 1
	AutoScaleMinPartitionCount *int64 `json:"auto_scale_min_partition_count,omitempty"`
	// Maximum number of partitions for automatic scale-up when auto scaling is enabled.
	AutoScaleMaxPartitionCount *int32 `json:"auto_scale_max_partition_count,omitempty"`
	//
	DataSchema string `json:"data_schema,omitempty"`
	//
	CsvProperties CsvProperties `json:"csv_properties,omitempty"`
	// Compression type of data. Currently, the value can be:
	// snappy
	// gzip
	// zip
	// Data is not compressed by default.
	// Enumeration values:
	// snappy
	// gzip
	// zip
	CompressionFormat string `json:"compression_format,omitempty"`
	// List of stream tags.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
	// Stream enterprise projects.
	SysTags []tags.ResourceTag `json:"sys_tags,omitempty"`
}

type CsvProperties struct {
	// Data separator.
	Delimiter string `json:"delimiter,omitempty"`
}

// POST /v2/{project_id}/streams

type CreateStreamResponse struct {
}
