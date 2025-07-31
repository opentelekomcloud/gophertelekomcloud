package resource_tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Delete is a method to delete tags in batch using given parameters.
func Delete(client *golangsdk.ServiceClient, opts BatchOpts) ([]FailedResource, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1.0/resource-tags/batch-delete
	raw, err := client.Post(client.ServiceURL("resource-tags", "batch-delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
	})
	if err != nil {
		return nil, err
	}

	var res []FailedResource
	err = extract.IntoSlicePtr(raw.Body, &res, "failed_resources")
	return res, err
}
