package key

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListBoundOpts struct {
	GatewayID string `json:"-"`
	ApiID     string `q:"api_id"`
	SignID    string `q:"sign_id"`
	EnvID     string `q:"env_id"`
	ApiName   string `q:"api_name"`
	GroupID   string `q:"group_id"`
}

// ListAPIBoundKeys This func basically copies ListBoundKeys without signature info
// I guess the intended way was to return []ApiForSign here but something went wrong along the way
func ListAPIBoundKeys(client *golangsdk.ServiceClient, opts ListBoundOpts) ([]BindSignResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client: client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "sign-bindings",
			"binded-apis") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return BindingPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractBindings(pages)
}
