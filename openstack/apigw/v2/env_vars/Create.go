package env_vars

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID     string `json:"-"`
	GroupID       string `json:"group_id" required:"true"`
	EnvID         string `json:"env_id" required:"true"`
	Name          string `json:"name" required:"true"`
	VariableName  string `json:"variable_name" required:"true"`
	VariableValue string `json:"variable_value" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*EnvVarsResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "env-variables"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res EnvVarsResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type EnvVarsResp struct {
	EnvID         string `json:"env_id"`
	GroupID       string `json:"group_id"`
	ID            string `json:"id"`
	VariableName  string `json:"variable_name"`
	VariableValue string `json:"variable_value"`
}
