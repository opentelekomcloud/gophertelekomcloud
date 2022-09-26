package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type MysqlUpdateInstanceNameOpts struct {
	// Instance ID, which is compliant with the UUID format.
	InstanceId string
	// Instance name
	// Instances of the same type can have same names under the same tenant.
	// The name consists of 4 to 64 characters and starts with a letter.
	// It is case-sensitive and can contain only letters, digits, hyphens (-), and underscores (_).
	Name string `json:"name"`
}

func UpdateGaussMySqlInstanceName(client *golangsdk.ServiceClient, opts MysqlUpdateInstanceNameOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// PUT https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}/name
	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "name"), b, nil, nil)
	return extraJob(err, raw)
}
