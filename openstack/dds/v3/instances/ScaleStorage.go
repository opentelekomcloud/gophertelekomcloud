package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ScaleStorageOpt struct {
	// Specifies the role ID.
	//
	// For a cluster instance, this parameter is set to the ID of the shard group.
	// This parameter is not transferred for replica set and single node instances.
	GroupId string `json:"group_id,omitempty"`
	// Specifies the requested disk capacity. The value must be an integer multiple of 10 and greater than the current storage space.
	//
	// In a cluster instance, this parameter indicates the storage space of shard nodes. The value range is from 10 GB to 2000 GB.
	// In a replica set instance, this parameter indicates the disk capacity of the DB instance to be expanded. The value range is from 10 GB to 2000 GB.
	// In a single node instance, this parameter indicates the disk capacity of the DB instance to be expanded. The value range is from 10 GB to 1000 GB.
	Size string `json:"size" required:"true"`
	// Specifies the instance
	InstanceId string `json:"-"`
}

func ScaleStorage(client *golangsdk.ServiceClient, opts ScaleStorageOpt) (*string, error) {
	b, err := build.RequestBody(opts, "volume")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/enlarge-volume
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "enlarge-volume"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return extractJob(err, raw)
}
