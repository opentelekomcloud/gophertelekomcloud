package volumetypes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves the volume type with the provided ID.
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
