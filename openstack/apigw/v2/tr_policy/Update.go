package throttling_policy

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	GatewayID             string `json:"-"`
	ThrottleID            string `json:"-"`
	Name                  string `json:"name" required:"true"`
	AppCallLimits         *int   `json:"app_call_limits,omitempty"`
	ApiCallLimits         *int   `json:"api_call_limits" required:"true"`
	TimeInterval          *int   `json:"time_interval" required:"true"`
	TimeUnit              string `json:"time_unit" required:"true"`
	Description           string `json:"remark,omitempty"`
	Type                  *int   `json:"type,omitempty"`
	IpCallLimits          *int   `json:"ip_call_limits,omitempty"`
	UserCallLimits        *int   `json:"user_call_limits,omitempty"`
	EnableAdaptiveControl string `json:"enable_adaptive_control,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*ThrottlingResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("apigw", "instances", opts.GatewayID, "throttles", opts.ThrottleID),
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
