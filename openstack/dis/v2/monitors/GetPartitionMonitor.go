package monitors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetPartitionMonitorOpts struct {
	// Name of the stream to be queried.
	// Maximum: 60
	StreamName string
	// Partition No.
	// The value can be in either of the following formats:
	//  - shardId-0000000000
	//  - 0
	// For example, if a stream has three partitions, the partition identifiers are 0, 1, and 2, or
	// shardId-0000000000,
	// shardId-0000000001,
	// and shardId-0000000002,
	// respectively.
	PartitionId string
	// Partition monitoring metric.
	// (Either label or label_list must be specified.
	// If both label_list and label are specified, label_list prevails.)
	// - total_put_bytes_per_stream: total input traffic (byte)
	// - total_get_bytes_per_stream: total output traffic (byte)
	// - total_put_records_per_stream: total number of input records
	// - total_get_records_per_stream: total number of output records
	// Enumeration values:
	//  total_put_bytes_per_stream
	//  total_get_bytes_per_stream
	//  total_put_records_per_stream
	//  total_get_records_per_stream
	Label string `q:"label,omitempty"`
	// List of labels separated by commas (,) to query multiple labels in batches.
	// (Either label or label_list must be specified.
	// If both label_list and label exist, label_list prevails.)
	LabelList string `q:"label_list,omitempty"`
	// Monitoring start time, which is a 10-digit timestamp.
	StartTime int64 `q:"start_time"`
	// Monitoring end time, which is a 10-digit timestamp.
	EndTime int64 `q:"end_time"`
}

func GetPartitionMonitor(client *golangsdk.ServiceClient, opts GetPartitionMonitorOpts) (*GetPartitionMonitorResponse, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/streams/{stream_name}/partitions/{partition_id}/metrics
	raw, err := client.Get(client.ServiceURL("streams", opts.StreamName, "partitions", opts.PartitionId, "metrics")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetPartitionMonitorResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetPartitionMonitorResponse struct {
	// Data object.
	Metrics Metrics `json:"metrics,omitempty"`
}
