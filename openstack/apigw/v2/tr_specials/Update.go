package special_policy

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	GatewayID       string `json:"-"`
	ThrottleID      string `json:"-"`
	SpecialPolicyID string `json:"strategy_id" required:"true"`
	CallLimits      int    `json:"call_limits" required:"true"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*ThrottlingResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("apigw", "instances", opts.GatewayID, "throttles", opts.ThrottleID,
		"throttle-specials", opts.SpecialPolicyID),
		b, nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	if err != nil {
		return nil, err
	}

	var res ThrottlingResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}
