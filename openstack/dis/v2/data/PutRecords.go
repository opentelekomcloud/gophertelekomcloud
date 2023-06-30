package data

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type PutRecordsOpts struct {
	// Name of the stream.
	// Maximum: 64
	StreamName string `json:"stream_name" required:"true"`
	// Unique ID of the stream.
	// If no stream is found by stream_name and stream_id is not empty, stream_id is used to search for the stream.
	// Note:
	// This parameter is mandatory when data is uploaded to the authorized stream.
	StreamId string `json:"stream_id,omitempty"`
	// List of records to be uploaded.
	Records []PutRecordsRequestEntry `json:"records" required:"true"`
}

func PutRecords(client *golangsdk.ServiceClient, opts PutRecordsOpts) (*PutRecordsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/records
	raw, err := client.Post(client.ServiceURL("records"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res PutRecordsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type PutRecordsRequestEntry struct {
	// Data to be uploaded. The uploaded data is the serialized binary data (character string encoded using Base64).
	// For example, if the character string data needs to be uploaded, the character string after Base64 encoding is ZGF0YQ==.
	Data string `json:"data"`
	// Hash value of the data to be written to the partition. The hash value overwrites the hash value of partition_key. Value range: 0–long.max
	ExplicitHashKey string `json:"explicit_hash_key,omitempty"`
	// Partition ID of the stream. The value can be in either of the following formats:
	// shardId-0000000000
	// 0
	// For example, if a stream has three partitions, the partition identifiers are 0, 1, and 2, or shardId-0000000000, shardId-0000000001, and shardId-0000000002, respectively.
	PartitionId string `json:"partition_id,omitempty"`
	// Partition to which data is written to. Note:
	// If the partition_id parameter is transferred, the partition_id parameter is used preferentially.
	// If partition_id is not transferred, partition_key is used.
	PartitionKey string `json:"partition_key,omitempty"`
}

type PutRecordsResponse struct {
	// Number of data records that fail to be uploaded.
	FailedRecordCount *int `json:"failed_record_count,omitempty"`

	Records []PutRecordsResultEntry `json:"records,omitempty"`
}

type PutRecordsResultEntry struct {
	// ID of the partition to which data is uploaded.
	PartitionId string `json:"partition_id,omitempty"`
	// Sequence number of the data to be uploaded.
	// A sequence number is a unique identifier for each record.
	// DIS automatically allocates a sequence number the data producer calls the PutRecords operation to add data to the DIS stream.
	// Sequence number of the same partition key usually changes with time. A longer interval between PutRecords requests results in a larger sequence number.
	SequenceNumber string `json:"sequence_number,omitempty"`
	// Error code.
	ErrorCode string `json:"error_code,omitempty"`
	// Error message.
	ErrorMessage string `json:"error_message,omitempty"`
}
