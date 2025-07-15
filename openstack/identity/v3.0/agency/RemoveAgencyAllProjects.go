package agency

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func RemoveAgencyAllProjects(client *golangsdk.ServiceClient, domainId, agencyId, roleId string) (err error) {
	// DELETE /v3.0/OS-INHERIT/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}/inherited_to_projects
	_, err = client.Delete(client.ServiceURL("OS-INHERIT", "domains", domainId, "agencies", agencyId, "roles", roleId, "inherited_to_projects"), nil)
	return err
}
