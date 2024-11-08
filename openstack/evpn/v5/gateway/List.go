package gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient) ([]Gateway, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("vpn-gateways").Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Gateway
	err = extract.IntoSlicePtr(raw.Body, &res, "vpn_gateways")
	return res, err
}
