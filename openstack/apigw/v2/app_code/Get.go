package app_code

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, gatewayID, appID, appCodeID string) (*CodeResp, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", gatewayID, "apps", appID,
		"app-codes", appCodeID),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res CodeResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
