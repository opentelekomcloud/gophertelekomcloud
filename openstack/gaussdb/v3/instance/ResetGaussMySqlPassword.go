package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type MysqlResetPasswordOpts struct {
	// Instance ID, which is compliant with the UUID format.
	InstanceId string `json:"instance_id"`
	// Database password.
	// Value range:
	// The password consists of 8 to 32 characters and contains at least three types of the following:
	// uppercase letters, lowercase letters, digits, and special characters (~!@#$%^*-_=+?,()&).
	// Enter a strong password to improve security, preventing security risks such as brute force cracking.
	// If you enter a weak password, the system automatically determines that the password is invalid.
	Password string `json:"password"`
}

func ResetGaussMySqlPassword(client *golangsdk.ServiceClient, opts MysqlResetPasswordOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}/password
	_, err = client.Post(client.ServiceURL("instances", opts.InstanceId), b, nil, nil)
	return err
}
