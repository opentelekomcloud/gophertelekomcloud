package postgres_extensions

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// This function is used to create an extension for a specified database.
func Create(client *golangsdk.ServiceClient, instanceId string, opts PostgresExtensionOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/extensions
	_, err = client.Post(client.ServiceURL("instances", instanceId, "extensions"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
