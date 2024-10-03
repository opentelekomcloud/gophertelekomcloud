package quota

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient) (*Quota, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("vpn", "quotas").Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Quota
	err = extract.IntoStructPtr(raw.Body, &res, "quotas")
	return &res, err
}

type Quota struct {
	Resources []Info `json:"resources"`
}

type Info struct {
	// Specifies a resource type.
	// Value range:
	// customer_gateway: customer gateway
	// vpn_connection: Enterprise Edition VPN connection
	// vpn_gateway: Enterprise Edition VPN gateway
	Type string `json:"type"`
	// Specifies the quota upper limit.
	Quota int `json:"quota"`
	// Specifies the number of resources in use.
	Used int `json:"used"`
}
