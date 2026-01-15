package alarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DeleteOpts contains the options for deleting alarm rules in batches.
type DeleteOpts struct {
	// Specifies the IDs of the alarm rules to be deleted in batches.
	// A maximum of 10 alarm rules can be deleted at a time.
	AlarmIds []string `json:"alarm_ids" required:"true"`
}

// DeleteResponse contains the response from the Delete request.
type DeleteResponse struct {
	// Specifies the IDs of the alarm rules that are deleted.
	AlarmIds []string `json:"alarm_ids"`
}

// Delete deletes alarm rules in batches.
func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (*DeleteResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/alarms/batch-delete
	raw, err := client.Post(client.ServiceURL("alarms", "batch-delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
