package dns

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to obtain the list of domain names in a domain name group.
// groupId: Domain name group ID. It is the same as ID retuned while creating a Domain Name Group.
// firewallId: Firewall Instance ID.
func ListDomainNames(client *golangsdk.ServiceClient, groupId, firewallId string) ([]DomainInfo, error) {
	// GET /v1/{project_id}/domain-set/domains/{domain_set_id}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("domain-set", "domains", groupId).WithQueryParams(&GetDomainNameListQueryParams{
		FwInstanceID: firewallId,
		Limit:        1024,
		Offset:       "0",
	}).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetDomainNamesDataResponse
	err = extract.Into(raw.Body, &res)
	return res.Data.Records, err
}
