package throttling_policy

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func UnbindPolicy(client *golangsdk.ServiceClient, gatewayID, throttleBindID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "throttle-bindings", throttleBindID), nil)
	return
}
