package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ResizeInstanceOpts struct {
	InstanceId string
	Resize     ResizeInstanceOption `json:"resize"`
}

type ResizeInstanceOption struct {
	// Resource specification code of the new specification
	TargetSpecCode string `json:"target_spec_code"`
}

func ResizeInstance(client *golangsdk.ServiceClient, opts ResizeInstanceOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/resize
	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "resize"), b, nil, nil)
	return extraJob(err, raw)
}
