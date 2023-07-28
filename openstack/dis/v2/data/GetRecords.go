package data

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"net/url"
)

type GetRecordsOpts struct {
	// Data cursor, which needs to be obtained through the API for obtaining data cursors.
	// Value range: a string of 1 to 512 characters
	// Note:
	// The validity period of a data cursor is 5 minutes.
	PartitionCursor string `q:"partition-cursor"`
	// Maximum number of bytes that can be obtained for each request.
	// Note:
	// If the value is less than the size of a single record in the partition, the record cannot be obtained.
	MaxFetchBytes *int `q:"max_fetch_bytes,omitempty"`
}

func GetRecords(client *golangsdk.ServiceClient, opts GetRecordsOpts) (*GetRecordsResponse, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/records
	raw, err := client.Get(client.ServiceURL("records")+q.String(), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res GetRecordsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetRecordsResponse struct {
	Records []Record `json:"records,omitempty"`
	// Next iterator.
	// Note:
	// The validity period of a data cursor is 5 minutes.
	NextPartitionCursor string `json:"next_partition_cursor,omitempty"`
}

type Record struct {
	// Partition key set when data is being uploaded.
	// Note:
	// If the partition_key parameter is passed when data is uploaded, this parameter will be returned when data is downloaded.
	// If partition_id instead of partition_key is passed when data is uploaded, no partition_key is returned.
	PartitionKey string `json:"partition_key,omitempty"`
	// Sequence number of the data record.
	SequenceNumber string `json:"sequence_number,omitempty"`
	// Downloaded data.
	// The downloaded data is the serialized binary data (Base64-encoded character string).
	// For example, the data returned by the data download API is "ZGF0YQ==", which is "data" after Base64 decoding.
	Data string `json:"data,omitempty"`
	// Timestamp when the record is written to DIS.
	CreatedAt *int64 `json:"timestamp,omitempty"`
	// Timestamp data type.
	// CreateTime: creation time.
	// Default: CreateTime
	TimestampType string `json:"timestamp_type,omitempty"`
}
