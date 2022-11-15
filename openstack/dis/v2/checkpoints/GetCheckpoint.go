package checkpoints

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetCheckpointOpts struct {
	// Name of the stream to which the checkpoint belongs.
	StreamName string `q:"stream_name"`
	// Identifier of the stream partition to which the checkpoint belongs.
	// The value can be in either of the following formats:
	//  - shardId-0000000000
	//  - 0
	// For example, if a stream has three partitions, the partition identifiers are 0, 1, and 2,
	// or shardId-0000000000, shardId-0000000001, and shardId-0000000002, respectively.
	PartitionId string `q:"partition_id"`
	// Name of the app associated with the checkpoint.
	AppName string `q:"app_name"`
	// Type of the checkpoint.
	// LAST_READ: Only sequence numbers are recorded in databases.
	// Enumeration values:
	// LAST_READ
	CheckpointType string `q:"checkpoint_type"`
}

func GetCheckpoint(client *golangsdk.ServiceClient, opts GetCheckpointOpts) (*GetCheckpointResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/checkpoints
	raw, err := client.Get(client.ServiceURL("checkpoints")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetCheckpointResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetCheckpointResponse struct {
	// Sequence number used to record the consumption checkpoint of the stream.
	SequenceNumber string `json:"sequence_number,omitempty"`
	// Metadata information of the consumer application.
	Metadata string `json:"metadata,omitempty"`
}
