package resourcegroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains the options for creating a resource group.
type CreateOpts struct {
	// Specifies the resource group name.
	// It can contain 1 to 128 characters and must start with a letter.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	GroupName string `json:"group_name" required:"true"`
	// Specifies the enterprise project ID.
	// The value can be a UUID or "0".
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// Specifies how resources are added to the resource group.
	// Possible values: EPS, TAG, Manual
	// Default: Manual
	Type string `json:"type,omitempty"`
	// Specifies the tags for dynamic resource matching.
	// This parameter is mandatory when type is set to TAG.
	// A maximum of 10 tags can be specified.
	Tags []ResourceGroupTag `json:"tags,omitempty"`
	// Specifies the enterprise project IDs for the EPS type.
	// This parameter is mandatory when type is set to EPS.
	// A maximum of 10 enterprise project IDs can be specified.
	AssociationEpIds []string `json:"association_ep_ids,omitempty"`
}

// ResourceGroupTag represents a tag for dynamic resource matching.
type ResourceGroupTag struct {
	// Specifies the tag key.
	// It can contain 1 to 36 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	Key string `json:"key" required:"true"`
	// Specifies the tag value.
	// It can contain 0 to 43 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	Value string `json:"value,omitempty"`
}

// Create creates a new resource group.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /v2/{project_id}/resource-groups
	raw, err := client.Post(client.ServiceURL("resource-groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		GroupId string `json:"group_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.GroupId, err
}
