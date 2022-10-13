package rescueunrescue

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// RescueOpts represents the configuration options used to control a Rescue option.
type RescueOpts struct {
	// AdminPass is the desired administrative password for the instance in RESCUE mode.
	// If it's left blank, the server will generate a password.
	AdminPass string `json:"adminPass,omitempty"`
	// RescueImageRef contains reference on an image that needs to be used as rescue image.
	// If it's left blank, the server will be rescued with the default image.
	RescueImageRef string `json:"rescue_image_ref,omitempty"`
}

// Rescue instructs the provider to place the server into RESCUE mode.
func Rescue(client *golangsdk.ServiceClient, id string, opts RescueOpts) (string, error) {
	b, err := build.RequestBody(opts, "rescue")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("servers", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		AdminPass string `json:"adminPass"`
	}
	err = extract.Into(raw.Body, &res)
	return res.AdminPass, err
}

// UnRescue instructs the provider to return the server from RESCUE mode.
func UnRescue(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Post(client.ServiceURL("servers", id, "action"), map[string]interface{}{"unrescue": nil}, nil, nil)
	return
}
