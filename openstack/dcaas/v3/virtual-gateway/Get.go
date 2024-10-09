package virtual_gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(c *golangsdk.ServiceClient, id string) (*VirtualGateway, error) {
	raw, err := c.Get(c.ServiceURL("dcaas", "virtual-gateways", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res VirtualGateway
	err = extract.IntoStructPtr(raw.Body, &res, "virtual_gateway")
	return &res, err
}
