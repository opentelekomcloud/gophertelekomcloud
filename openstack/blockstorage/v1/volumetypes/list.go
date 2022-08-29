package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, client.ServiceURL("types"), func(r pagination.PageResult) pagination.Page {
		return VolumeTypePage{pagination.SinglePageBase(r)}
	})
}

type VolumeType struct {
	// user-defined metadata
	ExtraSpecs map[string]interface{} `json:"extra_specs"`
	// unique identifier
	ID string `json:"id"`
	// display name
	Name string `json:"name"`
}

type VolumeTypePage struct {
	pagination.SinglePageBase
}

func (r VolumeTypePage) IsEmpty() (bool, error) {
	volumeTypes, err := ExtractVolumeTypes(r)
	return len(volumeTypes) == 0, err
}

func ExtractVolumeTypes(r pagination.Page) ([]VolumeType, error) {
	var res struct {
		VolumeTypes []VolumeType `json:"volume_types"`
	}

	err := extract.Into(r.(VolumeTypePage).Result.Body, &res)
	return res.VolumeTypes, err
}
