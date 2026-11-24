package endpoints

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete is used to delete an endpoint.
// Send DELETE /v2.1/endpoints/{endpoint_id}
func Delete(client *golangsdk.ServiceClient, endpointID string) error {
	if endpointID == "" {
		return golangsdk.ErrMissingInput{Argument: "endpointID"}
	}

	_, err := client.Delete(v21URL(client, "endpoints", endpointID), &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return err
}
