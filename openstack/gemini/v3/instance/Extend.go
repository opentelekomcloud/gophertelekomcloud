package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ExtendVolumeOpts struct {
	InstanceId string `json:"-"`
	Size       int    `json:"size" required:"true"`
}

func ExtendVolume(client *golangsdk.ServiceClient, opts ExtendVolumeOpts) (*ExtendVolumeResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "extend-volume"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res ExtendVolumeResponse
	return &res, extract.Into(raw.Body, &res)
}

type ExtendVolumeResponse struct {
	JobId string `json:"job_id"`
}
