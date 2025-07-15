package ips

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type SetFeatureStatusOpts struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `json:"object_id" required:"true"`
	// Patch type. Its value can only be 2 (virtual patch).
	IpsType int `json:"ips_type" required:"true"`
	// IPS feature status: 0 (disabled), 1 (enabled).
	Status *int `json:"status" required:"true"`
}

// This function is used to enable or disable the feature.
func SetIPSFeatureStatus(client *golangsdk.ServiceClient, opts SetFeatureStatusOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v1/{project_id}/ips/switch
	_, err = client.Post(client.ServiceURL("ips", "switch"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return err
}
