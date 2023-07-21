package roles

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func CheckGroupAllProjects(client *golangsdk.ServiceClient, domainId, groupId, roleId string) error {
	// HEAD https://{Endpoint}/v3/OS-INHERIT/domains/{domain_id}/groups/{group_id}/roles/{role_id}/inherited_to_projects
	_, err := client.Head(client.ServiceURL("OS-INHERIT", "domains", domainId, "groups", groupId, "roles", roleId, "inherited_to_projects"),
		nil)
	if err != nil {
		return err
	}
	return nil
}
