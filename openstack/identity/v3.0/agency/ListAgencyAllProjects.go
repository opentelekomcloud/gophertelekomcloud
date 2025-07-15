package agency

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListAgencyAllProjects(client *golangsdk.ServiceClient, domainId, agencyId string) (*Roles, error) {
	// GET /v3.0/OS-INHERIT/domains/{domain_id}/agencies/{agency_id}/roles/inherited_to_projects
	raw, err := client.Get(client.ServiceURL("OS-INHERIT", "domains", domainId, "agencies", agencyId, "roles", "inherited_to_projects"), nil,
		nil)
	if err != nil {
		return nil, err
	}

	var res Roles

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Roles struct {
	Roles []Role `json:"roles"`
	Links Link   `json:"links"`
}

type Role struct {
	Id    string `json:"id"`
	Links Link   `json:"links"`
	Name  string `json:"name"`
}

type Link struct {
	Self string `json:"self"`
}
