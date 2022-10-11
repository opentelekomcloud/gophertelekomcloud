package streams

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
	StreamType *CreateStreamReqStreamType `json:"stream_type,omitempty"`
	// Source data type.
	// BLOB: a set of binary data stored in a database management system.
	// Default value: BLOB
	// Enumeration values:
	// BLOB
	DataType *CreateStreamReqDataType `json:"data_type,omitempty"`
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
	DataSchema *string `json:"data_schema,omitempty"`
	//
	CsvProperties *CsvProperties `json:"csv_properties,omitempty"`
	// Compression type of data. Currently, the value can be:
	// snappy
	// gzip
	// zip
	// Data is not compressed by default.
	// Enumeration values:
	// snappy
	// gzip
	// zip
	CompressionFormat *CreateStreamReqCompressionFormat `json:"compression_format,omitempty"`
	// List of stream tags.
	Tags *[]Tag `json:"tags,omitempty"`
	// Stream enterprise projects.
	SysTags *[]SysTag `json:"sys_tags,omitempty"`
}
type CreateStreamReqStreamType struct {
	value string
}
type CreateStreamReqDataType struct {
	value string
}
type CsvProperties struct {
	//
	Delimiter *string `json:"delimiter,omitempty"`
}
type CreateStreamReqCompressionFormat struct {
	value string
}
type Tag struct {
	// Key.
	// This field cannot be left blank.
	// The key value of a resource must be unique.
	// Character set: A-Z, a-z, 0-9, '-', '_', and Unicode characters (\u4E00-\u9FFF).
	// Minimum: 1
	// Maximum: 36
	Key *string `json:"key,omitempty"`
	// Value.
	// The value contains a maximum of 43 characters.
	// Character set: A-Z, a-z, 0-9, '. ', '-', '_', and Unicode characters (\u4E00-\u9FFF).
	// The value can contain only digits, letters, hyphens (-), and underscores (_).
	// Minimum: 0
	// Maximum: 43
	Value *string `json:"value,omitempty"`
}

type SysTag struct {
	// Key.
	// This field cannot be left blank.
	// The value must be _sys_enterprise_project_id.
	// Enumeration values:
	// _sys_enterprise_project_id
	Key *SysTagKey `json:"key,omitempty"`
	// Value. The value is the enterprise project ID, which needs to be obtained on the enterprise management page.
	// 36-digit UUID
	Value *string `json:"value,omitempty"`
}
type SysTagKey struct {
	value string
}

// POST /v2/{project_id}/streams

type CreateStreamResponse struct {
	HttpStatusCode int `json:"-"`
}
