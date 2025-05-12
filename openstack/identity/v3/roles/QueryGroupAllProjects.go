package roles

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func QueryGroupAllProjects(client *golangsdk.ServiceClient, domainId, groupId string) ([]RoleList, error) {
	// GET https://{Endpoint}/v3/OS-INHERIT/domains/{domain_id}/groups/{group_id}/roles/inherited_to_projects
	raw, err := client.Get(client.ServiceURL("OS-INHERIT", "domains", domainId, "groups", groupId, "roles", "inherited_to_projects"),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res []RoleList
	err = extract.IntoSlicePtr(raw.Body, &res, "roles")
	return res, err
}

type RoleList struct {
	Flag        string `json:"flag"`
	Catalog     string `json:"catalog"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Links       Links  `json:"links"`
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
	Policy      Policy `json:"policy"`
	UpdatedTime string `json:"updated_time"`
	CreatedTime string `json:"created_time"`
}

type Links struct {
	Self     string `json:"self"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
}
