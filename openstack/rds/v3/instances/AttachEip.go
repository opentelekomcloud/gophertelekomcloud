package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type AttachEipOpts struct {
	InstanceId string `json:"-"`
	// NOTICE
	// When is_bind is true, public_ip is mandatory.
	// Specifies the EIP to be bound. The value must be in the standard IP address format.
	PublicIp string `json:"public_ip,omitempty"`
	// NOTICE
	// When is_bind is true, public_ip_id is mandatory.
	// Specifies the EIP ID. The value must be in the standard UUID format.
	PublicIpId string `json:"public_ip_id,omitempty"`
	// true: Bind an EIP.
	// false: Unbind an EIP.
	IsBind *bool `json:"is_bind" required:"true"`
}

func AttachEip(client *golangsdk.ServiceClient, opts AttachEipOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/public-ip
	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "public-ip"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
