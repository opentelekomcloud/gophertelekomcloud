package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateInstanceNameOpts struct {
	InstanceId string `json:"-"`
	// Specifies the DB instance name.
	// DB instances of the same type can have same names under the same tenant.
	// The parameter must be 4 to 64 characters long, start with a letter, and contain only letters (case-sensitive), digits, hyphens (-), and underscores (_).
	Name string `json:"name"`
}

func UpdateInstanceName(client *golangsdk.ServiceClient, opts UpdateInstanceNameOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/name
	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "name"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
