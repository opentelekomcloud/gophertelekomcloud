package app_code

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GenerateAppCode(client *golangsdk.ServiceClient, gatewayID, appID string) (*CodeResp, error) {
	raw, err := client.Put(client.ServiceURL("apigw", "instances", gatewayID, "apps",
		appID, "app-codes"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res CodeResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}
