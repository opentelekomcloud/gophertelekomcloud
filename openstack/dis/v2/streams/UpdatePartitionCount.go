package streams

type UpdatePartitionCountOpts struct {
	// Name of the stream whose partition quantity needs to be changed.
	// Maximum: 64
	StreamName string `json:"stream_name"`

	Body UpdatePartitionCountRequestBody `json:"body,omitempty"`
}

type UpdatePartitionCountRequestBody struct {
	// Name of the stream whose partition quantity needs to be changed.
	// Maximum: 64
	StreamName string `json:"stream_name"`
	// Number of the target partitions.
	// The value is an integer greater than 0.
	// If the value is greater than the number of current partitions, scaling-up is required.
	// If the value is less than the number of current partitions, scale-down is required.
	// Note: A maximum of five scale-up/down operations can be performed for each stream within one hour.
	// If a scale-up/down operation is successfully performed, you cannot perform one more scale-up/down operation within the next one hour.
	// Minimum: 0
	TargetPartitionCount int32 `json:"target_partition_count"`
}

// PUT /v2/{project_id}/streams/{stream_name}

type UpdatePartitionCountResponse struct {
}
