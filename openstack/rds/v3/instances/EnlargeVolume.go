package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type EnlargeVolumeRdsOpts struct {
	InstanceId string `json:"-"`
	// The minimum start value of each scaling is 10 GB. A DB instance can be scaled up only by a multiple of 10 GB. Value range: 10 GB to 4000 GB
	Size int `json:"size" required:"true"`
}

func EnlargeVolume(client *golangsdk.ServiceClient, opts EnlargeVolumeRdsOpts) (*string, error) {
	b, err := build.RequestBody(&opts, "enlarge_volume")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/action
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "action"), b, nil, nil)
	return extraJob(err, raw)
}
