package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ScalingOpts struct {
	SwitchOption     bool `json:"switch_option" required:"true"`
	LimitSize        int  `json:"limit_size,omitempty"`
	TriggerThreshold int  `json:"trigger_threshold,omitempty"`
}

func ManageAutoScaling(client *golangsdk.ServiceClient, id string, opts ScalingOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/disk-auto-expansion
	_, err = client.Put(client.ServiceURL("instances", id, "disk-auto-expansion"), b, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
