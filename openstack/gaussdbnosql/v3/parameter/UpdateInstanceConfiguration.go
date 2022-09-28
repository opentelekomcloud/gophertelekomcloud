package parameter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func UpdateInstanceConfiguration(client *golangsdk.ServiceClient, opts UpdateInstanceNameOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/name
	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "name"), b, nil, nil)
	return
}
