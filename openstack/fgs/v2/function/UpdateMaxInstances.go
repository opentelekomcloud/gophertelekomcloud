package function

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateFuncInstancesOpts struct {
	FuncUrn        string `json:"-"`
	MaxInstanceNum int    `json:"max_instance_num,omitempty"`
}

func UpdateMaxInstances(client *golangsdk.ServiceClient, opts UpdateFuncInstancesOpts) (*FuncGraph, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("fgs", "functions", opts.FuncUrn, "config-max-instance"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res FuncGraph
	return &res, extract.Into(raw.Body, &res)
}
