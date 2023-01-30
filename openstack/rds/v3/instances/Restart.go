package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type RestartOpts struct {
	InstanceId string `json:"-"`
	// This parameter is left blank.
	Restart struct{} `json:"restart"`
}

func Restart(client *golangsdk.ServiceClient, opts RestartOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/action
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return extraJob(err, raw)
}
