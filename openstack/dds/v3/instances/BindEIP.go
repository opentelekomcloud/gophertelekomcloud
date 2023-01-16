package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type BindEIPOpts struct {
	PublicIpId string `json:"public_ip_id" required:"true"`
	PublicIp   string `json:"public_ip" required:"true"`
	NodeId     string `json:"-"`
}

func BindEIP(client *golangsdk.ServiceClient, opts BindEIPOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("nodes", opts.NodeId, "bind-eip"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return extractJob(err, raw)
}
