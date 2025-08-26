package natgateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query details about a specified private NAT gateway.
func Get(client *golangsdk.ServiceClient, gatewayId string) (*GatewayCommonResponse, error) {
	// GET /v3/{project_id}/private-nat/gateways/{gateway_id}
	raw, err := client.Get(client.ServiceURL("private-nat", "gateways", gatewayId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GatewayCommonResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
