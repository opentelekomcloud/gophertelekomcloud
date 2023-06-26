package checkpoints

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CommitCheckpointOpts struct {
	// Name of the app, which is the unique identifier of a user data consumption program.
	AppName string `json:"app_name"`
	// Type of the checkpoint.
	// LAST_READ: Only sequence numbers are recorded in databases.
	// Enumeration values:
	// LAST_READ
	CheckpointType string `json:"checkpoint_type"`
	// Name of the stream.
	StreamName string `json:"stream_name"`
	// Partition identifier of the stream. The value can be in either of the following formats:
	// shardId-0000000000
	// 0
	// For example, if a stream has three partitions, the partition identifiers are 0, 1, and 2, or shardId-0000000000, shardId-0000000001, and shardId-0000000002, respectively.
	PartitionId string `json:"partition_id"`
	// Sequence number to be submitted, which is used to record the consumption checkpoint of the stream.
	// Ensure that the sequence number is within the valid range.
	SequenceNumber string `json:"sequence_number"`
	// Metadata information of the consumer application.
	// The metadata information can contain a maximum of 1,000 characters.
	// Maximum: 1000
	Metadata string `json:"metadata,omitempty"`
}

func CommitCheckpoint(client *golangsdk.ServiceClient, opts CommitCheckpointOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/checkpoints
	_, err = client.Post(client.ServiceURL("checkpoints"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return err
}
