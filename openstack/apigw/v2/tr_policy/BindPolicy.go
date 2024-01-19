package throttling_policy

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BindOpts struct {
	GatewayID  string   `json:"-"`
	PolicyID   string   `json:"strategy_id" required:"true"`
	PublishIds []string `json:"publish_ids" required:"true"`
}

func BindPolicy(client *golangsdk.ServiceClient, opts BindOpts) ([]BindThrottleResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "throttle-bindings"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{201},
		})
	if err != nil {
		return nil, err
	}

	var res []BindThrottleResp

	err = extract.IntoSlicePtr(raw.Body, &res, "throttle_applys")
	return res, err
}

type BindThrottleResp struct {
	PublishID  string `json:"publish_id"`
	Scope      int    `json:"scope"`
	StrategyID string `json:"strategy_id"`
	ApplyTime  string `json:"apply_time"`
	ID         string `json:"id"`
}
