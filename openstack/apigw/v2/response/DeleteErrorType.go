package response

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func DeleteErrorType(client *golangsdk.ServiceClient, gatewayID, groupID, id, respType string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "api-groups", groupID, "gateway-responses", id, respType), nil)
	return
}
