package dns

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to obtain the list of domain name groups.
// firewallId: Firewall Instance ID.
// groupName: Name of Domain Name Group
func GetDomainNameGroups(client *golangsdk.ServiceClient, groupName, firewallId string) (*DomainSetVO, error) {
	// GET /v1/{project_id}/domain-sets
	url, err := golangsdk.NewURLBuilder().WithEndpoints("domain-sets").WithQueryParams(&GetDomainNameGroupListQueryParams{
		FwInstanceID: firewallId,
		Limit:        1024,
		Offset:       "0",
		Keyword:      groupName,
	}).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetDomainNameGroupDataResponse
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}
	if len(res.Data.Records) != 0 {
		return &res.Data.Records[0], nil
	}
	return nil, fmt.Errorf("%s not found in domain name group", groupName)
}
