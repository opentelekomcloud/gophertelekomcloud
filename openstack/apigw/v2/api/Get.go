package api

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, gatewayID, apiID string) (*ApiResp, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", gatewayID, "apis", apiID),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res ApiResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
