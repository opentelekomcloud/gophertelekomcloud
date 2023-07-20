package groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/policies"
)

func QueryGroupAllProjects(client *golangsdk.ServiceClient, domainId, groupId, roleId string) ([]policies.ListPolicy, error) {
	// GET https://{Endpoint}/v3/OS-INHERIT/domains/{domain_id}/groups/{group_id}/roles/{role_id}/inherited_to_projects
	raw, err := client.Get(client.ServiceURL("OS-INHERIT", "domains", domainId, "groups", groupId, "roles", roleId, "inherited_to_projects"),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res []policies.ListPolicy
	err = extract.IntoSlicePtr(raw.Body, &res, "policy")
	return res, err
}
