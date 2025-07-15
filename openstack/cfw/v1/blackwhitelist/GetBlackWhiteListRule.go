package blackwhitelist

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query a single blacklist or whitelist rule.
// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
// In the return value, find the ID in ProtectObjects[n].ObjectID.
// If the value of type is 0, the protected object ID belongs to the Internet border.
// If the value of type is 1, the protected object ID belongs to the VPC border.
// listType: 4 (blacklist), 5 (whitelist).
// ipAddress: IP address used to create rule.
func GetBlacklistOrWhitelistRule(client *golangsdk.ServiceClient, objectId string, listType int, ipAddress string) (*BlackWhiteListRecord, error) {
	// GET /v1/{project_id}/black-white-lists
	url, err := golangsdk.NewURLBuilder().WithEndpoints("black-white-lists").WithQueryParams(&ListQueryParameters{
		ObjectID: objectId,
		ListType: listType,
		Address:  ipAddress,
		Limit:    1024,
		Offset:   "0",
	}).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetListResponse
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}
	if len(res.Data.Records) != 0 {
		return &res.Data.Records[0], nil
	}
	return nil, fmt.Errorf("no blacklist/whitelist rule found")
}
