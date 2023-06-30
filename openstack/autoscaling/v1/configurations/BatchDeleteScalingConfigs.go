package configurations

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type BatchDeleteOpts struct {
	ScalingConfigurationId []string `json:"scaling_configuration_id"`
}

func BatchDeleteScalingConfigs(client *golangsdk.ServiceClient, opts BatchDeleteOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /autoscaling-api/v1/{project_id}/scaling_configurations
	_, err = client.Post(client.ServiceURL("scaling_configurations"), b, nil, nil)
	return
}
