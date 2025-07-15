package response

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, gatewayID, groupID, id string) (*Response, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", gatewayID, "api-groups", groupID, "gateway-responses", id),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res Response
	err = extract.Into(raw.Body, &res)
	return &res, err
}
