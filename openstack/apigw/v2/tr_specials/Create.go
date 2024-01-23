package special_policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID  string `json:"-"`
	ThrottleID string `json:"throttle_id" required:"true"`
	CallLimits int    `json:"call_limits" required:"true"`
	ObjectID   string `json:"object_id" required:"true"`
	ObjectType string `json:"object_type" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*ThrottlingResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "throttles",
		opts.ThrottleID, "throttle-specials"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res ThrottlingResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ThrottlingResp struct {
	ID         string `json:"id"`
	CallLimits int    `json:"call_limits"`
	ApplyTime  string `json:"apply_time"`
	AppName    string `json:"app_name"`
	AppID      string `json:"app_id"`
	ObjectID   string `json:"object_id"`
	ObjectType string `json:"object_type"`
	ObjectName string `json:"object_name"`
	ThrottleID string `json:"throttle_id"`
}
