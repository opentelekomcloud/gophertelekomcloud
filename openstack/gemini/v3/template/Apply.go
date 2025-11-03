package template

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ApplyOpts struct {
	ConfigId    string   `json:"-"`
	InstanceIds []string `json:"instance_ids" required:"true"`
}

func Apply(client *golangsdk.ServiceClient, opts ApplyOpts) (*ApplyResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("configurations", opts.ConfigId, "apply"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ApplyResp
	return &res, extract.Into(raw.Body, &res)
}

type ApplyResp struct {
	JobId   string `json:"job_id"`
	Success bool   `json:"success"`
}
