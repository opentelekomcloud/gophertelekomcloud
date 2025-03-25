package addressgroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListQueryParameters struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `q:"object_id" required:"true"`
	// Offset, which specifies the start position of the record to be returned. The value must be a number no less than 0. The default value is 0.
	Offset string `q:"offset" required:"true"`
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `q:"limit" required:"true"`
}

// This function is used to query the address group list.
// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
// In the return value, find the ID in ProtectObjects[n].ObjectID.
// If the value of type is 0, the protected object ID belongs to the Internet border.
// If the value of type is 1, the protected object ID belongs to the VPC border.
func ListAddressGroups(client *golangsdk.ServiceClient, objectId string) ([]AddressGroupRecord, error) {
	// GET /v1/{project_id}/address-sets
	url, err := golangsdk.NewURLBuilder().WithEndpoints("address-sets").WithQueryParams(&ListQueryParameters{
		ObjectID: objectId,
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

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return res.Data.Records, err
}

type ListResponse struct {
	// Returned data for querying the address group list.
	Data AddressGroupsData `json:"data"`
}

type AddressGroupsData struct {
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `json:"limit"`
	// Offset, which specifies the start position of the record to be returned.
	Offset int `json:"offset"`
	// Query the total number of address group records.
	Total int `json:"total"`
	// The list of address group records
	Records []AddressGroupRecord `json:"records"`
}

type AddressGroupRecord struct {
	// Address group ID.
	SetID string `json:"set_id"`
	// Number of times an address group is referenced by rules.
	RefCount int `json:"ref_count"`
	// Description.
	Description string `json:"description"`
	// Internet protocol type of an address: 0 (IPv4), 1 (IPv6).
	AddressType int `json:"address_type"`
	// Protected object ID, used to distinguish between Internet border protection and VPC border protection.
	ObjectID string `json:"object_id"`
	// Address group type:
	// 0 - User-defined address group
	// 1 - WAF back-to-source IP address group
	// 2 - DDoS back-to-source IP address group
	// 3 - NAT64 address group
	AddressSetType int `json:"address_set_type"`
}
