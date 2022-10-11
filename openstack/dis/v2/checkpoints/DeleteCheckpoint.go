package checkpoints

type DeleteCheckpointRequest struct {
	// Name of the stream to which the checkpoint belongs.
	StreamName string `json:"stream_name"`
	// Name of the application associated with the checkpoint.
	// Minimum: 1
	// Maximum: 50
	AppName string `json:"app_name"`
	// Type of the checkpoint. LAST_READ: Only sequence numbers are recorded in databases.
	// Enumeration values:
	// LAST_READ
	CheckpointType DeleteCheckpointRequestCheckpointType `json:"checkpoint_type"`
	// Identifier of the stream partition to which the checkpoint belongs. The value can be in either of the following formats:
	// shardId-0000000000
	// 0
	// For example, if a stream has three partitions, the partition identifiers are 0, 1, and 2, and shardId-0000000000, shardId-0000000001, shardId-0000000002, respectively.
	PartitionId *string `json:"partition_id,omitempty"`
}
type DeleteCheckpointRequestCheckpointType struct {
	value string
}

// DELETE /v2/{project_id}/checkpoints

type DeleteCheckpointResponse struct {
	HttpStatusCode int `json:"-"`
}
