package zones

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DisableDNSSEC removes the signatures from a public zone.
// The DS record must be deleted at the registrar first, otherwise validating resolvers fail to resolve the zone.
func DisableDNSSEC(client *golangsdk.ServiceClient, zoneID string) (*DNSSEC, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("zones", zoneID, "disable-dnssec").Build()
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
