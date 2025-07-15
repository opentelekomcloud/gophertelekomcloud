package dns

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to obtain the info of a domain name in a domain name group.
// domainName: Domain name, for example, www.test.com.
// groupId: Domain name group ID. It is the same as ID retuned while creating a Domain Name Group.
// firewallId: Firewall Instance ID.
func GetDomainNameInfo(client *golangsdk.ServiceClient, domainName, groupId, firewallId string) (*DomainInfo, error) {
	// GET /v1/{project_id}/domain-set/domains/{domain_set_id}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("domain-set", "domains", groupId).WithQueryParams(&GetDomainNameListQueryParams{
		FwInstanceID: firewallId,
		Limit:        1024,
		Offset:       "0",
		DomainName:   domainName,
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
	if err != nil {
		return nil, err
	}
	if len(res.Data.Records) != 0 {
		return &res.Data.Records[0], nil
	}
	return nil, fmt.Errorf("%s not found in domain name group", domainName)
}
