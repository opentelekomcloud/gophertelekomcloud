package checkpoints

type GetCheckpointRequest struct {
	// Name of the stream to which the checkpoint belongs.
	StreamName string `json:"stream_name"`
	// Identifier of the stream partition to which the checkpoint belongs. The value can be in either of the following formats:
	// shardId-0000000000
	// 0
	// For example, if a stream has three partitions, the partition identifiers are 0, 1, and 2, or shardId-0000000000, shardId-0000000001, and shardId-0000000002, respectively.
	PartitionId string `json:"partition_id"`
	// Name of the app associated with the checkpoint.
	AppName string `json:"app_name"`
	// Type of the checkpoint.
	// LAST_READ: Only sequence numbers are recorded in databases.
	// Enumeration values:
	// LAST_READ
	CheckpointType GetCheckpointRequestCheckpointType `json:"checkpoint_type"`
}
type GetCheckpointRequestCheckpointType struct {
	value string
}

// GET /v2/{project_id}/checkpoints

type GetCheckpointResponse struct {
	// Sequence number used to record the consumption checkpoint of the stream.
	SequenceNumber *string `json:"sequence_number,omitempty"`
	// Metadata information of the consumer application.
	Metadata       *string `json:"metadata,omitempty"`
	HttpStatusCode int     `json:"-"`
}
