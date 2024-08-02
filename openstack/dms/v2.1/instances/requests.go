package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CrossVpcUpdateOpts is the structure required by the UpdateCrossVpc method to update the internal IP address for
// cross-VPC access.
type CrossVpcUpdateOpts struct {
	// User-defined advertised IP contents key-value pair.
	// The key is the listeners IP.
	// The value is advertised.listeners IP, or domain name.
	Contents map[string]string `json:"advertised_ip_contents" required:"true"`
}

// UpdateCrossVpc is a method to update the internal IP address for cross-VPC access using given parameters.
func UpdateCrossVpc(c *golangsdk.ServiceClient, instanceId string, opts CrossVpcUpdateOpts) (*CrossVpc, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := c.Post(crossVpcURL(c, instanceId), body, nil, nil)
	if err != nil {
		return nil, err
	}

	var r CrossVpc
	err = extract.Into(raw.Body, &r)

	return &r, err
}

type PasswordOpts struct {
	NewPassword string `json:"new_password" required:"true"`
}

// ChangePassword is a method to update the password using given parameters.
func ChangePassword(c *golangsdk.ServiceClient, instanceId string, opts PasswordOpts) error {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = c.Post(changePasswordURL(c, instanceId), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
