package gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, gatewayId string) (*Gateway, error) {
	raw, err := client.Get(client.ServiceURL("vpn-gateways", gatewayId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Gateway
	err = extract.IntoStructPtr(raw.Body, &res, "vpn_gateway")
	return &res, err
}
