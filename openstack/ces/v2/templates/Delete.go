package templates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DeleteOpts contains the options for batch deleting custom alarm templates.
type DeleteOpts struct {
	// Specifies the list of alarm template IDs to delete.
	// A maximum of 100 templates can be deleted at a time.
	TemplateIds []string `json:"template_ids" required:"true"`
	// Specifies whether to delete the alarm rules associated with the alarm template.
	// true: The alarm rules are also deleted.
	// false: The alarm rules are not deleted.
	DeleteAssociateAlarm bool `json:"delete_associate_alarm"`
}

// DeleteResponse contains the response from the batch delete request.
type DeleteResponse struct {
	// Specifies the list of deleted alarm template IDs.
	TemplateIds []string `json:"template_ids"`
}

// Delete batch deletes custom alarm templates.
func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (*DeleteResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/alarm-templates/batch-delete
	raw, err := client.Post(client.ServiceURL("alarm-templates", "batch-delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
