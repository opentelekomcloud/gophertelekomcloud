package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*VolumeType, error) {
	raw, err := client.Get(client.ServiceURL("types", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		VolumeType VolumeType `json:"volume_type"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.VolumeType, err
}
