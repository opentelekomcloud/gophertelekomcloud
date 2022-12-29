package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type RestartOpts struct {
	InstanceId string `json:"-"`
	TargetType string `json:"target_type,omitempty"`
	TargetId   string `json:"target_id" required:"true"`
}

func Restart(client *golangsdk.ServiceClient, opts RestartOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "restart"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return extractJob(err, raw)
}
