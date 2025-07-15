package response

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, groupID, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "api-groups", groupID, "gateway-responses", id), nil)
	return
}
