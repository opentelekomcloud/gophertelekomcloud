package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// The size of the volume, in GB
	Size int `json:"size" required:"true"`
	// The availability zone
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// ConsistencyGroupID is the ID of a consistency group
	ConsistencyGroupID string `json:"consistencygroup_id,omitempty"`
	// The volume description
	Description string `json:"description,omitempty"`
	// One or more metadata key and value pairs to associate with the volume
	Metadata map[string]string `json:"metadata,omitempty"`
	// The volume name
	Name string `json:"name,omitempty"`
	// the ID of the existing volume snapshot
	SnapshotID string `json:"snapshot_id,omitempty"`
	// SourceReplica is a UUID of an existing volume to replicate with
	SourceReplica string `json:"source_replica,omitempty"`
	// the ID of the existing volume
	SourceVolID string `json:"source_volid,omitempty"`
	// The ID of the image from which you want to create the volume.
	// Required to create a bootable volume.
	ImageID string `json:"imageRef,omitempty"`
	// The associated volume type
	VolumeType string `json:"volume_type,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Volume, error) {
	b, err := golangsdk.BuildRequestBody(opts, "volume")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("volumes"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res Volume
	err = extract.Into(raw.Body, &res)
	return &res, err
}
