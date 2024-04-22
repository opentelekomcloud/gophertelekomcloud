package api

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type VersionApiOpts struct {
	GatewayID string `json:"-"`
	ApiID     string `json:"-"`
	VersionID string `json:"version_id" required:"true"`
}

func SwitchVersion(client *golangsdk.ServiceClient, opts ManageOpts) (*ManageApiResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("apigw", "instances", opts.GatewayID, "apis", "publish", opts.ApiID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ManageApiResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}
