package volumes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Size               int               `json:"size" required:"true"`
	AvailabilityZone   string            `json:"availability_zone,omitempty"`
	ConsistencyGroupID string            `json:"consistencygroup_id,omitempty"`
	Description        string            `json:"description,omitempty"`
	Metadata           map[string]string `json:"metadata,omitempty"`
	Name               string            `json:"name,omitempty"`
	SnapshotID         string            `json:"snapshot_id,omitempty"`
	SourceReplica      string            `json:"source_replica,omitempty"`
	SourceVolID        string            `json:"source_volid,omitempty"`
	ImageID            string            `json:"imageRef,omitempty"`
	VolumeType         string            `json:"volume_type,omitempty"`
	Multiattach        bool              `json:"multiattach,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Volume, error) {
	b, err := build.RequestBody(opts, "volume")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("volumes"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		Volume Volume `json:"volume"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Volume, err
}
