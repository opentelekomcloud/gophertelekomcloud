package authorizer

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, gatewayID, authorizerID string) (*AuthorizerResp, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", gatewayID, "authorizers", authorizerID),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res AuthorizerResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
