package throttling_policy

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, gatewayID, throttleID string) (*ThrottlingResp, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", gatewayID, "throttles", throttleID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ThrottlingResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
