package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateDataIpOpts struct {
	InstanceId string `json:"-"`
	// Indicates the private IP address.
	NewIp string `json:"new_ip"`
}

func UpdateDataIp(c *golangsdk.ServiceClient, opts UpdateDataIpOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/ip
	raw, err := c.Put(c.ServiceURL("instances", opts.InstanceId, "ip"), b, nil,
		&golangsdk.RequestOpts{OkCodes: []int{200}})
	return extra(err, raw)
}
