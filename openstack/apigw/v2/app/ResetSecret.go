package app

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ResetOpts struct {
	GatewayID string `json:"-"`
	AppID     string `json:"-"`
	AppSecret string `json:"app_secret,omitempty"`
}

func ResetSecret(client *golangsdk.ServiceClient, opts ResetOpts) (*AppResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("apigw", "instances", opts.GatewayID, "apps", "secret",
		opts.AppID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res AppResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}
