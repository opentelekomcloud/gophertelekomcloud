package endpoints

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// AddIP is used to bind an IP address to an endpoint.
// The response doesn't contain the IP address ID, use ListIPs to get it.
// Send POST /v2.1/endpoints/{endpoint_id}/ipaddresses
func AddIP(client *golangsdk.ServiceClient, endpointID string, opts IPAddressInfo) (*Endpoint, error) {
	if endpointID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "endpointID"}
	}

	b, err := build.RequestBody(opts, "ipaddress")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(v21URL(client, "endpoints", endpointID, "ipaddresses"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res Endpoint
	err = extract.IntoStructPtr(raw.Body, &res, "endpoint")
	return &res, err
}
