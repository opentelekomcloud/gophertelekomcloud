package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsPublic    *bool  `json:"is_public,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*VolumeType, error) {
	b, err := golangsdk.BuildRequestBody(opts, "volume_type")
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
