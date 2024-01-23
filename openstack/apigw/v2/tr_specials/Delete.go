package special_policy

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, gatewayID, throttleID, strategyID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "throttles", throttleID,
		"throttle-specials", strategyID), nil)
	return
}
