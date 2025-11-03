package template

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateParametersOpts struct {
	InstanceId string            `json:"-"`
	Values     map[string]string `json:"values" required:"true"`
}

func UpdateInstanceParameters(client *golangsdk.ServiceClient, opts UpdateParametersOpts) (*UpdateParametersResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "configurations"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateParametersResp
	return &res, extract.Into(raw.Body, &res)
}

type UpdateParametersResp struct {
	JobId           string `json:"job_id"`
	RestartRequired bool   `json:"restart_required"`
}
