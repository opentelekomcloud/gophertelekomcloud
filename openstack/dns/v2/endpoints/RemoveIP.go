package endpoints

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// RemoveIP is used to unbind an IP address from an endpoint.
// Send DELETE /v2.1/endpoints/{endpoint_id}/ipaddresses/{ipaddress_id}
func RemoveIP(client *golangsdk.ServiceClient, endpointID, ipAddressID string) error {
	if endpointID == "" {
		return golangsdk.ErrMissingInput{Argument: "endpointID"}
	}
	if ipAddressID == "" {
		return golangsdk.ErrMissingInput{Argument: "ipAddressID"}
	}

	_, err := client.Delete(v21URL(client, "endpoints", endpointID, "ipaddresses", ipAddressID), &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return err
}
