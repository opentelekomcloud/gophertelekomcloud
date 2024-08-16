package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ModifyPortOpt struct {
	Port       int    `json:"port" required:"true"`
	InstanceId string `json:"-"`
}

func ModifyPort(client *golangsdk.ServiceClient, opts ModifyPortOpt) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "modify-port"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return ExtractJob(err, raw)
}
