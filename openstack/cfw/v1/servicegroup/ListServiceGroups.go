package servicegroup

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

// This function is used to query the service group list.
// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
// In the return value, find the ID in ProtectObjects[n].ObjectID.
// If the value of type is 0, the protected object ID belongs to the Internet border.
// If the value of type is 1, the protected object ID belongs to the VPC border.
func ListServiceGroups(client *golangsdk.ServiceClient, objectId string) ([]ServiceGroupRecord, error) {
	// GET /v1/{project_id}/service-sets
	url, err := golangsdk.NewURLBuilder().WithEndpoints("service-sets").WithQueryParams(&ListQueryParameters{
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
	// Returned data for querying the service group list.
	Data ServiceGroupsData `json:"data"`
}

type ServiceGroupsData struct {
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `json:"limit"`
	// Offset, which specifies the start position of the record to be returned.
	Offset int `json:"offset"`
	// Query the total number of service group records.
	Total int `json:"total"`
	// The list of service group records
	Records []ServiceGroupRecord `json:"records"`
}

type ServiceGroupRecord struct {
	// Service group ID.
	SetID string `json:"set_id"`
	// Service Group Name
	Name string `json:"name"`
	// Description.
	Description string `json:"description"`
	// Service group type:
	// 0 - user-defined service group
	// 1 - common web service
	// 2 - common remote login and ping
	// 3 - common database
	ServiceSetType int `json:"service_set_type"`
	// Number of times an service group is referenced by rules.
	RefCount int `json:"ref_count"`
	// Project ID.
	ProjectID string `json:"project_id"`
	// Protocol list. Protocol type: 6 (TCP), 17 (UDP), 1 (ICMP), 58 (ICMPv6), or -1 (any).
	Protocols []int `json:"protocols"`
}
