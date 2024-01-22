package app_code

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID string `json:"-"`
	AppID     string `json:"-"`
	AppCode   string `json:"app_code" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CodeResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "apps",
		opts.AppID, "app-codes"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res CodeResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CodeResp struct {
	ID         string `json:"id"`
	AppID      string `json:"app_id"`
	AppCode    string `json:"app_code"`
	CreateTime string `json:"create_time"`
}
