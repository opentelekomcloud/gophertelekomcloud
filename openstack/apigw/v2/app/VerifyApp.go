package app

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func VerifyApp(client *golangsdk.ServiceClient, gatewayID, appID string) (*ValidationResp, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", gatewayID, "apps", "validation",
		appID),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res ValidationResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ValidationResp struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"remark"`
}
