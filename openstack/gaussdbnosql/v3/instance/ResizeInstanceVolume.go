package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ResizeInstanceVolumeOpts struct {
	InstanceId string
	// The requested disk capacity. The value must be an integer greater than the current disk capacity.
	// The maximum disk capacity depends on the engine type and specifications.
	Size int32 `json:"size"`
}

func ResizeInstanceVolume(client *golangsdk.ServiceClient, opts ResizeInstanceVolumeOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/extend-volume
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "extend-volume"), b, nil, nil)
	return extraJob(err, raw)
}
