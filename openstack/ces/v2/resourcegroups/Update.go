package resourcegroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts contains the options for modifying a resource group.
type UpdateOpts struct {
	// Specifies the resource group name.
	// It can contain 1 to 128 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	GroupName string `json:"group_name" required:"true"`
	// Specifies the tags for dynamic resource matching.
	// This parameter is used when the resource group type is TAG.
	// A maximum of 10 tags can be specified.
	Tags []ResourceGroupTag `json:"tags,omitempty"`
}

// Update modifies a resource group.
func Update(client *golangsdk.ServiceClient, groupId string, opts UpdateOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// PUT /v2/{project_id}/resource-groups/{group_id}
	_, err = client.Put(client.ServiceURL("resource-groups", groupId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
