package api

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ManageOpts struct {
	GatewayID   string `json:"-"`
	Action      string `json:"action" required:"true"`
	EnvID       string `json:"env_id" required:"true"`
	ApiID       string `json:"api_id" required:"true"`
	Description string `json:"remark,omitempty"`
}

func ManageApi(client *golangsdk.ServiceClient, opts ManageOpts) (*ManageApiResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "apis", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res ManageApiResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ManageApiResp struct {
	PublishID   string `json:"publish_id"`
	ApiID       string `json:"api_id"`
	ApiName     string `json:"api_name"`
	EnvID       string `json:"env_id"`
	Description string `json:"remark"`
	VersionID   string `json:"version_id"`
	PublishTime string `json:"publish_time"`
}
