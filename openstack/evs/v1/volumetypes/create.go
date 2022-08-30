package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// See VolumeType.
	ExtraSpecs map[string]interface{} `json:"extra_specs,omitempty"`
	// See VolumeType.
	Name string `json:"name,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*VolumeType, error) {
	b, err := golangsdk.BuildRequestBody(opts, "volume_type")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("types"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
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
