package key

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, groupID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "signs", groupID), nil)
	return
}
