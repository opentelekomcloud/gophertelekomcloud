package streams

type DescribeStreamOpts struct {
	// Stream to be queried.
	// Maximum: 60
	StreamName string `json:"stream_name"`
	// Name of the partition to start the partition list with. The returned partition list does not contain this partition.
	StartPartitionId string `json:"start_partitionId,omitempty"`
	// Maximum number of partitions to list in a single API call. Value range: 1-1,000 Default value: 100
	// Minimum: 1
	// Maximum: 1000
	// Default: 100
	LimitPartitions *int32 `json:"limit_partitions,omitempty"`
}

// GET /v2/{project_id}/streams/{stream_name}

type DescribeStreamResponse struct {
	// Name of the stream.
	StreamName string `json:"stream_name,omitempty"`
	// Time when a stream is created. The value is a 13-bit timestamp.
	CreateTime *int64 `json:"create_time,omitempty"`
	// Time when a stream is the most recently modified. The value is a 13-bit timestamp.
	LastModifiedTime *int64 `json:"last_modified_time,omitempty"`
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
	// A list of partitions that comprise the DIS stream.
	Partitions []PartitionResult `json:"partitions,omitempty"`
	// Specifies whether there are more matching partitions of the DIS stream to list.
	// true: yes
	// false: no
	HasMorePartitions *bool `json:"has_more_partitions,omitempty"`
	// Period for storing data in units of hours.
	RetentionPeriod *int32 `json:"retention_period,omitempty"`
	// Unique identifier of the stream.
	StreamId string `json:"stream_id,omitempty"`
	// Source data type.
	// BLOB: a set of binary data stored in a database management system.
	// Default value: BLOB
	// Enumeration values:
	// BLOB
	DataType string `json:"data_type,omitempty"`
	//
	DataSchema string `json:"data_schema,omitempty"`
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
	// Attributes of data in CSV format, such as delimiter.
	CsvProperties CsvProperties `json:"csv_properties,omitempty"`
	// Total number of writable partitions (including partitions in ACTIVE state only).
	WritablePartitionCount *int32 `json:"writable_partition_count,omitempty"`
	// Total number of readable partitions (including partitions in ACTIVE and DELETED state).
	ReadablePartitionCount *int32 `json:"readable_partition_count,omitempty"`
	// List of scaling operation records.
	UpdatePartitionCounts []UpdatePartitionCount `json:"update_partition_counts,omitempty"`
	// List of stream tags.
	Tags []Tag `json:"tags,omitempty"`
	// Enterprise project of a stream.
	SysTags []SysTag `json:"sys_tags,omitempty"`
	// Specifies whether to enable auto scaling.
	// true: auto scaling is enabled.
	// false: auto scaling is disabled.
	// This function is disabled by default.
	AutoScaleEnabled *bool `json:"auto_scale_enabled,omitempty"`
	// Minimum number of partitions for automatic scale-down when auto scaling is enabled.
	AutoScaleMinPartitionCount *int32 `json:"auto_scale_min_partition_count,omitempty"`
	// Maximum number of partitions for automatic scale-up when auto scaling is enabled.
	AutoScaleMaxPartitionCount *int32 `json:"auto_scale_max_partition_count,omitempty"`
}

type PartitionResult struct {
	// Current status of the partition. Possible values:
	// CREATING: The stream is being created.
	// ACTIVE: The stream is available.
	// DELETED: The stream is being deleted.
	// EXPIRED: The stream has expired.
	// Enumeration values:
	// CREATING
	// ACTIVE
	// DELETED
	// EXPIRED
	Status string `json:"status,omitempty"`
	// Unique identifier of the partition.
	PartitionId string `json:"partition_id,omitempty"`
	// Possible value range of the hash key used by the partition.
	HashRange string `json:"hash_range,omitempty"`
	// Sequence number range of the partition.
	SequenceNumberRange string `json:"sequence_number_range,omitempty"`
	// Parent partition.
	ParentPartitions string `json:"parent_partitions,omitempty"`
}

type UpdatePartitionCount struct {
	// Scaling execution timestamp, which is a 13-digit timestamp.
	CreateTimestamp *int64 `json:"create_timestamp,omitempty"`
	// Number of partitions before scaling.
	SrcPartitionCount *int32 `json:"src_partition_count,omitempty"`
	// Number of partitions after scaling.
	TargetPartitionCount *int32 `json:"target_partition_count,omitempty"`
	// Response code of the scaling operation.
	ResultCode *int32 `json:"result_code,omitempty"`
	// Response to the scaling operation.
	ResultMsg *int32 `json:"result_msg,omitempty"`
	// Specifies whether the scaling operation is automatic.
	// true: Auto scaling is enabled.
	// false: Manual scaling is enabled.
	AutoScale *bool `json:"auto_scale,omitempty"`
}
