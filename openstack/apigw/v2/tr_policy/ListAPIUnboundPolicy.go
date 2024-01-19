package throttling_policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListUnoundOpts struct {
	GatewayID  string `json:"-"`
	ThrottleID string `q:"throttle_id"`
	EnvID      string `q:"env_id"`
	GroupID    string `q:"group_id"`
	ApiID      string `q:"api_id"`
	ApiName    string `q:"api_name"`
}

func ListAPIUnoundPolicy(client *golangsdk.ServiceClient, opts ListBoundOpts) ([]ApiThrottle, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client: client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "throttle-bindings",
			"unbinded-apis") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return BindingPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractBindings(pages)
}
