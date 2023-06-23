package streams

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateStreamOpts struct {
	// Name of the stream whose partition quantity needs to be changed.
	StreamName string `json:"stream_name" required:"true"`
	// Period of time for which data is retained in the stream.
	// Value range: 24-72 Unit: hour
	// Default value: 24 If this parameter is left blank, the default value is used.
	// Maximum: 168
	// Default: 24
	DataDuration *int `json:"data_duration,omitempty"`
	// Source data type.
	// - BLOB: a collection of binary data stored as a single entity in a database management system.
	// Default value: BLOB.
	// Enumeration values:
	// BLOB
	DataType string `json:"data_type,omitempty"`
	// Source data structure that define JSON and CSV formats.
	// It is described in the syntax of the Avro schema.
	DataSchema string `json:"data_schema,omitempty"`
	// Specifies whether to enable auto-scaling.
	// - true: auto scaling is enabled.
	// - false:  auto-scaling is disabled. This function is disabled by default.
	// Default: false
	// Enumeration values:
	// true
	// false
	AutoScaleEnabled *bool `json:"auto_scale_enabled,omitempty"`
	// Minimum number of partitions for automatic scale-down when auto-scaling is enabled.
	// Minimum: 1
	AutoScaleMinPartitionCount *int64 `json:"auto_scale_min_partition_count,omitempty"`
	// Maximum number of partitions for automatic scale-up when auto-scaling is enabled.
	AutoScaleMaxPartitionCount *int64 `json:"auto_scale_max_partition_count,omitempty"`
}

func UpdateStream(client *golangsdk.ServiceClient, opts UpdateStreamOpts) error {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// PUT /v3/{project_id}/streams/{stream_name}
	_, err = client.Put(client.ServiceURL("streams", opts.StreamName), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
