package throttling_policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID             string `json:"-"`
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

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*ThrottlingResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "throttles"), b, nil, &golangsdk.RequestOpts{
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
	AppCallLimits          int    `json:"app_call_limits"`
	Name                   string `json:"name"`
	TimeUnit               string `json:"time_unit"`
	Description            string `json:"remark"`
	ApiCallLimits          int    `json:"api_call_limits"`
	Type                   int    `json:"type"`
	EnableAdaptiveControl  string `json:"enable_adaptive_control"`
	UserCallLimits         int    `json:"user_call_limits"`
	TimeInterval           int    `json:"time_interval"`
	IpCallLimits           int    `json:"ip_call_limits"`
	ID                     string `json:"id"`
	BindNum                int    `json:"bind_num"`
	IsIncluSpecialThrottle int    `json:"is_inclu_special_throttle"`
	CreateTime             string `json:"create_time"`
}
