package ips

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type SetProtectionModeOpts struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `json:"object_id" required:"true"`
	// IPS protection mode: 0 (observation mode), 1 (strict mode), 2 (medium mode), or 3 (loose mode).
	Mode *int `json:"mode" required:"true"`
}

// This function is used to change the protection mode.
func SetIPSProtectionMode(client *golangsdk.ServiceClient, opts SetProtectionModeOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v1/{project_id}/ips/protect
	_, err = client.Post(client.ServiceURL("ips", "protect"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return err
}
