package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// // This function is used to update log configurations.
func UpdateLogConfig(client *golangsdk.ServiceClient, opts LogConfigOpts) (*string, error) {
	// PUT /v1/{project_id}/cfw/logs/configuration
	url, err := golangsdk.NewURLBuilder().WithEndpoints("cfw", "logs", "configuration").WithQueryParams(&QueryParameters{
		FwInstanceID: opts.FWInstanceID,
	}).Build()
	if err != nil {
		return nil, err
	}
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res DataResponse
	return &res.Data, extract.Into(raw.Body, &res)
}
