package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// // This function is used to update log configurations.
func UpdateLogConfig(client *golangsdk.ServiceClient, opts LogConfigOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/cfw/logs/configuration
	raw, err := client.Put(client.ServiceURL("cfw", "logs", "configuration"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res DataResponse
	return &res.Data, extract.Into(raw.Body, &res)
}
