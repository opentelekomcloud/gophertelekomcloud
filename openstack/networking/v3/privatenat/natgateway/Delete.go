package natgateway

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// This function is used to delete a private NAT gateway.
func Delete(client *golangsdk.ServiceClient, gatewayId string) error {
	// DELETE /v3/{project_id}/private-nat/gateways/{gateway_id}
	_, err := client.Delete(client.ServiceURL("private-nat", "gateways", gatewayId), &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return err
}
