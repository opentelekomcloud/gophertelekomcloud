package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// The name of the volume type
	Name string `json:"name" required:"true"`
	// The volume type description
	Description string `json:"description,omitempty"`
	// the ID of the existing volume snapshot
	IsPublic *bool `json:"os-volume-type-access:is_public,omitempty"`
	// Extra spec key-value pairs defined by the user.
	ExtraSpecs map[string]string `json:"extra_specs"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*VolumeType, error) {
	b, err := golangsdk.BuildRequestBody(opts, "volume_type")
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
