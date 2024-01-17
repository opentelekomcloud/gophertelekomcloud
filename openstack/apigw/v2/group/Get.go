package group

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, gatewayID, groupID string) (*GroupResp, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", gatewayID, "api-groups", groupID),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res GroupResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
