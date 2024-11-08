package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ModifySpecOpt struct {
	TargetType     string `json:"target_type,omitempty"`
	TargetId       string `json:"target_id" required:"true"`
	TargetSpecCode string `json:"target_spec_code" required:"true"`
	InstanceId     string `json:"-"`
}

func ModifySpec(client *golangsdk.ServiceClient, opts ModifySpecOpt) (*string, error) {
	b, err := build.RequestBody(opts, "resize")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "resize"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return ExtractJob(err, raw)
}
