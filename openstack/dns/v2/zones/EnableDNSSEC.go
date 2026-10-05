package zones

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// EnableDNSSEC signs a public zone and returns the generated key material.
// The DS record has to be published at the registrar afterwards to complete the chain of trust.
func EnableDNSSEC(client *golangsdk.ServiceClient, zoneID string) (*DNSSEC, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("zones", zoneID, "enable-dnssec").Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL(url.String()), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}
	var res DNSSEC
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
