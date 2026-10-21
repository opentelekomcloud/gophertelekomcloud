package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteTagOpts struct {
	// List of service resource tags to delete.
	Tags []CreateTags `json:"tags,omitempty"`
}

// This function is used to delete a tag
func DeleteTag(client *golangsdk.ServiceClient, firewallId string, opts DeleteTagOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// DELETE /v2/{project_id}/cfw-cfw/{fw_instance_id}/tags/delete
	_, err = client.DeleteWithBody(client.ServiceURL("cfw-cfw", firewallId, "tags", "delete"), b, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return err
}
