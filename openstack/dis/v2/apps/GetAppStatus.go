package apps

type GetAppStatusOpts struct {
	// Name of the app to be queried.
	AppName string `json:"app_name"`
	// Name of the stream to be queried.
	// Maximum: 60
	StreamName string `json:"stream_name"`
	// Max. number of partitions to list in a single API call.
	// The minimum value is 1 and the maximum value is 1,000.
	// The default value is 100.
	// Minimum: 1
	// Maximum: 1000
	// Default: 100
	Limit *int32 `json:"limit,omitempty"`
	// Name of the partition to start the partition list with.
	// The returned partition list does not contain this partition.
	StartPartitionId string `json:"start_partion_id,omitempty"`
	// Type of the checkpoint.
	// - LAST_READ: Only sequence numbers are recorded in databases.
	// Enumeration values:
	// LAST_READ
	CheckpointType string `json:"checkpoint_type"`
}

// GET /v2/{project_id}/apps/{app_name}/streams/{stream_name}

type GetAppStatusResponse struct {
	// Name of the app.
	AppName string `json:"app_name,omitempty"`
	// Unique identifier of the app.
	AppId string `json:"app_id,omitempty"`
	// Time when the app is created, in milliseconds.
	CreateTime *int64 `json:"create_time,omitempty"`
	// List of associated streams.
	CommitCheckPointStreamNames []string `json:"commit_check_point_stream_names,omitempty"`
}
