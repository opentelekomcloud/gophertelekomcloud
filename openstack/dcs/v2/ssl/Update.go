package ssl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type SslOpts struct {
	InstanceId string `json:"-"`
	Enabled    *bool  `json:"enabled" required:"true"`
}

func Update(client *golangsdk.ServiceClient, opts SslOpts) (*SslUpdateResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url := client.ServiceURL("instances", opts.InstanceId, "ssl")
	raw, err := client.Put(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	if err != nil {
		return nil, err
	}
	var res SslUpdateResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type SslUpdateResp struct {
	// Actual output parameters of API in which `result` is always `null`
	JobId      string `json:"jobId"`
	InstanceId string `json:"instanceId"`
	Result     string `json:"result"`
}
