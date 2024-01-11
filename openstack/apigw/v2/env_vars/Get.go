package env_vars

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, gatewayID, envVarID string) (*EnvVarsResp, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", gatewayID, "env-variables", envVarID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res EnvVarsResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
