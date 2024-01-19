package app

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID   string `json:"-"`
	Name        string `json:"name" required:"true"`
	Description string `json:"remark,omitempty"`
	AppKey      string `json:"app_key,omitempty"`
	AppSecret   string `json:"app_secret,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*AppResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "apps"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res AppResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type AppResp struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"remark"`
	Creator      string `json:"creator"`
	UpdateTime   string `json:"update_time"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
	RegisterTime string `json:"register_time"`
	Status       int    `json:"status"`
	AppType      string `json:"app_type"`
	BindNum      int    `json:"bind_num"`
}
