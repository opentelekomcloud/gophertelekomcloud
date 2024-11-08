package customer_gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, gatewayId string) (*CustomerGateway, error) {
	raw, err := client.Get(client.ServiceURL("customer-gateways", gatewayId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res CustomerGateway
	err = extract.IntoStructPtr(raw.Body, &res, "customer_gateway")
	return &res, err
}
