package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ModifyInternalIpOpts struct {
	NewIp      string `json:"new_ip" required:"true"`
	NodeId     string `json:"node_id" required:"true"`
	InstanceId string `json:"-"`
}

func ModifyInternalIp(client *golangsdk.ServiceClient, opts ModifyInternalIpOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "modify-internal-ip"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return ExtractJob(err, raw)
}
