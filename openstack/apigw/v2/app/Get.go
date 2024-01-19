package app

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, gatewayID, appID string) (*AppResp, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", gatewayID, "apps", appID),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res AppResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
