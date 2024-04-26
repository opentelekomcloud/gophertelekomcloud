package acl

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BindOpts struct {
	GatewayID string `json:"-"`
	// Access control policy ID.
	PolicyID string `json:"acl_id" required:"true"`
	// API publication record ID.
	PublishIds []string `json:"publish_ids" required:"true"`
}

func BindPolicy(client *golangsdk.ServiceClient, opts BindOpts) ([]BindAclResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "acl-bindings"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{201},
		})
	if err != nil {
		return nil, err
	}

	var res []BindAclResp

	err = extract.IntoSlicePtr(raw.Body, &res, "acl_bindings")
	return res, err
}

type BindAclResp struct {
	// Binding record ID.
	ID string `json:"id"`
	// API ID.
	ApiId string `json:"api_id"`
	// Environment ID.
	EnvId string `json:"env_id"`
	// Access control policy ID.
	PolicyID string `json:"acl_id"`
	// Binding time.
	AppliedAt string `json:"create_time"`
}
