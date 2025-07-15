package dns

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to obtain the list of domain name groups.
// firewallId: Firewall Instance ID.
// groupName: Name of Domain Name Group
// objectId: Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
// In the return value, find the ID in ProtectObjects[n].ObjectID.
// If the value of type is 0, the protected object ID belongs to the Internet border.
// If the value of type is 1, the protected object ID belongs to the VPC border.
func GetDomainNameGroup(client *golangsdk.ServiceClient, groupName, firewallId, objectID string) (*DomainSetVO, error) {
	// GET /v1/{project_id}/domain-sets
	url, err := golangsdk.NewURLBuilder().WithEndpoints("domain-sets").WithQueryParams(&GetDomainNameGroupListQueryParams{
		FwInstanceID: firewallId,
		ObjectID:     objectID,
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
