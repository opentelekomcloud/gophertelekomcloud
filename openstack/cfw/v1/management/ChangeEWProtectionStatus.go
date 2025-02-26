package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ChangeEWProtectionStatusOpts struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in this package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectId string `json:"object_id" required:"true"`
	// Protection status: 0 (enable), 1 (disable).
	Status int `json:"status" required:"true"`
}

// This function is used to enable or disable east-west protection.
func ChangeEWProtectionStatus(client *golangsdk.ServiceClient, firewallId string, opts ChangeEWProtectionStatusOpts) (*ChangeEWProtectionStatusResponseData, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/firewall/east-west/protect
	raw, err := client.Post(client.ServiceURL("firewall", "east-west", "protect"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ChangeEWProtectionStatusResponse
	return &res.Data, extract.Into(raw.Body, &res)
}

type ChangeEWProtectionStatusResponse struct {
	// Data returned for modifying east-west protection.
	Data ChangeEWProtectionStatusResponseData `json:"data"`
}

type ChangeEWProtectionStatusResponseData struct {
	// East-west protected object ID.
	Id string `json:"id"`
}
