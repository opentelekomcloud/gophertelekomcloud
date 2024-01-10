package env

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID   string `json:"-"`
	Name        string `json:"name" required:"true"`
	Description string `json:"remark,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*EnvResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "envs"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res EnvResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type EnvResp struct {
	ID          string `json:"id"`
	CreateTime  string `json:"create_time"`
	Name        string `json:"name"`
	Description string `json:"remark"`
}
