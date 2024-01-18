package key

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func UnbindKey(client *golangsdk.ServiceClient, gatewayID, signBindingID string) (err error) {
	_, err = client.Delete(client.ServiceURL("apigw", "instances", gatewayID, "sign-bindings", signBindingID), nil)
	return
}
