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

	var res VolumeType
	err = extract.IntoStructPtr(raw.Body, &res, "volume_type")
	return &res, err
}
