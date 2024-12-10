package gateway

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type FeatureOpts struct {
	GatewayID string `json:"-"`
	// Feature name.
	Name string `json:"name" required:"true"`
	// Indicates whether to enable the feature.
	Enable *bool `json:"enable" required:"true"`
	// Parameter configuration.
	Config string `json:"config,omitempty"`
}

func ConfigureFeature(client *golangsdk.ServiceClient, opts FeatureOpts) (*FeatureResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "features"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res FeatureResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type FeatureResp struct {
	// Feature ID.
	ID string `json:"id"`
	// Feature name.
	Name string `json:"name"`
	// Indicates whether to enable the feature.
	Enabled bool `json:"enable"`
	// Parameter configuration.
	Config string `json:"config"`
	// Gateway ID.
	GatewayID string `json:"instance_id"`
	// Feature update time.
	UpdatedAt string `json:"update_time"`
}
