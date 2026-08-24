package volumetypes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Name is the name of the volume type.
	Name string `json:"name" required:"true"`
	// Description is the description of the volume type.
	Description string `json:"description,omitempty"`
	// IsPublic controls whether the volume type is publicly visible.
	IsPublic *bool `json:"os-volume-type-access:is_public,omitempty"`
	// ExtraSpecs contains additional specifications for the volume type.
	ExtraSpecs map[string]string `json:"extra_specs,omitempty"`
}

// Create creates a volume type.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*VolumeType, error) {
	b, err := build.RequestBody(opts, "volume_type")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("types"), b, nil, &golangsdk.RequestOpts{
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
