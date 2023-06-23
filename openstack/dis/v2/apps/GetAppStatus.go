package apps

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetAppStatusOpts struct {
	// Name of the app to be queried.
	AppName string
	// Name of the stream to be queried.
	// Maximum: 60
	StreamName string
	// Max. number of partitions to list in a single API call.
	// The minimum value is 1 and the maximum value is 1,000.
	// The default value is 100.
	// Minimum: 1
	// Maximum: 1000
	// Default: 100
	Limit *int32 `q:"limit,omitempty"`
	// Name of the partition to start the partition list with.
	// The returned partition list does not contain this partition.
	StartPartitionId string `q:"start_partition_id,omitempty"`
	// Type of the checkpoint.
	// - LAST_READ: Only sequence numbers are recorded in databases.
	// Enumeration values:
	// LAST_READ
	CheckpointType string `q:"checkpoint_type"`
}

func GetAppStatus(client *golangsdk.ServiceClient, opts GetAppStatusOpts) (*GetAppStatusResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/apps/{app_name}/streams/{stream_name}
	raw, err := client.Get(client.ServiceURL("apps", opts.AppName, "streams", opts.StreamName)+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetAppStatusResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetAppStatusResponse struct {
	// Specify whether there are more matching DIS streams to list. Possible values:
	// true: yes
	// false: no
	// Default: false.
	HasMore bool `json:"has_more"`
	// Stream Name
	StreamName string `json:"stream_name"`
	// App name
	AppName string `json:"app_name"`
	// Partition consuming state list
	PartitionConsumingStates []struct {
		// Partition Id
		PartitionId string `json:"partition_id"`
		// Partition Sequence Number
		SequenceNumber string `json:"sequence_number"`
		// Partition data latest offset
		LatestOffset int `json:"latest_offset"`
		// Partition data earliest offset
		EarliestOffset int `json:"earliest_offset"`
		// Type of the checkpoint.
		// LAST_READ: Only sequence numbers are recorded in databases.
		// Enumeration values:
		// LAST_READ
		CheckpointType string `json:"checkpoint_type"`
		// Partition Status:
		// CREATING
		// ACTIVE
		// DELETED
		// EXPIRED
		Status string `json:"status"`
	} `json:"partition_consuming_states"`
}
