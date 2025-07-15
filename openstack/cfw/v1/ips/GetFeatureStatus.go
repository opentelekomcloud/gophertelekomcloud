package ips

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetIPSFeatureStatusQueryParameters struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `q:"object_id" required:"true"`
}

// This function is used to query the status of the IPS feature.
// objectId: Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
// In the return value, find the ID in ProtectObjects[n].ObjectID.
// If the value of type is 0, the protected object ID belongs to the Internet border.
// If the value of type is 1, the protected object ID belongs to the VPC border.
func GetIPSFeatureStatus(client *golangsdk.ServiceClient, objectId string) (*IpsSwitchResponseDTO, error) {
	// GET /v1/{project_id}/ips/switch
	url, err := golangsdk.NewURLBuilder().WithEndpoints("ips", "switch").WithQueryParams(&GetIPSFeatureStatusQueryParameters{
		ObjectID: objectId,
	}).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetIPSFeatureStatusResponse
	err = extract.Into(raw.Body, &res)
	return &res.Data, err

}

type GetIPSFeatureStatusResponse struct {
	// Returned value for querying the IPS switch.
	Data IpsSwitchResponseDTO `json:"data"`
}

type IpsSwitchResponseDTO struct {
	// IPS switch ID.
	ID string `json:"id"`
	// Basic defense status: 0 (disabled), 1 (enabled).
	BasicDefenseStatus int `json:"basic_defense_status"`
	// Virtual patch status: 0 (disabled), 1 (enabled).
	VirtualPatchesStatus int `json:"virtual_patches_status"`
}
