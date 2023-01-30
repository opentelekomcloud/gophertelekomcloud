package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type SwitchSslOpts struct {
	// Specifies the DB instance ID.
	InstanceId string `json:"-"`
	// Specifies whether to enable SSL.
	// true: SSL is enabled.
	// false: SSL is disabled.
	SslOption bool `json:"ssl_option"`
}

// SwitchSsl This API is supported for MySQL only.
func SwitchSsl(c *golangsdk.ServiceClient, opts SwitchSslOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/ssl
	_, err = c.Put(c.ServiceURL("instances", opts.InstanceId, "ssl"), b, nil,
		&golangsdk.RequestOpts{OkCodes: []int{200}})
	return
}
