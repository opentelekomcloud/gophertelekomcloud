package throttling_policy

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, throttleID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "throttles", throttleID), nil)
	return
}
