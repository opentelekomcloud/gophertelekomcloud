package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Specifies the AS configuration name.
	// Supports fuzzy search.
	Name string `q:"scaling_configuration_name"`
	// Specifies the image ID. It is same as imageRef.
	ImageID string `q:"image_id"`
	// Specifies the start line number. The default value is 0. The minimum parameter value is 0.
	StartNumber int `q:"start_number"`
	// Specifies the number of query records. The default value is 20. The value range is 0 to 100.
	Limit int `q:"limit"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListScalingConfigsResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("scaling_configuration")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListScalingConfigsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListScalingConfigsResponse struct {
	TotalNumber           int32           `json:"total_number,omitempty"`
	StartNumber           int32           `json:"start_number,omitempty"`
	Limit                 int32           `json:"limit,omitempty"`
	ScalingConfigurations []Configuration `json:"scaling_configurations,omitempty"`
}
