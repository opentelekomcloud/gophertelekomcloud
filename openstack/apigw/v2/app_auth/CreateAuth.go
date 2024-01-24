package app_auth

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateAuthOpts struct {
	GatewayID string   `json:"-"`
	EnvID     string   `json:"env_id" required:"true"`
	AppIDs    []string `json:"app_ids" required:"true"`
	ApiIDs    []string `json:"api_ids" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateAuthOpts) ([]AppAuthResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "app-auths"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res []AppAuthResp

	err = extract.IntoSlicePtr(raw.Body, &res, "auths")
	return res, err
}

type AppAuthResp struct {
	ID            string     `json:"id"`
	ApiID         string     `json:"api_id"`
	AuthResult    AuthResult `json:"auth_result"`
	AuthTime      string     `json:"auth_time"`
	AppID         string     `json:"app_id"`
	AuthRole      string     `json:"auth_role"`
	AuthTunnel    string     `json:"auth_tunnel"`
	AuthWhitelist []string   `json:"auth_whitelist"`
	AuthBlacklist []string   `json:"auth_blacklist"`
	VisitParams   string     `json:"visit_params"`
}

type AuthResult struct {
	Status    string `json:"status"`
	ErrorMsg  string `json:"error_msg"`
	ErrorCode string `json:"error_code"`
	ApiName   string `json:"api_name"`
	AppName   string `json:"app_name"`
}
