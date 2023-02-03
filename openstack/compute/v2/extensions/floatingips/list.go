package floatingips

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List returns a Pager that allows you to iterate over a collection of FloatingIPs.
func List(client *golangsdk.ServiceClient) ([]FloatingIP, error) {
	raw, err := client.Get(client.ServiceURL("os-floating-ips"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []FloatingIP
	err = extract.IntoSlicePtr(raw.Body, &res, "floating_ips")
	return res, err
}
