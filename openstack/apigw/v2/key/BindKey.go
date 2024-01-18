package key

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BindOpts struct {
	GatewayID  string   `json:"-"`
	SignID     string   `json:"sign_id" required:"true"`
	PublishIds []string `json:"publish_ids" required:"true"`
}

func BindKey(client *golangsdk.ServiceClient, opts BindOpts) ([]BindSignResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "sign-bindings"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{201},
		})
	if err != nil {
		return nil, err
	}

	var res []BindSignResp

	err = extract.IntoSlicePtr(raw.Body, &res, "bindings")
	return res, err
}

type BindSignResp struct {
	PublishID      string `json:"publish_id"`
	ApiID          string `json:"api_id"`
	ApiType        int    `json:"api_type"`
	ApiName        string `json:"api_name"`
	ApiDescription string `json:"api_remark"`
	EnvID          string `json:"env_id"`
	EnvName        string `json:"env_name"`
	GroupName      string `json:"group_name"`
	Name           string `json:"name"`
	SignType       string `json:"sign_type"`
	SignKey        string `json:"sign_key"`
	SignSecret     string `json:"sign_secret"`
	SignAlgorithm  string `json:"sign_algorithm"`
	BindingTime    string `json:"binding_time"`
	ID             string `json:"id"`
	SignID         string `json:"sign_id"`
	SignName       string `json:"sign_name"`
}
