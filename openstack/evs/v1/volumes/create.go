package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Size             int               `json:"size" required:"true"`
	AvailabilityZone string            `json:"availability_zone,omitempty"`
	Description      string            `json:"display_description,omitempty"`
	Metadata         map[string]string `json:"metadata,omitempty"`
	Name             string            `json:"display_name,omitempty"`
	SnapshotID       string            `json:"snapshot_id,omitempty"`
	SourceVolID      string            `json:"source_volid,omitempty"`
	ImageID          string            `json:"imageRef,omitempty"`
	VolumeType       string            `json:"volume_type,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Volume, error) {
	b, err := build.RequestBodyMap(opts, "volume")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("volumes"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	var res Volume
	err = extract.IntoStructPtr(raw.Body, &res, "volume")
	return &res, err
}
