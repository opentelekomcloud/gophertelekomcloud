package data

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetCursorOpts struct {
	// Name of the stream.
	StreamName string `q:"stream-name"`
	// Partition ID of the stream. The value can be in either of the following formats:
	// shardId-00000000000
	// For example, if a stream has three partitions, the partition identifiers are 0, 1, and 2,
	// or shardId-0000000000, shardId-0000000001, and shardId-0000000002, respectively.
	PartitionId string `q:"partition-id"`
	// Cursor type.
	// AT_SEQUENCE_NUMBER: Data is read from the position denoted by a specific sequence number
	// (that is defined by starting-sequence-number). This is the default cursor type.
	// AFTER_SEQUENCE_NUMBER: Data is read right after the position denoted by a specific sequence number
	// (that is defined by starting-sequence-number).
	// TRIM_HORIZON: Data is read from the earliest data record in the partition.
	// For example, a tenant uses a DIS stream to upload three pieces of data A1, A2, and A3. N days later,
	// A1 has expired and A2 and A3 are still in the validity period.
	// In this case, if the tenant uses TRIM_HORIZON to download the data, the system downloads data from A2.
	// LATEST: Data is read from the latest record in the partition.
	// This setting ensures that you always read the latest record in the partition.
	// AT_TIMESTAMP: Data is read from the position denoted by a specific timestamp.
	// Enumeration values:
	// AT_SEQUENCE_NUMBER
	// AFTER_SEQUENCE_NUMBER
	// TRIM_HORIZON
	// LATEST
	// AT_TIMESTAMP
	CursorType string `q:"cursor-type,omitempty"`
	// Serial number. A sequence number is a unique identifier for each record.
	// DIS automatically allocates a sequence number when the data producer calls the PutRecords operation to add data to the DIS stream.
	// SN of the same partition key usually changes with time.
	// A longer interval between PutRecords requests results in a larger sequence number.
	// The sequence number is closely related to cursor types AT_SEQUENCE_NUMBER and AFTER_SEQUENCE_NUMBER.
	// The two parameters determine the position of the data to be read.
	// Value range: 0 to 9,223,372,036,854,775,807
	StartingSequenceNumber string `q:"starting-sequence-number,omitempty"`
	// Timestamp when the data record starts to be read, which is closely related to cursor type AT_TIMESTAMP.
	// The two parameters determine the position of the data to be read.
	// Note:
	// The timestamp is accurate to milliseconds.
	Timestamp *int64 `q:"timestamp,omitempty"`
	// Unique ID of the stream. This parameter is mandatory for obtaining the iterator of an authorized stream.
	StreamId string `q:"stream-id,omitempty"`
}

func GetCursor(client *golangsdk.ServiceClient, opts GetCursorOpts) (*GetCursorResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("cursors").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/cursors
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetCursorResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetCursorResponse struct {
	// Data cursor. Value range: a string of 1 to 512 characters
	// Note:
	// The validity period of a data cursor is 5 minutes.
	// Minimum: 1
	// Maximum: 512
	PartitionCursor string `json:"partition_cursor,omitempty"`
}
