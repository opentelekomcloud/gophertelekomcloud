package dns

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to obtain the list of domain name groups.
// firewallId: Firewall Instance ID.
func ListDomainNameGroups(client *golangsdk.ServiceClient, firewallId string) ([]DomainSetVO, error) {
	// GET /v1/{project_id}/domain-sets
	url, err := golangsdk.NewURLBuilder().WithEndpoints("domain-sets").WithQueryParams(&GetDomainNameGroupListQueryParams{
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

	var res GetDomainNameGroupDataResponse
	err = extract.Into(raw.Body, &res)
	return res.Data.Records, err
}
