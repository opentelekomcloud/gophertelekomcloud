package streams

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdatePartitionCountOpts struct {
	// Name of the stream whose partition quantity needs to be changed.
	// Maximum: 64
	StreamName string `json:"stream_name" required:"true"`
	// Number of the target partitions.
	// The value is an integer greater than 0.
	// If the value is greater than the number of current partitions, scaling-up is required.
	// If the value is less than the number of current partitions, scale-down is required.
	// Note: A maximum of five scale-up/down operations can be performed for each stream within one hour.
	// If a scale-up/down operation is successfully performed, you cannot perform one more scale-up/down operation within the next one hour.
	// Minimum: 0
	TargetPartitionCount int `json:"target_partition_count"`
}

func UpdatePartitionCount(client *golangsdk.ServiceClient, opts UpdatePartitionCountOpts) error {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// PUT /v2/{project_id}/streams/{stream_name}
	_, err = client.Put(client.ServiceURL("streams", opts.StreamName), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
