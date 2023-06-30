package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type EnableConfigIpOpts struct {
	Type       string `json:"type" required:"true"`
	TargetId   string `json:"target_id,omitempty"`
	Password   string `json:"password" required:"true"`
	InstanceId string `json:"-"`
}

func EnableConfigIp(client *golangsdk.ServiceClient, opts EnableConfigIpOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("instances", opts.InstanceId, "create-ip"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return err
	}
	return nil
}
