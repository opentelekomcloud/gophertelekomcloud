package ips

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetProtectionModeQueryParameters struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `q:"object_id" required:"true"`
}

// This function is used to query a protection mode.
// objectId: Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
// In the return value, find the ID in ProtectObjects[n].ObjectID.
// If the value of type is 0, the protected object ID belongs to the Internet border.
// If the value of type is 1, the protected object ID belongs to the VPC border.
func GetProtectionMode(client *golangsdk.ServiceClient, objectId string) (*IpsProtectModeObject, error) {
	//GET /v1/{project_id}/ips/protect
	url, err := golangsdk.NewURLBuilder().WithEndpoints("ips", "protect").WithQueryParams(&GetProtectionModeQueryParameters{
		ObjectID: objectId,
	}).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetProtectionModeResponse
	err = extract.Into(raw.Body, &res)
	return &res.Data, err

}

type GetProtectionModeResponse struct {
	// Returned value for querying the IPS switch.
	Data IpsProtectModeObject `json:"data"`
}

type IpsProtectModeObject struct {
	// IPS protection mode ID.
	ID string `json:"id"`
	// IPS protection mode: 0 (observation mode), 1 (strict mode), 2 (medium mode), or 3 (loose mode).
	// The observation mode is the default mode.
	Mode int `json:"mode"`
}
