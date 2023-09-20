package roles

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func UpdateGroupAllProjects(client *golangsdk.ServiceClient, domainId, groupId, roleId string) error {
	// PUT https://{Endpoint}/v3/OS-INHERIT/domains/{domain_id}/groups/{group_id}/roles/{role_id}/inherited_to_projects
	_, err := client.Put(client.ServiceURL("OS-INHERIT", "domains", domainId, "groups", groupId, "roles", roleId, "inherited_to_projects"),
		nil, nil, &golangsdk.RequestOpts{
			OkCodes: []int{204},
		})
	if err != nil {
		return err
	}
	return nil
}
