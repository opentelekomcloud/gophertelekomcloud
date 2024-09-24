package tag

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type BatchTagActionOpts struct {
	InstanceId string
	// Operation identifier. Valid value:
	// create: indicates to add tags.
	// delete: indicates to delete tags.
	Action string `json:"action"`
	// Tag list.
	Tags []BatchTagActionTagOption `json:"tags"`
}

type BatchTagActionTagOption struct {
	// Tag key. It contains a maximum of 36 Unicode characters. It cannot be null or an empty string or contain spaces.
	// Before verifying and using key, spaces are automatically filtered out. Character set: 0-9, A-Z, a-z, "_", and "-".
	Key string `json:"key"`
	// Tag value. It contains a maximum of 43 Unicode characters, can be an empty string, and cannot contain spaces.
	// Before verifying or using value, spaces are automatically filtered out.
	// Character set: 0-9, A-Z, a-z, "_", and "-".
	// If action is set to create, this parameter is mandatory.
	// If action is set to delete, this parameter is optional.
	// NOTE If value is specified, tags are deleted by key and value.
	// If value is not specified, tags are deleted by key.
	Value string `json:"value,omitempty"`
}

func BatchTagAction(client *golangsdk.ServiceClient, opts BatchTagActionOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/tags/action
	_, err = client.Post(client.ServiceURL("instances", opts.InstanceId, "tags", "action"), b, nil, nil)
	return
}
