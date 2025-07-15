package connection_monitoring

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Specifies a VPN connection ID.
	ConnectionId string `q:"vpn_connection_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Monitor, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("connection-monitors").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Monitor
	err = extract.IntoSlicePtr(raw.Body, &res, "connection_monitors")
	return res, err
}
