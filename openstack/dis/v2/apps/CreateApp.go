package apps

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateAppOpts struct {
	// Unique identifier of the consumer application to be created.
	// The application name contains 1 to 200 characters, including letters, digits, underscores (_), and hyphens (-).
	// Minimum: 1
	// Maximum: 200
	AppName string `json:"app_name"`
}

func CreateApp(client *golangsdk.ServiceClient, opts CreateAppOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/apps
	_, err = client.Post(client.ServiceURL("apps"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return err
}
