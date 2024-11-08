package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
	Backup *Backup `json:"backup" required:"true"`
}

type Backup struct {
	// Specifies the instance ID, which can be obtained by calling the API for querying instances.
	// If you do not have an instance, you can call the API used for creating an instance.
	InstanceId string `json:"instance_id" required:"true"`
	// Specifies the manual backup name.
	// The value must be 4 to 64 characters in length and start with a letter (from A to Z or from a to z).
	// It is case-sensitive and can contain only letters, digits (from 0 to 9), hyphens (-), and underscores (_).
	Name string `json:"name" required:"true"`
	// Specifies the manual backup description.
	// The description must consist of a maximum of 256 characters and cannot
	// contain the following special characters: >!<"&'=
	Description string `json:"description,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Job, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/backups
	raw, err := client.Post(client.ServiceURL("backups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	return extractJob(err, raw)
}
