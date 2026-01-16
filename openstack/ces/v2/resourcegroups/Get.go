package resourcegroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ResourceGroupDetail represents the detailed information of a resource group.
type ResourceGroupDetail struct {
	// Specifies the resource group name.
	GroupName string `json:"group_name"`
	// Specifies the resource group ID.
	GroupId string `json:"group_id"`
	// Specifies the time when the resource group was created.
	// The value is in UTC format: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	CreateTime string `json:"create_time"`
	// Specifies the enterprise project ID.
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// Specifies how resources are added to the resource group.
	// Possible values: EPS, TAG, Manual
	Type string `json:"type"`
	// Specifies the enterprise project IDs for the EPS type.
	AssociationEpIds []string `json:"association_ep_ids"`
	// Specifies the tags for dynamic resource matching.
	Tags []ResourceGroupTag `json:"tags"`
}

// Get retrieves details of a resource group.
func Get(client *golangsdk.ServiceClient, groupId string) (*ResourceGroupDetail, error) {
	// GET /v2/{project_id}/resource-groups/{group_id}
	raw, err := client.Get(client.ServiceURL("resource-groups", groupId), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ResourceGroupDetail
	err = extract.Into(raw.Body, &res)
	return &res, err
}
