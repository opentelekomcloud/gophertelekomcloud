package group

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, groupID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "api-groups", groupID), nil)
	return
}
