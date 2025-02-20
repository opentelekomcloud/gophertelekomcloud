package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateTagOpts struct {
	// List of service resource tags. After tags are added to firewall resources, you can query resources and combine CDRs by key and value.
	Tags []CreateTags `json:"tags,omitempty"`
}

// Create function is used to create a tag
func CreateTag(client *golangsdk.ServiceClient, firewallId string, opts CreateTagOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/cfw-cfw/{fw_instance_id}/tags/create
	_, err = client.Post(client.ServiceURL("cfw-cfw", firewallId, "tags", "create"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return err
}
