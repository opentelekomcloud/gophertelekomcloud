package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient) ([]VolumeType, error) {
	raw, err := client.Get(client.ServiceURL("types"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []VolumeType
	err = extract.IntoSlicePtr(raw.Body, &res, "volume_types")
	return res, err
}

type VolumeType struct {
	// user-defined metadata
	ExtraSpecs map[string]any `json:"extra_specs"`
	// unique identifier
	ID string `json:"id"`
	// display name
	Name string `json:"name"`
}
