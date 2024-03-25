package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateInstanceNameOpts struct {
	InstanceId string
	// New instance name
	// Value range:
	// The value must be 4 to 64 characters in length and start with a letter (from A to Z or from a to z).
	// It is case-sensitive and can contain only letters, digits (from 0 to 9), hyphens (-), and underscores (_).
	Name string `json:"name"`
}

func UpdateInstanceName(client *golangsdk.ServiceClient, opts UpdateInstanceNameOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/name
	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "name"), b, nil, nil)
	return
}
