package volumetypes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Name is the name of the volume type.
	Name string `json:"name,omitempty"`
	// Description is the description of the volume type.
	Description string `json:"description,omitempty"`
	// IsPublic controls whether the volume type is publicly visible.
	IsPublic *bool `json:"is_public,omitempty"`
}

// Update updates the volume type with the provided ID.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*VolumeType, error) {
	b, err := build.RequestBody(opts, "volume_type")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("types", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		VolumeType VolumeType `json:"volume_type"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.VolumeType, err
}
