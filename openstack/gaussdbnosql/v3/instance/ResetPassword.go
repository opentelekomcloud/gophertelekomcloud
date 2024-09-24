package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ResetPasswordOpts struct {
	InstanceId string
	// Database password
	// The value must be 8 to 32 characters in length and contain uppercase letters (A to Z),
	// lowercase letters (a to z), digits (0 to 9), and special characters, such as ~!@#%^*-_=+?
	// You are advised to enter a strong password to improve security, preventing security risks such as brute force cracking.
	Password string `json:"password"`
}

func ResetPassword(client *golangsdk.ServiceClient, opts ResetPasswordOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/password
	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "password"), b, nil, nil)
	return
}
