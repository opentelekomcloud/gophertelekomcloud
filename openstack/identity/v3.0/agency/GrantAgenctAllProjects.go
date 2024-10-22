package agency

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func GrantAgencyAllProjects(client *golangsdk.ServiceClient, domainId, agencyId, roleId string) (err error) {
	// PUT /v3.0/OS-INHERIT/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}/inherited_to_projects
	_, err = client.Put(client.ServiceURL("OS-INHERIT", "domains", domainId, "agencies", agencyId, "roles", roleId, "inherited_to_projects"), nil,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{204},
		})
	return err
}
